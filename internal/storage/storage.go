package storage

import (
	"github.com/MlDenis/dm-go-musthave-metrics/internal/metric"
)

type MetricData struct {
	metricType   string
	gaugeValue   metric.Gauge
	counterValue metric.Counter
}

type MemStorage struct {
	data map[string]*MetricData
}

func NewMetricsStorage() MemStorage {
	return MemStorage{make(map[string]*MetricData)}
}

func (ms *MemStorage) UpdateMetricInStorage(
	metricType string,
	metricName string,
	gaugeValue metric.Gauge,
	counterValue metric.Counter) {

	_, ok := ms.data[metricName]
	if ok {
		switch metricType {
		case metric.GaugeString:
			(ms.data[metricName]).gaugeValue = gaugeValue
		case metric.CounterString:
			ms.data[metricName].counterValue += counterValue
		}
	} else {
		newElement := new(MetricData)
		newElement.metricType = metricType
		newElement.gaugeValue = gaugeValue
		newElement.counterValue = counterValue

		ms.data[metricName] = newElement
	}
}
