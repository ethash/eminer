package metricstat

import (
	"sync"
	"time"

	metrics "github.com/rcrowley/go-metrics"
	"github.com/ethash/eminer/http/ts"
)

// MetricStat structure
type MetricStat struct {
	sync.RWMutex

	registry metrics.Registry
	duration time.Duration
	timespan time.Duration

	metricseries map[string]*ts.Series
}

// New return metricstat
func New(r metrics.Registry, t, d time.Duration) *MetricStat {
	m := &MetricStat{
		registry:     r,
		duration:     d,
		timespan:     t,
		metricseries: make(map[string]*ts.Series),
	}

	go m.start()

	return m
}

func (m *MetricStat) start() {
	for {
		m.update()

		time.Sleep(m.duration)
	}
}

// FromDuration returns series range of buckets
func (m *MetricStat) FromDuration(d time.Duration) map[string][]*ts.Bucket {
	m.RLock()
	defer m.RUnlock()

	buckets := make(map[string][]*ts.Bucket)

	for name, series := range m.metricseries {
		buckets[name] = series.FromDuration(d)
	}

	return buckets
}

// Clear metric stats
func (m *MetricStat) Clear() {
	m.Lock()
	defer m.Unlock()

	for name := range m.metricseries {
		m.metricseries[name] = ts.New(m.timespan, m.duration)
	}
}

func (m *MetricStat) update() {
	m.Lock()
	defer m.Unlock()

	m.registry.Each(func(name string, i interface{}) {
		var series *ts.Series
		var ok bool

		if series, ok = m.metricseries[name]; !ok {
			series = ts.New(m.timespan, m.duration)
			m.metricseries[name] = series
		}

		switch metric := i.(type) {
		case metrics.Counter:
			series.Insert(time.Now(), float64(metric.Count()))
		case metrics.Histogram:
			series.Insert(time.Now(), float64(metric.Count()))
		case metrics.Gauge:
			series.Insert(time.Now(), float64(metric.Value()))
		case metrics.GaugeFloat64:
			series.Insert(time.Now(), metric.Value())
		case metrics.Meter:
			snapshot := metric.Snapshot()
			series.Insert(time.Now(), snapshot.Rate1())
		}
	})
}
