package server

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func (s *MSServer) GetSingleValue(ctx *gin.Context) {
	vt := ctx.Params.ByName(metrics.TypeS)
	name := ctx.Params.ByName(metrics.NameS)

	switch vt {
	case metrics.GaugeS:
		value, err := s.MetricsStorage.GetMetricInfo(vt, name)
		if err != nil {
			ctx.String(http.StatusNotFound, services.CtText)
			return
		}
		ctx.String(http.StatusOK, "%+v", value)
		log.Printf("#DEBUG run GetSingleValue with: value type = %s, name = %s, value = %s.\n", vt, name, value)
		return
	case metrics.CounterS:
		value, err := s.MetricsStorage.GetMetricInfo(vt, name)
		if err != nil {
			ctx.String(http.StatusNotFound, services.CtText)
			return
		}
		ctx.String(http.StatusOK, "%+v", value)
		log.Printf("#DEBUG run GetSingleValue with: value type = %s, name = %s, value = %s.\n", vt, name, value)
		return
	}
	ctx.String(http.StatusNotFound, services.CtText)
}
