package healthz

import (
	"github.com/statuzproj/genie/utils/prometheus"
	"net/http"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {

	prometheus.SetWebpagePrometheusMetric()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status": "up"}`))
}
