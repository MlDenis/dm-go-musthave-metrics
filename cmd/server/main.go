package main

import (
	"flag"
	"github.com/MlDenis/dm-go-musthave-metrics/internal/configurator"
	"github.com/MlDenis/dm-go-musthave-metrics/internal/server"
)

var serverAdress *string

func init() {
	serverAdress = configurator.GetEnv("ADDRESS", flag.String("a", "localhost:8080", "SendingAdress"))
}

func main() {
	flag.Parse()

	s := server.MakeNewMSServer(*serverAdress)
	s.ServerStart()

}
