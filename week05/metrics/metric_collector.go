package metrics

import (
	"sync"
	"time"
)

// MetricCollector 接口
type MetricCollector interface {
	Update(Result)
	Reset()
}

type DefaultMetricCollector struct {
	mutex *sync.RWMutex

	numRequests *Number
	errors      *Number
	successes   *Number

	totalDuration *Timing
	runDuration   *Timing
}

// Result result
type Result struct {
	Attempts      float64
	Errors        float64
	Successes     float64
	TotalDuration time.Duration
	RunDuration   time.Duration
}

func NewDefaultCollector() MetricCollector {
	m := &DefaultMetricCollector{}
	m.mutex = &sync.RWMutex{}
	m.Reset()
	return m
}

func (d *DefaultMetricCollector) NumRequests() *Number {
	d.mutex.RLock()
	defer d.mutex.RUnlock()
	return d.numRequests
}

func (d *DefaultMetricCollector) Errors() *Number {
	d.mutex.RLock()
	defer d.mutex.RUnlock()
	return d.errors
}

func (d *DefaultMetricCollector) Successes() *Number {
	d.mutex.RLock()
	defer d.mutex.RUnlock()
	return d.errors
}

func (d *DefaultMetricCollector) TotalDurations() *Timing {
	d.mutex.RLock()
	defer d.mutex.RUnlock()
	return d.totalDuration
}

func (d *DefaultMetricCollector) RunDuration() *Timing {
	d.mutex.RLock()
	defer d.mutex.RUnlock()
	return d.runDuration
}

func (d *DefaultMetricCollector) Update(r Result) {
	d.mutex.RLock()
	defer d.mutex.RUnlock()

	d.numRequests.Increment(r.Attempts)
	d.errors.Increment(r.Errors)
	d.successes.Increment(r.Successes)
	d.totalDuration.Add(r.TotalDuration)
	d.runDuration.Add(r.RunDuration)
}

func (d *DefaultMetricCollector) Reset() {
	d.mutex.Lock()
	defer d.mutex.Unlock()

	d.numRequests = NewNumber()
	d.errors = NewNumber()
	d.successes = NewNumber()
	d.totalDuration = NewTiming()
	d.runDuration = NewTiming()
}
