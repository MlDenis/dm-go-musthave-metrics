package storage

import (
	"github.com/MlDenis/dm-go-musthave-metrics/internal/metric"
	"log"
)

type MetricData struct {
	MetricType   string
	GaugeValue   metric.Gauge
	CounterValue metric.Counter
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
	counterValue metric.Counter,
) {
	_, ok := ms.data[metricName]
	if ok {
		switch metricType {
		case metric.GaugeString:
			(ms.data[metricName]).GaugeValue = gaugeValue
		case metric.CounterString:
			ms.data[metricName].CounterValue += counterValue
		}
	} else {
		newElement := MetricData{
			metricType,
			gaugeValue,
			counterValue,
		}

		ms.data[metricName] = &newElement
	}

	log.Printf("#DEBUG UpdateMetricInStorage sucessfully complete with: %+v", ms.data[metricName])
}
