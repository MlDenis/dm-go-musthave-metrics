package main

import (
	"github.com/MlDenis/dm-go-musthave-metrics/environment"
	"github.com/MlDenis/dm-go-musthave-metrics/internal/server"
)

var serverAdress *string

func main() {
	scfg := environment.NewServerConfig()
	s := server.MakeNewMSServer(scfg.HostAdress)
	s.ServerStart()

}
