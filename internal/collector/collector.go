package collector

import (
	"github.com/MlDenis/dm-go-musthave-metrics/internal/metric"
	"github.com/MlDenis/dm-go-musthave-metrics/internal/storage"
	"math/rand"
	"runtime"
)

type MetricsDataBuffer struct {
	Data map[string]*storage.MetricData
}

func MakeNewDataBuffer() MetricsDataBuffer {

	mdf := MetricsDataBuffer{make(map[string]*storage.MetricData)}

	return mdf
}

func (mdf *MetricsDataBuffer) CollectMetricData() {
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)

	mdf.Data["Alloc"] = &storage.MetricData{MetricType: "gauge", GaugeValue: metric.Gauge(memStats.Alloc), CounterValue: -1}
	mdf.Data["BuckHashSys"] = &storage.MetricData{MetricType: "gauge", GaugeValue: metric.Gauge(memStats.BuckHashSys), CounterValue: -1}
	mdf.Data["Frees"] = &storage.MetricData{MetricType: "gauge", GaugeValue: metric.Gauge(memStats.Frees), CounterValue: -1}

	mdf.Data["GCCPUFraction"] = &storage.MetricData{MetricType: "gauge", GaugeValue: metric.Gauge(memStats.GCCPUFraction), CounterValue: -1}
	mdf.Data["GCSys"] = &storage.MetricData{MetricType: "gauge", GaugeValue: metric.Gauge(memStats.GCSys), CounterValue: -1}
	mdf.Data["HeapAlloc"] = &storage.MetricData{MetricType: "gauge", GaugeValue: metric.Gauge(memStats.HeapAlloc), CounterValue: -1}

	mdf.Data["HeapIdle"] = &storage.MetricData{MetricType: "gauge", GaugeValue: metric.Gauge(memStats.HeapIdle), CounterValue: -1}
	mdf.Data["HeapInuse"] = &storage.MetricData{MetricType: "gauge", GaugeValue: metric.Gauge(memStats.HeapInuse), CounterValue: -1}
	mdf.Data["HeapObjects"] = &storage.MetricData{MetricType: "gauge", GaugeValue: metric.Gauge(memStats.HeapObjects), CounterValue: -1}

	mdf.Data["HeapReleased"] = &storage.MetricData{MetricType: "gauge", GaugeValue: metric.Gauge(memStats.HeapReleased), CounterValue: -1}
	mdf.Data["HeapSys"] = &storage.MetricData{MetricType: "gauge", GaugeValue: metric.Gauge(memStats.HeapSys), CounterValue: -1}
	mdf.Data["LastGC"] = &storage.MetricData{MetricType: "gauge", GaugeValue: metric.Gauge(memStats.LastGC), CounterValue: -1}

	mdf.Data["Lookups"] = &storage.MetricData{MetricType: "gauge", GaugeValue: metric.Gauge(memStats.Lookups), CounterValue: -1}
	mdf.Data["MCacheInuse"] = &storage.MetricData{MetricType: "gauge", GaugeValue: metric.Gauge(memStats.MCacheInuse), CounterValue: -1}
	mdf.Data["MCacheSys"] = &storage.MetricData{MetricType: "gauge", GaugeValue: metric.Gauge(memStats.MCacheSys), CounterValue: -1}

	mdf.Data["MSpanInuse"] = &storage.MetricData{MetricType: "gauge", GaugeValue: metric.Gauge(memStats.MSpanInuse), CounterValue: -1}
	mdf.Data["MSpanSys"] = &storage.MetricData{MetricType: "gauge", GaugeValue: metric.Gauge(memStats.MSpanSys), CounterValue: -1}
	mdf.Data["Mallocs"] = &storage.MetricData{MetricType: "gauge", GaugeValue: metric.Gauge(memStats.Mallocs), CounterValue: -1}

	mdf.Data["NextGC"] = &storage.MetricData{MetricType: "gauge", GaugeValue: metric.Gauge(memStats.NextGC), CounterValue: -1}
	mdf.Data["NumForcedGC"] = &storage.MetricData{MetricType: "gauge", GaugeValue: metric.Gauge(memStats.NumForcedGC), CounterValue: -1}
	mdf.Data["NumGC"] = &storage.MetricData{MetricType: "gauge", GaugeValue: metric.Gauge(memStats.NumGC), CounterValue: -1}

	mdf.Data["OtherSys"] = &storage.MetricData{MetricType: "gauge", GaugeValue: metric.Gauge(memStats.OtherSys), CounterValue: -1}
	mdf.Data["PauseTotalNs"] = &storage.MetricData{MetricType: "gauge", GaugeValue: metric.Gauge(memStats.PauseTotalNs), CounterValue: -1}
	mdf.Data["StackInuse"] = &storage.MetricData{MetricType: "gauge", GaugeValue: metric.Gauge(memStats.StackInuse), CounterValue: -1}

	mdf.Data["StackSys"] = &storage.MetricData{MetricType: "gauge", GaugeValue: metric.Gauge(memStats.StackSys), CounterValue: -1}
	mdf.Data["Sys"] = &storage.MetricData{MetricType: "gauge", GaugeValue: metric.Gauge(memStats.Sys), CounterValue: -1}
	mdf.Data["TotalAlloc"] = &storage.MetricData{MetricType: "gauge", GaugeValue: metric.Gauge(memStats.TotalAlloc), CounterValue: -1}

	rv := metric.Gauge(rand.Float64())
	mdf.Data["RandomValue"] = &storage.MetricData{MetricType: "gauge", GaugeValue: rv, CounterValue: -1}

	_, ok := mdf.Data["PollCount"]
	if ok {
		(mdf.Data["PollCount"]).CounterValue += 1
	} else {
		mdf.Data["PollCount"] = &storage.MetricData{MetricType: "counter", GaugeValue: -1.0, CounterValue: 1}
	}
}
