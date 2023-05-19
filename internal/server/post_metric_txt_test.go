package server

import (
	"github.com/MlDenis/dm-go-musthave-metrics/internal/metric"
	"github.com/MlDenis/dm-go-musthave-metrics/internal/storage"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPostSingleValue(t *testing.T) {
	gin.SetMode(gin.TestMode)

	tests := []struct {
		vt, name, value string
		expectedStatus  int
	}{
		{metric.GaugeString, "Mallocs", "22669", http.StatusOK},
		{metric.CounterString, "PollCount", "15", http.StatusOK},
		{metric.GaugeString, "TestInvalid", "none", http.StatusBadRequest},
		{"baga", "TestUnknown", "111", http.StatusNotImplemented},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)

			c.Params = []gin.Param{
				{Key: TypeS, Value: tt.vt},
				{Key: NameS, Value: tt.name},
				{Key: ValueS, Value: tt.value},
			}

			s := &MSServer{MS: storage.NewMetricsStorage()}
			s.PostSingleValue(c)

			assert.Equal(t, tt.expectedStatus, w.Code)
		})
	}
}
