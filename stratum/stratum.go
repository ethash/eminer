package stratum

import (
	"bufio"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"net"
	"net/url"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
)

// Client structure
type Client struct {
	sync.RWMutex

	ID uint64

	server        *url.URL
	backupServers []*url.URL
	serverIdx     int

	conn net.Conn

	username string
	password string

	worker string

	reader *bufio.Reader
	writer *bufio.Writer

	timeout time.Duration

	callBackCh     map[uint64]*Callback
	notificationCh map[string]chan *json.RawMessage

	noReconnect bool

	sick        bool
	sickRate    int
	successRate int
	FailsCount  uint64

	work *stratumWork

	mode int

	closed int32

	minerVersion string

	tls bool
}

type stratumWork struct {
	jobid          string
	seedHash       string
	headerHash     string
	target         *big.Int
	extraNonce     *big.Int
	extraNonceSize int
}

// Req struct
type Req struct {
	ID     uint64      `json:"id"`
	Method string      `json:"method"`
	Params interface{} `json:"params"`
	Worker string      `json:"worker,omitempty"`
}

// Resp struct
type Resp struct {
	ID      uint64           `json:"id,omitempty"`
	Method  string           `json:"method,omitempty"`
	Version string           `json:"jsonrpc,omitempty"`
	Result  *json.RawMessage `json:"result,omitempty"`
	Params  *json.RawMessage `json:"params,omitempty"`
	Error   interface{}      `json:"error,omitempty"`
}

// Callback struct
type Callback struct {
	method string
	resp   chan *Resp
}

const (
	connectTimeout   = 10 //10 sec for connect timeout
	responseTimeout  = 10 //10 sec for response timeout
	connectRetryTime = 5

	bufferSize = 1024
)

// stratum mode
const (
	etherproxy int = iota
	etherstratum
	stratum
)

// New stratum server client
func New(servers string, timeout time.Duration, version, username, password, worker string, tls bool) (*Client, error) {
	serversArr := strings.Split(servers, ",")

	c := &Client{
		timeout: timeout,

		username: username,
		password: password,

		worker: worker,

		callBackCh:     make(map[uint64]*Callback),
		notificationCh: make(map[string]chan *json.RawMessage),

		minerVersion: version,

		work: new(stratumWork),

		tls: tls,
	}

	for idx, server := range serversArr {
		if !strings.Contains(server, "://") {
			server = "etherproxy://" + server
		}

		url, err := url.Parse(server)
		if err != nil {
			return nil, err
		}

		if idx == 0 {
			c.server = url
		}

		c.backupServers = append(c.backupServers, url)
	}

	return c, nil
}

// Dial connect to stratum server
func (c *Client) Dial() (err error) {
	remoteAddr, err := net.ResolveTCPAddr("tcp", c.server.Host)
	if err != nil {
		return c.dialNextServer(err)
	}

	dialer := &net.Dialer{
		Timeout: connectTimeout * time.Second,
	}

	if c.tls {
		c.conn, err = tls.DialWithDialer(dialer, "tcp", c.server.Host, &tls.Config{InsecureSkipVerify: false})
	} else {
		c.conn, err = dialer.Dial("tcp", remoteAddr.String())
	}

	if err != nil {
		return c.dialNextServer(err)
	}

	atomic.StoreInt32(&c.closed, 0)
	atomic.StoreUint64(&c.ID, 0)

	if !c.tls {
		c.conn.(*net.TCPConn).SetKeepAlive(true)
		c.conn.(*net.TCPConn).SetNoDelay(true)
	}

	c.reader = bufio.NewReaderSize(c.conn, bufferSize)
	c.writer = bufio.NewWriterSize(c.conn, bufferSize)

	switch c.server.Scheme {
	case "stratum+tcp":
		if strings.Contains(c.server.Host, "nicehash.com") {
			c.mode = etherstratum
			break
		}
		c.mode = stratum
	default:
		c.mode = etherproxy
	}

	log.Info("Connected to stratum server", "stratum", c.server)

	go c.Read()

	return c.Subscribe()
}

// Read from stratum server
func (c *Client) Read() {

	for !c.isClosed() {
		c.conn.SetReadDeadline(time.Now().Add(c.timeout))

		data, err := c.reader.ReadString('\n')
		if err != nil {
			if c.noReconnect {
				return
			}

			if err.Error() == "EOF" {
				err = errors.New("Connection disconnect detected by EOF")
			}

			log.Error(err.Error(), "stratum", c.server)
			c.reconnect()
			return
		}

		if len(data) < 4 {
			continue
		}

		if !c.tls {
			log.Trace("Received stratum message", "message", strings.TrimSuffix(string(data), "\n"))
		}

		var resp Resp
		err = json.Unmarshal([]byte(data), &resp)
		if err != nil {
			log.Warn(err.Error(), "stratum", c.server)
			c.markSick()
			continue
		}

		c.handleStratumMessage(&resp)
	}
}

