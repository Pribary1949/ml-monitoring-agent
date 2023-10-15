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
