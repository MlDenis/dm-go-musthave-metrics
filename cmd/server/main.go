package main

import (
	"github.com/MlDenis/dm-go-musthave-metrics/internal/server"
)

func main() {

	var (
		serverAdress string
		serverPort   string
	)

	serverAdress = "localhost"
	serverPort = "8080"

	s := server.MakeNewMSServer(serverAdress, serverPort)
	s.ServerStart()

}
