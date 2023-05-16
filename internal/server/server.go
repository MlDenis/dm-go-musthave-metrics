package server

import (
	"fmt"
	"github.com/MlDenis/dm-go-musthave-metrics/internal/metric"
	"github.com/MlDenis/dm-go-musthave-metrics/internal/storage"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type ServerConfig struct {
	addr string
	prt  string
}

func NewServerConfig(address string, port string) ServerConfig {
	return ServerConfig{addr: address, prt: port}
}

type MSServer struct {
	MS     storage.MemStorage
	Config ServerConfig
	Router *gin.Engine
}

func MakeNewMSServer(adress string, port string) MSServer {
	sc := NewServerConfig(adress, port)
	ms := storage.NewMetricsStorage()
	s := MSServer{MS: ms, Config: sc}

	r := gin.Default()
	r.RedirectTrailingSlash = false

	r.POST("/update/:type/:name/:value", s.PostSingleValue)
	r.GET("/value/:type/:name", s.GetSingleValue)
	r.GET("/", s.GetMSDataHowHTML)

	s.Router = r

	return s
}

func (s *MSServer) DoTheJob() {
	mux := http.NewServeMux()
	mux.HandleFunc("/update/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("#DEBUG we in mux.HandleFunc(\"/update/\", func(w http.ResponseWriter, r *http.Request) ")

		w.Header().Set("Content-Type", "text/plain")

		segmentsData := strings.Split(r.URL.Path, "/")
		vt := segmentsData[2]
		if vt != "gauge" && vt != "counter" {
			http.Error(w, "Incorrect request", http.StatusBadRequest)
			return
		}
		name := segmentsData[3]
		if name == "" {
			http.Error(w, "Incorrect request", http.StatusNotFound)
			return
		}
		value := segmentsData[4]
		if value == "" {
			http.Error(w, "Incorrect request", http.StatusBadRequest)
			return
		}

		switch vt {
		case metric.GaugeString:
			convertedValueFloat, err := strconv.ParseFloat(value, 64)
			if err != nil {
				http.Error(w, "Incorrect data", http.StatusBadRequest)
				return
			}
			convertedValueGauge := metric.Gauge(convertedValueFloat)

			s.MS.UpdateMetricInStorage(
				vt,
				name,
				convertedValueGauge,
				-1)
			w.WriteHeader(http.StatusOK)
			return
		case metric.CounterString:
			convertedValueInt, err := strconv.ParseInt(value, 0, 64)
			if err != nil {
				http.Error(w, "Incorrect data", http.StatusBadRequest)
				return
			}
			convertedValueCounter := metric.Counter(convertedValueInt)

			s.MS.UpdateMetricInStorage(
				vt,
				name,
				0.0,
				convertedValueCounter)
			w.WriteHeader(http.StatusOK)
			return
		}
		http.Error(w, "Incorrect request", http.StatusBadRequest)
	})

	workAdress := fmt.Sprintf("%s:%s", s.Config.addr, s.Config.prt)

	err := http.ListenAndServe(workAdress, mux)
	if err != nil {
		log.Println(fmt.Sprint(err))
	}
	log.Printf("Server listen and serve on %s", workAdress)
}
