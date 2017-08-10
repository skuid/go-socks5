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
		Name: "num_requests_success",
		Help: "Number of requests that failed to be handled properly",
	})

	promRequestFailed = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "num_requests_failed",
		Help: "Number of requests that were successfully fulfilled",
	})

	promRequestDuration = prometheus.NewSummaryVec(prometheus.SummaryOpts{
		Name:   "tcp_request_duration_microseconds",
		Help:   "Response duration summary in microseconds.",
		MaxAge: time.Hour,
	}, []string{"remote_addr"})

	promRequestLatency = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name: "tcp_request_latencies",
			Help: "Response latency distribution in microseconds for request",
			// Use buckets ranging from 125 ms to 8 seconds.
			Buckets: prometheus.ExponentialBuckets(125000, 2.0, 7),
		},
		[]string{"remote_addr"},
	)

	// Request Setup
	promSetupRequestDuration = prometheus.NewSummaryVec(prometheus.SummaryOpts{
		Name:   "tcp_setup_request_duration_microseconds",
		Help:   "Response setup duration summary in microseconds.",
		MaxAge: time.Hour,
	}, []string{"remote_addr"})

	promSetupRequestLatency = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name: "tcp_setup_request_latencies",
			Help: "Response latency distribution in microseconds for request",
			// Use buckets ranging from 125 ms to 8 seconds.
			Buckets: prometheus.ExponentialBuckets(125000, 2.0, 7),
		},
		[]string{"remote_addr"},
	)
)

func instrumentRequestDuration(startTime time.Time, remoteAddr string) {
	elapsed := float64((time.Since(startTime)) / time.Microsecond)
	promRequestDuration.WithLabelValues(remoteAddr).Observe(elapsed)
	promRequestLatency.WithLabelValues(remoteAddr).Observe(elapsed)
}

func instrumentRequestSetupDuration(startTime time.Time, remoteAddr string) {
	elapsed := float64((time.Since(startTime)) / time.Microsecond)
	promSetupRequestDuration.WithLabelValues(remoteAddr).Observe(elapsed)
	promSetupRequestLatency.WithLabelValues(remoteAddr).Observe(elapsed)
}

func init() {
	prometheus.MustRegister(promRequestCounter)
	prometheus.MustRegister(promRequestSuccess)
	prometheus.MustRegister(promRequestFailed)
	prometheus.MustRegister(promRequestDuration)
	prometheus.MustRegister(promRequestLatency)
}
