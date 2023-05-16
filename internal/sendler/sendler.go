package sendler

import (
	"fmt"
	"github.com/MlDenis/dm-go-musthave-metrics/internal/collector"
	"io"
	"log"
	"net/http"
	"time"
)

type AgentConfig struct {
	sendingAdress  string
	pollInterval   time.Duration
	reportInterval time.Duration
}

func InitAgentConfig(
	sendingAdress string,
	pollInterval time.Duration,
	reportInterval time.Duration,
) AgentConfig {
	return AgentConfig{
		sendingAdress,
		pollInterval,
		reportInterval,
	}
}

type Agent struct {
	config            AgentConfig
	metricsDataBuffer collector.MetricsDataBuffer
	client            *http.Client
	pollTicker        *time.Ticker
	reportTicker      *time.Ticker
}

func MakeNewAgent(cfg AgentConfig) *Agent {
	mdb := collector.MakeNewDataBuffer()
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
	log.Printf("#DEBUG UpdateMetricsData sucessfully complete")
}

func (a *Agent) SendMetricData(metricName string) error {
	log.Printf("#DEBUG we in SendMetricData")
	c := &http.Client{}
	var requestData string

	switch (a.metricsDataBuffer.Data[metricName]).MetricType {
	case "gauge":
		requestData = fmt.Sprintf("http://%s:%s/update/%s/%s/%f",
			a.config.sendingAdress,
			a.metricsDataBuffer.Data[metricName].MetricType,
			metricName,
			a.metricsDataBuffer.Data[metricName].GaugeValue,
		)
	case "counter":
		requestData = fmt.Sprintf("http://%s/update/%s/%s/%d",
			a.config.sendingAdress,
			a.metricsDataBuffer.Data[metricName].MetricType,
			metricName,
			a.metricsDataBuffer.Data[metricName].CounterValue,
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

	for {
		for metric := range a.metricsDataBuffer.Data {
			err := a.SendMetricData(metric)
			if err != nil {
				return
			}
		}
	}

}
