package storage

import (
	"github.com/MlDenis/dm-go-musthave-metrics/internal/metric"
	"log"
)

type MetricData struct {
	metricType   string
	gaugeValue   metric.Gauge
	counterValue metric.Counter
}

type MemStorage struct {
	data map[string]*MetricData
}

func (ms *MemStorage) NewMetricsStorage() MemStorage {
	return MemStorage{}
}

func (ms *MemStorage) UpdateMetricInStorage(
	metricType string,
	metricName string,
	gaugeValue metric.Gauge,
	counterValue metric.Counter) {
	switch metricType {
	case metric.GaugeString:
		(ms.data[metricName]).gaugeValue = gaugeValue
		log.Printf("#DEBUG: gaugeValue %s : %v has been added to the MemStorage\n", metricName, gaugeValue)
	case metric.CounterString:
		ms.data[metricName].counterValue = counterValue
		log.Printf("#DEBUG: counterValue %s : %v has been added to the MemStorage\n", metricName, counterValue)
	}
}
