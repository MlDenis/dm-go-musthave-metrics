package main

import (
	"github.com/MlDenis/dm-go-musthave-metrics/internal/server"
)

func main() {
	s := server.MakeNewMSServer("localhost:8080")
	s.DoTheJob()
}
