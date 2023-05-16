package main

import (
	"flag"
	"github.com/MlDenis/dm-go-musthave-metrics/internal/configurator"
	"github.com/MlDenis/dm-go-musthave-metrics/internal/sendler"
	"log"
	"strconv"
	"strings"
	"sync"
	"time"
)

var wg sync.WaitGroup

var (
	sendingAdress *string
	pollIntervalS,
	reportIntervalS string
)

func init() {
	sendingAdress = configurator.GetEnv("ADDRESS", flag.String("a", "localhost:8080", "SENDING_ADRESS"))
	pollIntervalS = strings.Replace(
		*configurator.GetEnv("POLL_INTERVAL",
			flag.String("p", "2", "POLL_INTERVAL")),
		"s",
		"",
		-1)
	reportIntervalS = strings.Replace(
		*configurator.GetEnv("REPORT_INTERVAL",
			flag.String("r", "10", "REPORT_INTERVAL")),
		"s",
		"",
		-1)
}

func main() {
	flag.Parse()

	pollInterval, err := strconv.Atoi(pollIntervalS)
	if err != nil {
		log.Fatalf("Error happened in reading poll counter variable. Err: %s", err)
	}

	reportInterval, err := strconv.Atoi(reportIntervalS)
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
		time.Sleep(time.Second * time.Duration(pollInterval))
		a.UpdateMetricsData()
		time.Sleep(time.Second * time.Duration(pollInterval))
		a.UpdateMetricsData()
		time.Sleep(time.Second * time.Duration(pollInterval))
		a.UpdateMetricsData()
		time.Sleep(time.Second * time.Duration(pollInterval))
		a.UpdateMetricsData()
		time.Sleep(time.Second * time.Duration(pollInterval))
		a.UpdateMetricsData()
		a.SendMetricsData()
	}
}