// Write to stratum server
func (c *Client) Write(data []byte) error {
	if c.isClosed() {
		return errors.New("Connection closed")
	}

	if !c.tls {
		log.Trace("Sending message to stratum", "message", string(data))
	}

	c.conn.SetWriteDeadline(time.Now().Add(c.timeout))

	_, err := c.writer.Write(data)
	if err != nil {
		return err
	}

	_, err = c.writer.WriteString("\n")
	if err != nil {
		return err
	}

	return c.writer.Flush()
}

// Subscribe on stratum server
func (c *Client) Subscribe() error {
	if c.mode == etherproxy {
		return c.Auth()
	}

	params := []string{"eminer/" + c.minerVersion, "EthereumStratum/1.0.0"}

	var result []*json.RawMessage
	resp, err := c.request("mining.subscribe", params)
	if err != nil {
		return err
	}

	if resp.Error != nil {
		return fmt.Errorf("%v", resp.Error)
	}

	if resp.Result == nil {
		return errors.New("No response found")
	}

	err = json.Unmarshal(*resp.Result, &result)
	if err != nil {
		return err
	}

	if len(result) <= 0 {
		return errors.New("Subscribe failed")
	}

	if len(result) == 2 {
		var subscribeResult []string
		err = json.Unmarshal(*result[0], &subscribeResult)
		if err != nil {
			return err
		}

		if len(subscribeResult) == 3 && subscribeResult[2] == "EthereumStratum/1.0.0" {
			c.mode = etherstratum
		}
	}

	if c.mode == etherstratum {
		var extraNonce string
		err = json.Unmarshal(*result[1], &extraNonce)
		if err != nil {
			return err
		}

		c.work.extraNonceSize = len(extraNonce)

		for i := len(extraNonce); i < 16; i++ {
			extraNonce = extraNonce + "0"
		}

		c.work.extraNonce = new(big.Int).SetBytes(common.FromHex(extraNonce))

		go c.request("mining.extranonce.subscribe", []string{""})

		time.Sleep(500 * time.Millisecond) //give a time to socket, sometimes received "short write" error
	}

	return c.Auth()
}

// Auth on stratum server
func (c *Client) Auth() error {
	method := "eth_submitLogin"

	if c.mode != etherproxy {
		method = "mining.authorize"
	}

	var auth bool
	resp, err := c.request(method, []string{c.username, c.password})
	if err != nil {
		return err
	}

	if resp.Error != nil {
		return fmt.Errorf("%v", resp.Error)
	}

	if resp.Result == nil {
		return errors.New("No response found")
	}

	err = json.Unmarshal(*resp.Result, &auth)
	if err != nil {
		return err
	}

	if !auth {
		return errors.New("Authentication failed")
	}

	return nil
}

// GetWork from stratum server
func (c *Client) GetWork() ([]string, error) {
	if c.mode != etherproxy {
		return c.waitWorkFromStratum()
	}

	var getWork []string
	resp, err := c.request("eth_getWork", []string{})
	if err != nil {
		return nil, err
	}

	if resp.Error != nil {
		return getWork, fmt.Errorf("%v", resp.Error)
	}

	if resp.Result == nil {
		return getWork, errors.New("No response found")
	}

	err = json.Unmarshal(*resp.Result, &getWork)
	return getWork, err
}

// SubmitHashrate to stratum server
func (c *Client) SubmitHashrate(params interface{}) (bool, error) {
	if c.mode != etherproxy {
		return true, nil
	}

	resp, err := c.request("eth_submitHashrate", params)
	if err != nil {
		return false, nil
	}

	if resp.Error != nil {
		return false, fmt.Errorf("%v", resp.Error)
	}

	if resp.Result == nil {
		return false, nil
	}

	var ok bool
	err = json.Unmarshal(*resp.Result, &ok)
	if err != nil {
		return false, nil
	}

	return ok, nil
}

