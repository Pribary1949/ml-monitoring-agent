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
		}),
		errorRate: promauto.NewSummary(prometheus.SummaryOpts{
			Namespace: namespace,
			Name:      "inference_errors",
			Help:      "Summary of inference errors",
		}),
	}
}

func (c *MetricsCollector) RecordInference(duration time.Duration, success bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.requestCount.Inc()
	c.inferenceTime.Observe(duration.Seconds())
	if !success {
		c.errorRate.Observe(1.0)
	}
}

func (c *MetricsCollector) UpdateDrift(score float64) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.driftScore.Set(score)
}

func (c *MetricsCollector) StartSimulation() {
	go func() {
		for {
			time.Sleep(2 * time.Second)
			c.RecordInference(150*time.Millisecond, true)
			c.UpdateDrift(0.05)
		}
	}()
}
