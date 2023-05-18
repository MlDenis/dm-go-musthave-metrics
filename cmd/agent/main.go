package main

import (
	"github.com/MlDenis/dm-go-musthave-metrics/environment"
	"github.com/MlDenis/dm-go-musthave-metrics/internal/sendler"
)

func main() {

	acfg := environment.NewAgentConfig()

	a := sendler.MakeNewAgent(acfg)

	a.DoTheJob()

}
