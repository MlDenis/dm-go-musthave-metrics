package server

import (
	"github.com/MlDenis/dm-go-musthave-metrics/internal/metric"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (s *MSServer) PostSingleValue(ctx *gin.Context) {

	vt := ctx.Params.ByName(TypeS)
	name := ctx.Params.ByName(NameS)
	value := ctx.Params.ByName(ValueS)

	if value == "none" {
		ctx.String(http.StatusBadRequest, "text/plain")
		return
	}

	switch vt {
	case metric.GaugeString:
		floatVal, err := strconv.ParseFloat(value, 64)
		if err != nil {
			ctx.String(http.StatusInternalServerError, "text/plain")
			return
		} else {
			s.MS.UpdateMetricInStorage(vt, name, metric.Gauge(floatVal), 0)
			ctx.String(http.StatusOK, "text/plain")
			return
		}
	case metric.CounterString:
		intVal, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			ctx.String(http.StatusInternalServerError, "text/plain")
			return
		} else {
			s.MS.UpdateMetricInStorage(vt, name, 0, metric.Counter(intVal))
			ctx.String(http.StatusOK, "text/plain")
			return
		}
	default:
		ctx.String(http.StatusNotImplemented, "text/plain")
		return
	}

}
