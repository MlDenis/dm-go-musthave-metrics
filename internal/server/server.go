package server

import (
	"fmt"
	"github.com/MlDenis/dm-go-musthave-metrics/internal/metric"
	"github.com/MlDenis/dm-go-musthave-metrics/internal/storage"
	"log"
	"net/http"
	"strconv"
	"strings"
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
	mux.HandleFunc("/update/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")

		segmentsData := strings.Split(r.URL.Path, "/")
		if len(segmentsData) != 5 {
			http.Error(w, "Incorrect request", http.StatusBadRequest)
			return
		}
		vt := segmentsData[2]
		name := segmentsData[3]
		value := segmentsData[4]

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
		return
	})

	err := http.ListenAndServe(s.addr, mux)
	if err != nil {
		log.Printf(fmt.Sprint(err))
	}
	log.Printf("#DEBUG Server listen and serve on %s", s.addr)
}
