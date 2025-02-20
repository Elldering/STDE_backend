package controllers

import (
	"STDE_proj/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

var (
	httpRequestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests",
		},
		[]string{"method", "endpoint"},
	)
	httpDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_duration_seconds",
			Help:    "Duration of HTTP requests",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"endpoint"},
	)
)

func init() {
	prometheus.MustRegister(httpRequestsTotal, httpDuration)
}

func HelloHandler(c *gin.Context) {
	timer := prometheus.NewTimer(httpDuration.WithLabelValues("/hello"))
	defer timer.ObserveDuration()

	httpRequestsTotal.WithLabelValues(http.MethodGet, "/hello").Inc()
	c.String(http.StatusOK, "Hello, World!")
}

func MetricsHandler() gin.HandlerFunc {
	return gin.WrapH(promhttp.Handler())
}

func UserCountHandler(svc *services.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		count, err := svc.GetUserCount()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user count"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"user_count": count})
	}
}
