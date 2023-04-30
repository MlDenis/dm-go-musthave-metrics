package server

import (
	"fmt"
	"github.com/MlDenis/dm-go-musthave-metrics/internal/handlers"
	"github.com/MlDenis/dm-go-musthave-metrics/internal/storage"
	"log"
	"net/http"
)

type MSServer struct {
	MS   storage.MemStorage
	addr string
}

func MakeNewMSServer(adress string) MSServer {
	return MSServer{MS: storage.MemStorage{}, addr: adress}
}

func (s *MSServer) DoTheJob() {
	mux := http.NewServeMux()
	mux.HandleFunc("/update/", handlers.HandlerPostMetricDataAsText)
	mux.HandleFunc("/update/", handlers.HandlerPostMetricDataAsText)

	err := http.ListenAndServe(s.addr, mux)
	if err != nil {
		log.Printf(fmt.Sprint(err))
	}
	log.Printf("#DEBUG Server listen and serve on %s", s.addr)
}
