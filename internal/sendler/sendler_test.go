package sendler

import (
	"github.com/MlDenis/dm-go-musthave-metrics/internal/collector"
	"net/http"
	"testing"
	"time"
)

func TestAgent_UpdateMetricsData(t *testing.T) {
	type fields struct {
		config            AgentConfig
		metricsDataBuffer collector.MetricsDataBuffer
		client            *http.Client
		pollTicker        *time.Ticker
		reportTicker      *time.Ticker
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Agent{
				config:            tt.fields.config,
				metricsDataBuffer: tt.fields.metricsDataBuffer,
				client:            tt.fields.client,
				pollTicker:        tt.fields.pollTicker,
				reportTicker:      tt.fields.reportTicker,
			}
			a.UpdateMetricsData()
		})
	}
}
