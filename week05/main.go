package main

import "metric/metrics"

func main() {
	d := metrics.NewDefaultCollector()
	d.Update(metrics.Result{
		Attempts:      0,
		Errors:        0,
		Successes:     0,
		TotalDuration: 0,
		RunDuration:   0,
	})
}
