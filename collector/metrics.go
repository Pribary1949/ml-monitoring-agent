package collector

import (
	"sync"
	"time"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type MetricsCollector struct {
	mu            sync.Mutex
	inferenceTime prometheus.Histogram
	requestCount  prometheus.Counter
	driftScore    prometheus.Gauge
	errorRate     prometheus.Summary
}

func NewMetricsCollector(namespace string) *MetricsCollector {
	return &MetricsCollector{
		inferenceTime: promauto.NewHistogram(prometheus.HistogramOpts{
			Namespace: namespace,
			Name:      "inference_duration_seconds",
			Help:      "Latency of model inference in seconds",
			Buckets:   prometheus.DefBuckets,
		}),
		requestCount: promauto.NewCounter(prometheus.CounterOpts{
			Namespace: namespace,
			Name:      "total_requests",
			Help:      "Total number of inference requests",
		}),
		driftScore: promauto.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Name:      "model_drift_score",
			Help:      "Current calculated drift score (0-1)",