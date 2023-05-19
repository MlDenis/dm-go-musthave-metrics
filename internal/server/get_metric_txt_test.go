package server

import (
	"fmt"
	"github.com/MlDenis/dm-go-musthave-metrics/internal/metric"
	"github.com/MlDenis/dm-go-musthave-metrics/internal/storage"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetSingleValue(t *testing.T) {
	gin.SetMode(gin.TestMode)

	testCases := []struct {
		vt         string
		name       string
		statusCode int
		response   string
	}{
		{vt: metric.GaugeString, name: "Mallocs", statusCode: http.StatusOK, response: "10"},
		{vt: metric.CounterString, name: "PollCount", statusCode: http.StatusOK, response: "5"},
		{vt: "invalid", name: "invalid", statusCode: http.StatusNotFound, response: "text/plain"},
	}

	ms := storage.NewMetricsStorage()
	ms.UpdateMetricInStorage(metric.GaugeString, "Mallocs", 10.0, 0)
	ms.UpdateMetricInStorage(metric.CounterString, "PollCount", 0.0, 5)

	s := &MSServer{MS: ms}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%s_%s", tc.vt, tc.name), func(t *testing.T) {
			router := gin.Default()
			router.GET("/:type/:name", s.GetSingleValue)

			req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/%s/%s", tc.vt, tc.name), nil)
			resp := httptest.NewRecorder()

			router.ServeHTTP(resp, req)

			assert.Equal(t, tc.statusCode, resp.Code)
			if tc.statusCode == http.StatusOK {
				assert.Equal(t, tc.response, resp.Body.String())
			}
		})
	}
}
