package rpc

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/ethereum/go-ethereum/log"
)

//Client struct
type Client struct {
	sync.RWMutex

	URL        *url.URL
	backupURLs []*url.URL
	urlIdx     int

	client *http.Client

	sick        bool
	sickRate    int
	successRate int
	FailsCount  uint64

	timeout time.Duration
}

//GetBlockReply struct
type GetBlockReply struct {
	Number     string `json:"number"`
	Difficulty string `json:"difficulty"`
}

//JSONRpcResp struct
type JSONRpcResp struct {
	ID     *json.RawMessage       `json:"id"`
	Result *json.RawMessage       `json:"result"`
	Error  map[string]interface{} `json:"error"`
}

// New func
func New(rawURLs string, timeout time.Duration) (*Client, error) {
	c := new(Client)

	splitURLs := strings.Split(rawURLs, ",")

	if len(splitURLs) > 0 {
		for _, rawURL := range splitURLs {
			url, err := url.Parse(rawURL)
			if err != nil {
				log.Error("Error parse url", "url", url, "error", err.Error())
			}

			c.backupURLs = append(c.backupURLs, url)
		}
	}

	if len(c.backupURLs) == 0 {
		return nil, errors.New("No URL found")
	}

	c.URL = c.backupURLs[c.urlIdx]
	c.timeout = timeout

	c.client = &http.Client{
		Timeout: c.timeout,
	}

	return c, nil
}

//GetWork func
func (r *Client) GetWork() ([]string, error) {
	rpcResp, err := r.doPost("eth_getWork", []string{}, 73)
	var reply []string
	if err != nil {
		return reply, err
	}
	if rpcResp.Error != nil {
		return reply, errors.New("[JSON-RPC] " + rpcResp.Error["message"].(string))
	}
	err = json.Unmarshal(*rpcResp.Result, &reply)
	return reply, err
}

//SubmitWork func
func (r *Client) SubmitWork(params interface{}) (bool, error) {
	rpcResp, err := r.doPost("eth_submitWork", params, 73)
	var result bool
	if err != nil {
		log.Warn("[JSON-RPC] Submit work error", "error", err.Error())
		return true, nil //network error never rejected
	}

	if rpcResp.Error != nil {
		return false, errors.New("[JSON-RPC] " + rpcResp.Error["message"].(string))
	}

	json.Unmarshal(*rpcResp.Result, &result)
	if !result {
		return false, errors.New("[JSON-RPC] " + "rejected without reason")
	}
	return result, nil
}

//SubmitHashrate func
func (r *Client) SubmitHashrate(params interface{}) (bool, error) {
	rpcResp, err := r.doPost("eth_submitHashrate", params, 73)
	var result bool
	if err != nil {
		return false, nil
	}
	if rpcResp.Error != nil {
		return false, nil
	}

	json.Unmarshal(*rpcResp.Result, &result)

	return result, nil
}

func (r *Client) doPost(method string, params interface{}, id uint64) (JSONRpcResp, error) {
	if r.Sick() && len(r.backupURLs) > r.urlIdx+1 {
		log.Warn("RPC server sick", "url", r.URL.String())

		r.URL = r.backupURLs[r.urlIdx+1]
		r.urlIdx++

		if r.urlIdx+1 == len(r.backupURLs) {
			r.urlIdx = -1
		}

		log.Info("Trying another RPC server", "url", r.URL.String())

		//clear stats
		r.Lock()
		r.sick = false
		r.sickRate = 0
		r.successRate = 0
		r.Unlock()
	}

	var rpcResp JSONRpcResp

	jsonReq := map[string]interface{}{"jsonrpc": "2.0", "id": id, "method": method, "params": params}

	data, _ := json.Marshal(jsonReq)
	req, err := http.NewRequest("POST", r.URL.String(), bytes.NewBuffer(data))
	if err != nil {
		r.markSick()
		return rpcResp, errors.New("[JSON-RPC] " + err.Error())
	}

	req.Header.Set("Content-Length", (string)(len(data)))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	resp, err := r.client.Do(req)

	if err != nil {
		r.markSick()
		return rpcResp, errors.New("[JSON-RPC] " + err.Error())
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, &rpcResp)

	if rpcResp.Error != nil {
		r.markSick()
	}

	if err != nil {
		return rpcResp, errors.New("[JSON-RPC] " + err.Error())
	}

	return rpcResp, nil
}

//Check func
func (r *Client) Check() (bool, error) {
	_, err := r.GetWork()
	if err != nil {
		return false, errors.New("[JSON-RPC] " + err.Error())
	}
	r.markAlive()
	return !r.Sick(), nil
}

//Sick func
func (r *Client) Sick() bool {
	r.RLock()
	defer r.RUnlock()
	return r.sick
}

func (r *Client) markSick() {
	r.Lock()
	if !r.sick {
		atomic.AddUint64(&r.FailsCount, 1)
	}
	r.sickRate++
	r.successRate = 0
	if r.sickRate >= 5 {
		r.sick = true
	}
	r.Unlock()
}

func (r *Client) markAlive() {
	r.Lock()
	r.successRate++
	if r.successRate >= 5 {
		r.sick = false
		r.sickRate = 0
		r.successRate = 0
	}
	r.Unlock()
}