// SubmitWork to stratum server
func (c *Client) SubmitWork(params interface{}) (bool, error) {
	method := "eth_submitWork"
	var solution []interface{}

	if c.mode == stratum {
		method = "mining.submit"
		solution = append(solution, c.username)
		solution = append(solution, c.work.jobid)

		for _, param := range params.([]interface{}) {
			solution = append(solution, param)
		}
	} else if c.mode == etherstratum {
		method = "mining.submit"
		solution = append(solution, c.username)
		solution = append(solution, c.work.jobid)

		blockNonce := params.([]interface{})[0].(types.BlockNonce)
		nonce := common.Bytes2Hex(blockNonce[:])

		log.Trace("Submitting nonce with extranonce", "nonce", nonce, "extraNonce", common.Bytes2Hex(c.work.extraNonce.Bytes()))

		solution = append(solution, nonce[c.work.extraNonceSize:])
	} else {
		for _, param := range params.([]interface{}) {
			solution = append(solution, param)
		}
	}

	resp, err := c.request(method, solution)
	if err != nil {
		log.Warn("Submit work error", "stratum", c.server, "error", err.Error())
		return true, nil // network error, never rejected
	}

	if resp.Error != nil {
		return false, fmt.Errorf("%v", resp.Error) // rejected with error result
	}

	if resp.Result == nil {
		log.Warn("Submit work error", "stratum", c.server, "error", "No response found")
		return true, nil // like stratum internal server error
	}

	var ok bool
	err = json.Unmarshal(*resp.Result, &ok)
	if err != nil {
		log.Warn("Submit work error", "stratum", c.server, "error", err.Error())
		return true, nil // json error, never rejected
	}

	if !ok {
		// rejected without error result
		return ok, errors.New("Rejected without reason")
	}

	return ok, nil
}

// RegisterNotification for notifications
func (c *Client) RegisterNotification(method string, callback chan *json.RawMessage) {
	c.Lock()
	defer c.Unlock()

	c.notificationCh[method] = callback
}

// UnregisterNotification for notifications
func (c *Client) UnregisterNotification(method string) {
	c.Lock()
	defer c.Unlock()

	delete(c.notificationCh, method)
}

//Sick func
func (c *Client) Sick() bool {
	return c.sick
}

// Close the connection
func (c *Client) Close(noReconnect bool) {
	c.noReconnect = noReconnect

	if !c.isClosed() && c.conn != nil {
		atomic.StoreInt32(&c.closed, 1)
		c.conn.Close()
	}
}

func (c *Client) request(method string, params interface{}) (*Resp, error) {
	req := Req{
		ID:     atomic.AddUint64(&c.ID, 1),
		Method: method,
		Params: params,
	}

	if c.mode == etherproxy {
		req.Worker = c.worker
	}

	data, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	return c.send(req.ID, method, data)

}

func (c *Client) send(id uint64, method string, data []byte) (*Resp, error) {
	c.Lock()
	respCallback := make(chan *Resp)
	c.callBackCh[id] = &Callback{method: method, resp: respCallback}
	c.Unlock()

	defer c.deleteCallback(id)

	err := c.Write(data)
	if err != nil {
		return nil, err
	}

	select {
	case resp := <-respCallback:
		if resp != nil {
			c.markAlive()
			return resp, nil
		}
		return nil, errors.New("Request cancelled")
	case <-time.After(responseTimeout * time.Second):
		c.markSick()
		return nil, errors.New("Timed out, no response from server")
	}
}

func (c *Client) deleteCallback(id uint64) {
	c.Lock()
	defer c.Unlock()

	delete(c.callBackCh, id)
}

func (c *Client) reconnect() {
	if c.noReconnect {
		return
	}

	c.Close(false)

	time.Sleep(connectRetryTime * time.Second)

	c.Lock()
	for id, callback := range c.callBackCh {
		delete(c.callBackCh, id)
		close(callback.resp)
	}
	c.Unlock()

	log.Info("Reconnecting to stratum server", "stratum", c.server)

	if c.Sick() {
		if len(c.backupServers) > c.serverIdx+1 {
			log.Warn("Stratum server sick", "stratum", c.server)

			c.server = c.backupServers[c.serverIdx+1]
			c.serverIdx++

			if c.serverIdx+1 == len(c.backupServers) {
				c.serverIdx = -1
			}

			log.Info("Trying another server", "stratum", c.server)
		}
	}

	c.sick = false
	c.sickRate = 0
	c.successRate = 0

	err := c.Dial()
	if err != nil {
		log.Error(err.Error(), "stratum", c.server)
		c.reconnect()
		return
	}
}

func (c *Client) waitWorkFromStratum() ([]string, error) {
	if c.work.jobid != "" {
		return c.workToStringSlice(), nil
	}

	for c.work.jobid == "" {
		time.Sleep(1 * time.Second)
	}

	return c.workToStringSlice(), nil
}

func (c *Client) markSick() {
	if !c.sick {
		atomic.AddUint64(&c.FailsCount, 1)
	}
	c.sickRate++
	c.successRate = 0
	if c.sickRate >= 5 {
		c.sick = true
	}
}

func (c *Client) markAlive() {
	c.successRate++
	if c.successRate >= 5 {
		c.sick = false
		c.sickRate = 0
		c.successRate = 0
	}
}

