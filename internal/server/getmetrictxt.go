package server

import (
	"github.com/MlDenis/dm-go-musthave-metrics/internal/metric"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *MSServer) GetSingleValue(ctx *gin.Context) {
	vt := ctx.Params.ByName(TypeS)
	name := ctx.Params.ByName(NameS)

	switch vt {
	case metric.GaugeString:
		value, err := s.MS.GetStorageInfo(vt, name)
		if err != nil {
			ctx.String(http.StatusNotFound, "text/plain")
			return
		}
		ctx.String(http.StatusOK, "%+v", value)
		return
	case metric.CounterString:
		value, err := s.MS.GetStorageInfo(vt, name)
		if err != nil {
			ctx.String(http.StatusNotFound, "text/plain")
			return
		}
		ctx.String(http.StatusOK, "%+v", value)
		return
	}
	ctx.String(http.StatusNotFound, "text/plain")
}
