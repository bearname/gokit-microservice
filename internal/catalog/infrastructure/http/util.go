package http

import (
	"net/http"
	"strconv"
	"time"

	"github.com/go-kit/kit/metrics"
)

type MetricsHolder struct {
	RequestCount   metrics.Counter
	RequestLatency metrics.Histogram
}

type statusCodeWriter struct {
	http.ResponseWriter
	statusCode int
}

func (w *statusCodeWriter) WriteHeader(code int) {
	w.statusCode = code
	w.ResponseWriter.WriteHeader(code)
}

func NewMetricsHolder(counter metrics.Counter, latency metrics.Histogram) *MetricsHolder {
	return &MetricsHolder{
		RequestCount:   counter,
		RequestLatency: latency,
	}
}

func InstrumentingMiddleware(next http.Handler, metrics *MetricsHolder, endpointName string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		scw := &statusCodeWriter{w, http.StatusOK}

		defer func(start time.Time) {
			metrics.RequestCount.With("method", r.Method, "endpoint", endpointName, "status_code", strconv.Itoa(scw.statusCode)).Add(1)
			metrics.RequestLatency.With("method", r.Method, "endpoint", endpointName).Observe(time.Since(start).Seconds())
		}(time.Now())

		next.ServeHTTP(scw, r)
	})
}
