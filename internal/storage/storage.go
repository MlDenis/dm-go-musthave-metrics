package storage

import (
	"fmt"
	"github.com/MlDenis/dm-go-musthave-metrics/internal/metric"
	"github.com/gin-gonic/gin"
	"strconv"
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

}

// GetStorageInfo - A method for returning data about a specific metric in text form.
func (ms *MemStorage) GetStorageInfo(vt string, name string) (string, error) {

	switch vt {
	case metric.GaugeString:
		value, ok := ms.data[name]
		if ok {
			return strconv.FormatFloat(float64(value.GaugeValue), 'E', -1, 64), nil
		}
	case metric.CounterString:
		value, ok := ms.data[name]
		if ok {
			return strconv.FormatInt(int64(value.CounterValue), 10), nil
		}
	}
	return fmt.Sprintf("Value with key %s not found", name), gin.Error{}
}

// GetHTMLPageInfo - Method for returning data about all available metric values in the form of an html document.
func (ms *MemStorage) GetHTMLPageInfo() (string, error) {
	var pageInfo strings.Builder

	for name, value := range ms.data {
		switch value.MetricType {
		case metric.GaugeString:
			_, err := fmt.Fprintf(
				&pageInfo,
				"* %s :  %v \n",
				name,
				strconv.FormatFloat(float64(value.GaugeValue),
					'E',
					-1,
					64))
			if err != nil {
				return "", fmt.Errorf("failure return number of bytes with fmt.Fprintf. "+
					"Posible error on the storage side %w", err)
			}
		case metric.CounterString:
			_, err := fmt.Fprintf(
				&pageInfo,
				"* %s :  %v \n",
				name,
				strconv.FormatInt(int64(value.CounterValue),
					10))
			if err != nil {
				return "", fmt.Errorf("failure return number of bytes with fmt.Fprintf. "+
					"Posible error on the storage side %w", err)
			}
		}
	}
	return pageInfo.String(), nil

}
