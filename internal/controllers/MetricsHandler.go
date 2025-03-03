package controllers

import (
	"STDE_proj/internal/services"
	"github.com/gin-gonic/gin"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	userCount = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "user_count_total",
			Help: "Total number of users",
		},
	)
	activeProjects = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "active_projects_total",
			Help: "Total number of active projects",
		},
	)
	averageTaskTime = prometheus.NewHistogram(
		prometheus.HistogramOpts{
			Name:    "average_task_completion_time",
			Help:    "Average time to complete a task",
			Buckets: prometheus.DefBuckets,
		},
	)
)

func init() {
	prometheus.MustRegister(userCount, activeProjects, averageTaskTime)
}

func UpdateMetrics(c *gin.Context) {
	count, err := services.GetUserCount()
	if err == nil {
		userCount.Set(float64(count))
	}

	projects, err := services.GetActiveProjectsCount()
	if err == nil {
		activeProjects.Set(float64(projects))
	}

	time, err := services.GetAverageTaskTime()
	if err == nil {
		averageTaskTime.Observe(time)
	}
}
