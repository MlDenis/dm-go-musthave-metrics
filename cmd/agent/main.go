package main

import (
	"github.com/MlDenis/dm-go-musthave-metrics/internal/sendler"
	"sync"
	"time"
)

var wg sync.WaitGroup

const sendingAdress = "localhost"
const sendingPort = "8080"
const pollInterval = 2
const reportInterval = 10

func main() {
	ac := sendler.InitAgentConfig(
		sendingAdress,
		sendingPort,
		pollInterval,
		reportInterval,
	)

	a := sendler.MakeNewAgent(ac)

	for {
		time.Sleep(time.Second * pollInterval)
		a.UpdateMetricsData()
		time.Sleep(time.Second * pollInterval)
		a.UpdateMetricsData()
		time.Sleep(time.Second * pollInterval)
		a.UpdateMetricsData()
		time.Sleep(time.Second * pollInterval)
		a.UpdateMetricsData()
		time.Sleep(time.Second * pollInterval)
		a.UpdateMetricsData()
		a.SendMetricsData()
	}
}
