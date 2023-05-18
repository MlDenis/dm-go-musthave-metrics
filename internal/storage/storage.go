package storage

import (
	"fmt"
	"github.com/MlDenis/dm-go-musthave-metrics/internal/metric"
	"github.com/gin-gonic/gin"
	"strings"
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
			if metricName == "PollCount" {
				ms.data[metricName].CounterValue += counterValue
			} else {
				(ms.data[metricName]).CounterValue = counterValue
			}
		}
	} else {
		newElement := MetricData{
			metricType,
			gaugeValue,
			counterValue,
		}

		ms.data[metricName] = &newElement
	}

}

// GetStorageInfo - A method for returning data about a specific metric in text form.
func (ms *MemStorage) GetStorageInfo(vt string, name string) (string, error) {

	switch vt {
	case metric.GaugeString:
		value, ok := ms.data[name]
		if ok {
			return fmt.Sprintf("%+v", value.GaugeValue), nil
		}
	case metric.CounterString:
		value, ok := ms.data[name]
		if ok {
			return fmt.Sprintf("%+v", value.CounterValue), nil
		}
	}
	return fmt.Sprintf("Value with key %s not found", name), gin.Error{}
}

// GetHTMLPageInfo - Method for returning data about all available metric values in the form of an html document.
func (ms *MemStorage) GetHTMLPageInfo() string {
	var pageInfo strings.Builder

	for name, value := range ms.data {
		switch value.MetricType {
		case metric.GaugeString:
			_, err := fmt.Fprintf(&pageInfo, "* %s :  %v \n", name, fmt.Sprint(value.GaugeValue))
			if err != nil {
				return fmt.Sprint(err)
			}
		case metric.CounterString:
			_, err := fmt.Fprintf(&pageInfo, "* %s :  %v \n", name, fmt.Sprint(value.CounterValue))
			if err != nil {
				return fmt.Sprint(err)
			}
		}
	}
	return pageInfo.String()

}
