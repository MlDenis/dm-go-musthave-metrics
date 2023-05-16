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
		if floatVal, err := strconv.ParseFloat(value, 64); err == nil {
			s.MS.UpdateMetricInStorage(vt, name, metric.Gauge(floatVal), -1)
			ctx.String(http.StatusOK, "text/plain")
			return
		}
	case metric.CounterString:
		if intVal, err := strconv.ParseInt(value, 10, 64); err == nil {
			s.MS.UpdateMetricInStorage(vt, name, -1.0, metric.Counter(intVal))
			ctx.String(http.StatusOK, "text/plain")
			return
		}
	default:
		ctx.String(http.StatusNotImplemented, "text/plain")
		return
	}

}