func (c *Client) dialNextServer(err error) error {
	if len(c.backupServers) > c.serverIdx+1 {
		log.Warn("Stratum server sick", "stratum", c.server, "error", err.Error())

		c.server = c.backupServers[c.serverIdx+1]
		c.serverIdx++

		if c.serverIdx+1 == len(c.backupServers) {
			c.serverIdx = -1
		}

		log.Info("Trying another server", "stratum", c.server)

		time.Sleep(2 * time.Second)

		return c.Dial()
	}

	return err
}

func (c *Client) searchWaitingResponses(resp *Resp) {
	for _, callback := range c.callBackCh {
		if callback.method == "eth_getWork" {
			callback.resp <- resp
		}
	}
}

func (c *Client) handleNotification(resp *Resp) {
	c.Lock()
	defer c.Unlock()

	switch resp.Method {
	case "mining.notify":
		var notify []interface{}
		err := json.Unmarshal(*resp.Params, &notify)
		if err != nil {
			log.Error("Stratum mining notify json decode error", "error", err.Error())
			return
		}

		if len(notify) < 4 {
			return
		}

		if c.mode == etherstratum && len(notify) == 4 { //etherstratum mode
			c.work.jobid = notify[0].(string)      //jobid
			c.work.seedHash = notify[2].(string)   //seedhash
			c.work.headerHash = notify[1].(string) //headerhash
			if c.work.target == nil {
				c.work.target = difficultyToTarget(1)
			}
		}

		if c.mode == stratum && len(notify) == 5 { //stratum mode
			c.work.jobid = notify[0].(string)                                         //jobid
			c.work.seedHash = notify[1].(string)                                      //seedhash
			c.work.headerHash = notify[2].(string)                                    //headerhash
			c.work.target = new(big.Int).SetBytes(common.FromHex(notify[3].(string))) //target
		}

		c.callWorkCh()
		return
	case "mining.set_difficulty":
		var result []float64
		err := json.Unmarshal(*resp.Params, &result)
		if err != nil {
			log.Error("Stratum mining set_difficulty json decode error", "error", err.Error())
			return
		}

		c.work.target = difficultyToTarget(result[0])

		c.callWorkCh()
		return
	case "mining.set_extranonce":
		var result []string
		err := json.Unmarshal(*resp.Params, &result)
		if err != nil {
			log.Error("Stratum mining set_extranonce json decode error", "error", err.Error())
			return
		}

		extraNonce := result[0]
		c.work.extraNonceSize = len(extraNonce)

		for i := len(extraNonce); i < 16; i++ {
			extraNonce = extraNonce + "0"
		}

		c.work.extraNonce = new(big.Int).SetBytes(common.FromHex(extraNonce))

		c.callWorkCh()
		return
	default:
		if callback, ok := c.notificationCh["notify_Work"]; ok {
			c.markAlive()
			callback <- resp.Result
			return
		}
	}

	//Some pools send getWork response with wrong id
	c.searchWaitingResponses(resp)
}

func (c *Client) handleStratumMessage(resp *Resp) {
	if resp.ID == 0 {
		c.handleNotification(resp)
		return
	}

	c.Lock()
	defer c.Unlock()

	if callBack, ok := c.callBackCh[resp.ID]; ok {
		callBack.resp <- resp
	}
}

func (c *Client) isClosed() bool {
	return atomic.LoadInt32(&c.closed) == 1
}

func (c *Client) callWorkCh() {
	if callback, ok := c.notificationCh["notify_Work"]; ok {
		params, err := json.Marshal(c.workToStringSlice())
		if err != nil {
			log.Error("Stratum work json encode error", "error", err.Error())
			return
		}
		paramsRaw := json.RawMessage(params)
		c.markAlive()
		callback <- &paramsRaw
	}
}

func (c *Client) workToStringSlice() (work []string) {
	work = append(work, c.work.seedHash)
	work = append(work, c.work.headerHash)
	work = append(work, common.ToHex(c.work.target.Bytes()))

	if c.work.extraNonce != nil {
		work = append(work, common.ToHex(c.work.extraNonce.Bytes()))
		work = append(work, strconv.Itoa(c.work.extraNonceSize*4)) //sizebits
	}

	return
}

func difficultyToTarget(diff float64) (target *big.Int) {
	mantissa := 0x0000ffff / diff
	exp := 1
	tmp := mantissa
	for tmp >= 256.0 {
		tmp /= 256.0
		exp++
	}
	for i := 0; i < exp; i++ {
		mantissa *= 256.0
	}
	target = new(big.Int).Lsh(big.NewInt(int64(mantissa)), uint(26-exp)*8)
	return
}
