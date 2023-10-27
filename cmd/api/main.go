package main

import (
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/statuzproj/genie/utils/healthz"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/healthz", healthz.HealthCheck)
	http.Handle("/metrics", promhttp.Handler())
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("HTTP server error: %v", err)
	}
}
