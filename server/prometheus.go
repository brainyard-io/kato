package server

import (
	"net/http"

	//"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// initPrometheusEndpoint initiates a prometheus server
func (s *Server) initPrometheusEndpoint() *Server {
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":9000", nil)
	return s
}
