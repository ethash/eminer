package http

import (
	"bytes"
	"html/template"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/ethash/eminer/ethash"
	"github.com/ethash/eminer/http/metricstat"
	"github.com/ethereum/go-ethereum/log"
	"github.com/gin-gonic/gin"
	metrics "github.com/rcrowley/go-metrics"
)

// Server struct
type Server struct {
	openclMiner *ethash.OpenCLMiner
	metricstat  *metricstat.MetricStat
	engine      *gin.Engine
	listen      string
	workerName  string
	version     string
	debug       bool
}

// New server
func New(listen string, version, workerName string, debug bool) *Server {
	server := &Server{
		listen:     listen,
		engine:     gin.Default(),
		metricstat: metricstat.New(metrics.DefaultRegistry, 24*time.Hour, time.Minute),
		workerName: workerName,
		version:    version,
		debug:      debug,
	}

	return server
}

// ClearStats clear metric stats
func (s *Server) ClearStats() {
	s.metricstat.Clear()
}

// SetMiner for stats
func (s *Server) SetMiner(miner *ethash.OpenCLMiner) {
	s.openclMiner = miner
}

func (s *Server) showDash(c *gin.Context) {
	//TODO: using while dev
	//dashboard, err := template.New("dashboard.html").Delims("[[", "]]").ParseFiles("http/templates/dashboard.html")

	dashboardTemplate, err := templatesDashboardHtmlBytes()
	if err != nil {
		c.String(http.StatusOK, "template execution: %s", err)
		return
	}
	dashboard, _ := template.New("dashboard.html").Delims("[[", "]]").Parse(string(dashboardTemplate))

	data := make(map[string]interface{})

	data["version"] = s.version
	data["worker_name"] = s.workerName

	var doc bytes.Buffer

	err = dashboard.Execute(&doc, data)
	if err != nil {
		c.String(http.StatusOK, "template execution: %s", err)
		return
	}

	c.Writer.Header().Set("Content-Type", "text/html")
	c.String(http.StatusOK, doc.String())
}

func (s *Server) chartData(c *gin.Context) {
	c.Writer.Header().Add("Access-Control-Allow-Origin", "*")

	data := make(map[string][][2]interface{})
	var totalHashRate [][2]interface{}

	for name, buckets := range s.metricstat.FromDuration(24 * time.Hour) {
		var lastValue float64

		for i, b := range buckets {
			var item [2]interface{}
			item[0] = b.T.Unix() * 1000
			item[1] = b.V

			if strings.HasPrefix(name, s.workerName+".solutions") {
				item[1] = item[1].(float64) - lastValue
				lastValue = b.V

				if i == 0 {
					item[1] = 0
				}
			}

			if strings.HasSuffix(name, "hashrate") {
				item[1] = b.V / 1e6
				if len(totalHashRate) > i {
					totalHashRate[i][1] = totalHashRate[i][1].(float64) + b.V/1e6
				} else {
					totalHashRate = append(totalHashRate, item)
				}
			}

			data[name] = append(data[name], item)
		}
	}

	if len(totalHashRate) > 0 {
		data[s.workerName+".total.hashrate"] = totalHashRate
	}

	if c.Query("series") != "" {
		series := strings.Split(c.Query("series"), ",")
		tempData := make(map[string][][2]interface{})

		for _, s := range series {
			tempData[s] = data[s]
		}

		c.JSON(200, tempData)
		return
	}

	c.JSON(200, data)
}

func (s *Server) getStats(c *gin.Context) {
	c.Writer.Header().Add("Access-Control-Allow-Origin", "*")

	if s.openclMiner == nil {
		c.JSON(200, gin.H{"error": "Miner is not ready"})
		return
	}

	c.JSON(200, s.openclMiner)
}

func init() {
	devNull, _ := os.Open(os.DevNull)
	gin.DefaultWriter = devNull
	gin.SetMode(gin.ReleaseMode)
}

// Start http server
func (s *Server) Start() {

	s.engine.GET("/", s.showDash)

	v1 := s.engine.Group("/api/v1")
	{
		v1.GET("/stats", s.getStats)
		v1.GET("/chartData", s.chartData)
	}

	if s.debug {
		pprofRoutes(s.engine)
	}

	srv := &http.Server{
		Addr:         s.listen,
		Handler:      s.engine,
		ReadTimeout:  60 * time.Second,
		WriteTimeout: 60 * time.Second,
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Error("http listen error", "error", err.Error())
	}
}
