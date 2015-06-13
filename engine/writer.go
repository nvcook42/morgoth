package engine

import (
	metric "github.com/nathanielc/morgoth/metric/types"
	"time"
)

type Writer interface {
	Insert(datetime time.Time, metric metric.MetricID, value float64)
	RecordAnomalous(metric metric.MetricID, start, stop time.Time)
	DeleteMetric(metric metric.MetricID)
}
