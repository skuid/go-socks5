package socks5

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	promRequestCounter = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "num_requests_handled",
		Help: "Number of requests handled by the proxy",
	})

	promRequestSuccess = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "num_requests_failed",
		Help: "Number of requests that failed to be handled properly",
	})

	promRequestFailed = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "num_requests_success",
		Help: "Number of requests that were successfully fulfilled",
	})

	promRequestDuration = prometheus.NewSummaryVec(prometheus.SummaryOpts{
		Name:   "tcp_request_duration_milliseconds",
		Help:   "Response duration summary in milliseconds.",
		MaxAge: time.Hour,
	}, []string{"request"})

	promRequestLatency = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name: "tcp_request_latencies",
			Help: "Response latency distribution in microseconds for each verb and path",
			// Use buckets ranging from 125 ms to 8 seconds.
			Buckets: prometheus.ExponentialBuckets(125000, 2.0, 7),
		},
		[]string{"request"},
	)
)

func instrumentRequestDuration(startTime time.Time) {
	elapsed := float64((time.Since(startTime)) / time.Microsecond)
	promRequestDuration.WithLabelValues("request").Observe(elapsed)
	promRequestLatency.WithLabelValues("request").Observe(elapsed)
}

func init() {
	prometheus.MustRegister(promRequestCounter)
	prometheus.MustRegister(promRequestSuccess)
	prometheus.MustRegister(promRequestFailed)
	prometheus.MustRegister(promRequestDuration)
	prometheus.MustRegister(promRequestLatency)
}
