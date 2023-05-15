package client

import (
	"github.com/MlDenis/dm-go-musthave-metrics/internal/metric"
	"github.com/MlDenis/dm-go-musthave-metrics/internal/storage"
	"math/rand"
	"runtime"
	"time"
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

	// TODO: Обсудить 1 на 1 как избежать ручного ввода каждого параметра, и иметь всё параметры в одной структуре уже при старте, перерабатывая инфу из массива в цикле

	mdf.Data["Alloc"] = &storage.MetricData{"gauge", metric.Gauge(memStats.Alloc), -1}
	mdf.Data["BuckHashSys"] = &storage.MetricData{"gauge", metric.Gauge(memStats.BuckHashSys), -1}
	mdf.Data["Frees"] = &storage.MetricData{"gauge", metric.Gauge(memStats.Frees), -1}

	mdf.Data["GCCPUFraction"] = &storage.MetricData{"gauge", metric.Gauge(memStats.GCCPUFraction), -1}
	mdf.Data["GCSys"] = &storage.MetricData{"gauge", metric.Gauge(memStats.GCSys), -1}
	mdf.Data["HeapAlloc"] = &storage.MetricData{"gauge", metric.Gauge(memStats.HeapAlloc), -1}

	mdf.Data["HeapIdle"] = &storage.MetricData{"gauge", metric.Gauge(memStats.HeapIdle), -1}
	mdf.Data["HeapInuse"] = &storage.MetricData{"gauge", metric.Gauge(memStats.HeapInuse), -1}
	mdf.Data["HeapObjects"] = &storage.MetricData{"gauge", metric.Gauge(memStats.HeapObjects), -1}

	mdf.Data["HeapReleased"] = &storage.MetricData{"gauge", metric.Gauge(memStats.HeapReleased), -1}
	mdf.Data["HeapSys"] = &storage.MetricData{"gauge", metric.Gauge(memStats.HeapSys), -1}
	mdf.Data["LastGC"] = &storage.MetricData{"gauge", metric.Gauge(memStats.LastGC), -1}

	mdf.Data["Lookups"] = &storage.MetricData{"gauge", metric.Gauge(memStats.Lookups), -1}
	mdf.Data["MCacheInuse"] = &storage.MetricData{"gauge", metric.Gauge(memStats.MCacheInuse), -1}
	mdf.Data["MCacheSys"] = &storage.MetricData{"gauge", metric.Gauge(memStats.MCacheSys), -1}

	mdf.Data["MSpanInuse"] = &storage.MetricData{"gauge", metric.Gauge(memStats.MSpanInuse), -1}
	mdf.Data["MSpanSys"] = &storage.MetricData{"gauge", metric.Gauge(memStats.MSpanSys), -1}
	mdf.Data["Mallocs"] = &storage.MetricData{"gauge", metric.Gauge(memStats.Mallocs), -1}

	mdf.Data["NextGC"] = &storage.MetricData{"gauge", metric.Gauge(memStats.NextGC), -1}
	mdf.Data["NumForcedGC"] = &storage.MetricData{"gauge", metric.Gauge(memStats.NumForcedGC), -1}
	mdf.Data["NumGC"] = &storage.MetricData{"gauge", metric.Gauge(memStats.NumGC), -1}

	mdf.Data["OtherSys"] = &storage.MetricData{"gauge", metric.Gauge(memStats.OtherSys), -1}
	mdf.Data["PauseTotalNs"] = &storage.MetricData{"gauge", metric.Gauge(memStats.PauseTotalNs), -1}
	mdf.Data["StackInuse"] = &storage.MetricData{"gauge", metric.Gauge(memStats.StackInuse), -1}

	mdf.Data["StackSys"] = &storage.MetricData{"gauge", metric.Gauge(memStats.StackSys), -1}
	mdf.Data["Sys"] = &storage.MetricData{"gauge", metric.Gauge(memStats.Sys), -1}
	mdf.Data["TotalAlloc"] = &storage.MetricData{"gauge", metric.Gauge(memStats.TotalAlloc), -1}

	rv := metric.Gauge(rand.Float64())
	mdf.Data["RandomValue"] = &storage.MetricData{"gauge", rv, -1}

	_, ok := mdf.Data["PollCount"]
	if ok {
		(mdf.Data["PollCount"]).CounterValue += 1
	} else {
		mdf.Data["PollCount"] = &storage.MetricData{"counter", -1.0, 1}
	}

	time.Sleep(*time.Second)
}
