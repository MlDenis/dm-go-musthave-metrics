package sendler

import (
	"fmt"
	"github.com/MlDenis/dm-go-musthave-metrics/environment"
	"github.com/MlDenis/dm-go-musthave-metrics/internal/collector"
	"io"
	"log"
	"net/http"
	"time"
)

type Agent struct {
	config            environment.AgentConfig
	metricsDataBuffer collector.MetricsDataBuffer
	client            *http.Client
}

func MakeNewAgent(acfg environment.AgentConfig) *Agent {
	mdb := collector.MakeNewDataBuffer()
	clt := &http.Client{}

	return &Agent{
		config:            acfg,
		metricsDataBuffer: mdb,
		client:            clt,
	}
}

func (a *Agent) UpdateMetricsData() {
	a.metricsDataBuffer.CollectMetricData()
}

func (a *Agent) SendMetricData(metricName string) error {
	c := &http.Client{}
	var requestData string

	switch (a.metricsDataBuffer.Data[metricName]).MetricType {
	case "gauge":
		requestData = fmt.Sprintf("http://%s/update/%s/%s/%f",
			a.config.SendingAdress,
			a.metricsDataBuffer.Data[metricName].MetricType,
			metricName,
			a.metricsDataBuffer.Data[metricName].GaugeValue,
		)
	case "counter":
		requestData = fmt.Sprintf("http://%s/update/%s/%s/%d",
			a.config.SendingAdress,
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
		return fmt.Errorf("failed to read the request body: %w", err)
	}

	return nil
}

func (a *Agent) SendMetricsData() {
	for metric := range a.metricsDataBuffer.Data {
		err := a.SendMetricData(metric)
		if err != nil {
			log.Println("Got an error while executing the command a.SendMetricData(metric): ", err)
		}
	}

	log.Printf("Input counter value is: %v", a.metricsDataBuffer.Data["PollCount"])
}

func (a *Agent) DoTheJob() {
	for {
		time.Sleep(a.config.PollIntervalS)
		a.UpdateMetricsData()
		time.Sleep(a.config.PollIntervalS)
		a.UpdateMetricsData()
		time.Sleep(a.config.PollIntervalS)
		a.UpdateMetricsData()
		time.Sleep(a.config.PollIntervalS)
		a.UpdateMetricsData()
		time.Sleep(a.config.PollIntervalS)
		a.UpdateMetricsData()
		a.SendMetricsData()
	}
}
