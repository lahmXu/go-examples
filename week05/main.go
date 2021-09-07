package main

import "metric/metric"

func main() {
	d := metric.NewDefaultMetricCollector()
	d.Update(metric.Result{
		Attempts:      0,
		Errors:        0,
		Successes:     0,
		TotalDuration: 0,
		RunDuration:   0,
	})
}
