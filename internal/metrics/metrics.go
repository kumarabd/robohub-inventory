package metrics

import (
	"time"
)

// Metrics tracks application metrics
type Metrics struct {
	RequestCount    int64
	RequestDuration time.Duration
	ErrorCount      int64
}

var globalMetrics *Metrics

func Init() {
	globalMetrics = &Metrics{}
}

func Get() *Metrics {
	if globalMetrics == nil {
		Init()
	}
	return globalMetrics
}

func (m *Metrics) IncrementRequestCount() {
	m.RequestCount++
}

func (m *Metrics) IncrementErrorCount() {
	m.ErrorCount++
}

func (m *Metrics) AddRequestDuration(duration time.Duration) {
	m.RequestDuration += duration
}

func (m *Metrics) Reset() {
	m.RequestCount = 0
	m.RequestDuration = 0
	m.ErrorCount = 0
}
