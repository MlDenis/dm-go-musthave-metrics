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

	mdf.Data["Alloc"] = &storage.MetricData{MetricType: "gauge", GaugeValue: metric.Gauge(memStats.Alloc)}
	mdf.Data["BuckHashSys"] = &storage.MetricData{MetricType: "gauge", GaugeValue: metric.Gauge(memStats.BuckHashSys)}
	mdf.Data["Frees"] = &storage.MetricData{MetricType: "gauge", GaugeValue: metric.Gauge(memStats.Frees)}

	mdf.Data["GCCPUFraction"] = &storage.MetricData{MetricType: "gauge", GaugeValue: metric.Gauge(memStats.GCCPUFraction)}
	mdf.Data["GCSys"] = &storage.MetricData{MetricType: "gauge", GaugeValue: metric.Gauge(memStats.GCSys)}
	mdf.Data["HeapAlloc"] = &storage.MetricData{MetricType: "gauge", GaugeValue: metric.Gauge(memStats.HeapAlloc)}

	mdf.Data["HeapIdle"] = &storage.MetricData{MetricType: "gauge", GaugeValue: metric.Gauge(memStats.HeapIdle)}
	mdf.Data["HeapInuse"] = &storage.MetricData{MetricType: "gauge", GaugeValue: metric.Gauge(memStats.HeapInuse)}
	mdf.Data["HeapObjects"] = &storage.MetricData{MetricType: "gauge", GaugeValue: metric.Gauge(memStats.HeapObjects)}

	mdf.Data["HeapReleased"] = &storage.MetricData{MetricType: "gauge", GaugeValue: metric.Gauge(memStats.HeapReleased)}
	mdf.Data["HeapSys"] = &storage.MetricData{MetricType: "gauge", GaugeValue: metric.Gauge(memStats.HeapSys)}
	mdf.Data["LastGC"] = &storage.MetricData{MetricType: "gauge", GaugeValue: metric.Gauge(memStats.LastGC)}

	mdf.Data["Lookups"] = &storage.MetricData{MetricType: "gauge", GaugeValue: metric.Gauge(memStats.Lookups)}
	mdf.Data["MCacheInuse"] = &storage.MetricData{MetricType: "gauge", GaugeValue: metric.Gauge(memStats.MCacheInuse)}
	mdf.Data["MCacheSys"] = &storage.MetricData{MetricType: "gauge", GaugeValue: metric.Gauge(memStats.MCacheSys)}

	mdf.Data["MSpanInuse"] = &storage.MetricData{MetricType: "gauge", GaugeValue: metric.Gauge(memStats.MSpanInuse)}
	mdf.Data["MSpanSys"] = &storage.MetricData{MetricType: "gauge", GaugeValue: metric.Gauge(memStats.MSpanSys)}
	mdf.Data["Mallocs"] = &storage.MetricData{MetricType: "gauge", GaugeValue: metric.Gauge(memStats.Mallocs)}

	mdf.Data["NextGC"] = &storage.MetricData{MetricType: "gauge", GaugeValue: metric.Gauge(memStats.NextGC)}
	mdf.Data["NumForcedGC"] = &storage.MetricData{MetricType: "gauge", GaugeValue: metric.Gauge(memStats.NumForcedGC)}
	mdf.Data["NumGC"] = &storage.MetricData{MetricType: "gauge", GaugeValue: metric.Gauge(memStats.NumGC)}

	mdf.Data["OtherSys"] = &storage.MetricData{MetricType: "gauge", GaugeValue: metric.Gauge(memStats.OtherSys)}
	mdf.Data["PauseTotalNs"] = &storage.MetricData{MetricType: "gauge", GaugeValue: metric.Gauge(memStats.PauseTotalNs)}
	mdf.Data["StackInuse"] = &storage.MetricData{MetricType: "gauge", GaugeValue: metric.Gauge(memStats.StackInuse)}

	mdf.Data["StackSys"] = &storage.MetricData{MetricType: "gauge", GaugeValue: metric.Gauge(memStats.StackSys)}
	mdf.Data["Sys"] = &storage.MetricData{MetricType: "gauge", GaugeValue: metric.Gauge(memStats.Sys)}
	mdf.Data["TotalAlloc"] = &storage.MetricData{MetricType: "gauge", GaugeValue: metric.Gauge(memStats.TotalAlloc)}

	rv := metric.Gauge(rand.Float64())
	mdf.Data["RandomValue"] = &storage.MetricData{MetricType: "gauge", GaugeValue: rv}

	_, ok := mdf.Data["PollCount"]
	if ok {
		(mdf.Data["PollCount"]).CounterValue += 1
	} else {
		mdf.Data["PollCount"] = &storage.MetricData{MetricType: "counter", CounterValue: 1}
	}
}
