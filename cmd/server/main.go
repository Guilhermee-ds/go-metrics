package main

import (
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	paymentsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "ecomerce_payments_total",
			Help: "Total number of payments processed",
		},
		[]string{"status"},
	)
	httpDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "ecommerce_http_duration",
			Help:    "HTTP request duration is seconds",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"handler"},
	)
)

func middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		duration := prometheus.NewTimer(httpDuration.WithLabelValues(r.URL.Path))
		next.ServeHTTP(w, r)
		duration.ObserveDuration()
	})
}

func init() {
	prometheus.MustRegister(paymentsTotal, httpDuration)
}

func main() {
	mux := http.NewServeMux()

	mux.Handle("POST /payments", middleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		status := "success"

		if err := r.ParseForm(); err != nil {
			http.Error(w, "unprocessable emtity", http.StatusUnprocessableEntity)
			return
		}

		if r.PostForm.Get("status") != "success" {
			status = "failed"
		}

		paymentsTotal.WithLabelValues(status).Inc()
		w.Write([]byte("Payment requested"))
	})))

	mux.Handle("GET /metrics", promhttp.Handler())
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
