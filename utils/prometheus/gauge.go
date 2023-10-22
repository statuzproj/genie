package prometheus

import (
	"github.com/prometheus/client_golang/prometheus"
	"log"
)

func SetWebpagePrometheusMetric() {
	log.Println("test")
	metric := "statuz_genie"
	gauge := prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: metric,
			Help: "Indicates whether the webpage is up (1) or down (0)",
		},
		[]string{"service"}, // Add a "url" label
	)

	gauge.WithLabelValues("genie").Set(float64(1))

	prometheus.Register(gauge)
}
