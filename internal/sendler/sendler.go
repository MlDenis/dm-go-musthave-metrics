package client

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type Agent struct {
	config            AgentConfig
	metricsDataBuffer MetricsDataBuffer
	client            *http.Client
	pollTicker        *time.Ticker
	reportTicker      *time.Ticker
}

func MakeNewAgent(cfg AgentConfig) *Agent {
	mdb := MakeNewDataBuffer()
	clt := &http.Client{}

	return &Agent{
		config:            cfg,
		metricsDataBuffer: mdb,
		client:            clt,
		pollTicker:        time.NewTicker(cfg.pollInterval),
		reportTicker:      time.NewTicker(cfg.reportInterval),
	}
}

func (a *Agent) UpdateMetricsData() {
	a.metricsDataBuffer.CollectMetricData()
	//log.Printf("#DEBUG UpdateMetricsData sucessfully complete")
}

func (a *Agent) SendMetricData(metricName string) error {
	log.Printf("#DEBUG we in SendMetricData")
	c := &http.Client{}
	var requestData string

	switch (a.metricsDataBuffer.data[metricName]).MetricType {
	case "gauge":
		requestData = fmt.Sprintf("%s:%s/update/%s/%s/%f",
			a.config.sendingAdress,
			a.config.sendingPort,
			a.metricsDataBuffer.data[metricName].MetricType,
			metricName,
			a.metricsDataBuffer.data[metricName].GaugeValue,
		)
	case "counter":
		requestData = fmt.Sprintf("%s:%s/update/%s/%s/%d",
			a.config.sendingAdress,
			a.config.sendingPort,
			a.metricsDataBuffer.data[metricName].MetricType,
			metricName,
			a.metricsDataBuffer.data[metricName].CounterValue,
		)
	}

	request, err := http.NewRequest(http.MethodPost, requestData, nil)
	if err != nil {
		return err
	}

	request.Header.Add("Content-Type", "text/plain")
	response, err := c.Do(request)
	if err != nil {
		return err
	}

	defer response.Body.Close()
	_, err = io.ReadAll(response.Body)
	if err != nil {
		return err
	}

	log.Printf("#DEBUG SendMetricData sucessfully complete")

	return nil
}

func (a *Agent) SendMetricsData() {
	for metric, _ := range a.metricsDataBuffer.data {
		err := a.SendMetricData(metric)
		if err != nil {
			return
		}
	}
}

func (a *Agent) DoTheJob() {

	for {
		select {
		case <-a.pollTicker.C:
			a.UpdateMetricsData()
		case <-time.After(a.config.reportInterval * time.Second):
			a.SendMetricsData()
			//case <-a.reportTicker.C:
			//	a.SendMetricsData()
		}

	}

}
