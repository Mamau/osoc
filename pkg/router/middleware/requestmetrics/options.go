package requestmetrics

import "github.com/prometheus/client_golang/prometheus"

type Option func(*options)

type options struct {
	requestCounter  *prometheus.CounterVec
	requestDuration *prometheus.HistogramVec
}

func RequestCounter(totalRequests *prometheus.CounterVec) Option {
	return func(o *options) {
		o.requestCounter = totalRequests
	}
}

func RequestDuration(requestLatency *prometheus.HistogramVec) Option {
	return func(o *options) {
		o.requestDuration = requestLatency
	}
}
