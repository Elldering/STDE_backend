package services

import (
	"STDE_proj/internal/repositories"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	UserCount = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "user_count_total",
		Help: "Total number of users",
	})
	OrderCount = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "order_count_total",
		Help: "Total number of orders",
	})
	BlockedUsers = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "blocked_users_total",
		Help: "Total number of blocked users",
	})
	ReviewCount = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "review_count_total",
		Help: "Total number of reviews",
	})
	AverageReviewRating = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "average_review_rating",
		Help: "Average rating of reviews",
	})
	ActiveProjects = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "active_projects_total",
		Help: "Total number of active projects",
	})
	// Используем гистограмму для измерения времени выполнения задач
	AverageTaskTime = prometheus.NewHistogram(prometheus.HistogramOpts{
		Name:    "average_task_completion_time",
		Help:    "Average time to complete a task",
		Buckets: prometheus.DefBuckets,
	})
	AverageOrderPrice = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "average_order_price",
		Help: "Average price of orders",
	})
	TotalRevenue = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "total_revenue",
		Help: "Total revenue from orders",
	})
	OrderPositionCount = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "order_position_count_total",
		Help: "Total number of order positions",
	})
	AverageOrderItems = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "average_order_items",
		Help: "Average number of items per order",
	})

	// HTTP-метрики
	HttpRequestCountWithPath = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "http_requests_total_with_path",
		Help: "Number of HTTP requests by path.",
	}, []string{"url"})
	HttpRequestDuration = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name: "http_request_duration_seconds",
		Help: "Response time of HTTP request.",
	}, []string{"path"})
)

func init() {
	prometheus.MustRegister(
		UserCount,
		OrderCount,
		BlockedUsers,
		ReviewCount,
		AverageReviewRating,
		ActiveProjects,
		AverageTaskTime,
		AverageOrderPrice,
		TotalRevenue,
		OrderPositionCount,
		AverageOrderItems,
		HttpRequestCountWithPath,
		HttpRequestDuration,
	)
}

// Обёртки для вызова функций репозитория

func GetUserCount() (int, error) {
	return repositories.GetUserCount()
}

func GetOrderCount() (int, error) {
	return repositories.GetOrderCount()
}

func GetBlockedUserCount() (int, error) {
	return repositories.GetBlockedUserCount()
}

func GetReviewCount() (int, error) {
	return repositories.GetReviewCount()
}

func GetAverageReviewRating() (float64, error) {
	return repositories.GetAverageReviewRating()
}

func GetAverageOrderPrice() (float64, error) {
	return repositories.GetAverageOrderPrice()
}

func GetTotalRevenue() (float64, error) {
	return repositories.GetTotalRevenue()
}

func GetOrderPositionCount() (int, error) {
	return repositories.GetOrderPositionCount()
}

func GetAverageOrderItems() (float64, error) {
	return repositories.GetAverageOrderItems()
}

// Регистрация эндпоинта для метрик
func RegisterMetricsEndpoint(r *gin.Engine) {
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))
}
