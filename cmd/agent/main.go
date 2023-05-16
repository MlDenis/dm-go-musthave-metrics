package main

import (
	"flag"
	"github.com/MlDenis/dm-go-musthave-metrics/internal/sendler"
	"log"
	"sync"
	"time"
)

var wg sync.WaitGroup

var (
	sendingAdress,
	pollIntervalS,
	reportIntervalS *string
)

func init() {
	sendingAdress = flag.String("a", "localhost:8080", "SendingAdress")
	pollIntervalS = flag.String("p", "2", "POLL_INTERVAL")
	reportIntervalS = flag.String("r", "10", "REPORT_INTERVAL")
}

func main() {
	flag.Parse()

	pollInterval, err := time.ParseDuration(*pollIntervalS)
	if err != nil {
		log.Fatalf("Error happened in reading poll counter variable. Err: %s", err)
	}

	reportInterval, err := time.ParseDuration(*reportIntervalS)
	if err != nil {
		log.Fatalf("Error happened in reading poll counter variable. Err: %s", err)
	}

	ac := sendler.InitAgentConfig(
		*sendingAdress,
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
