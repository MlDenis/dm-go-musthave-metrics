package metric

type Gauge float64
type Counter int64

const (
	GaugeString   = "gauge"
	CounterString = "counter"
)
