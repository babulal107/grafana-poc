package router

import (
	"github.com/gin-gonic/gin"
	"github.com/grafana-poc/internal/model"
	"github.com/penglongli/gin-metrics/ginmetrics"
	"net/http"
)

// Init gin router & middleware
func Init() *gin.Engine {
	r := gin.Default()

	// get global Monitor object
	m := ginmetrics.GetMonitor()

	// +optional set metric path, default /debug/metrics
	m.SetMetricPath("/metrics")
	// +optional set slow time, default 5s
	m.SetSlowTime(10)
	// +optional set request duration, default {0.1, 0.3, 1.2, 5, 10}
	// used to p95, p99
	m.SetDuration([]float64{0.1, 0.3, 1.2, 5, 10})

	// set middleware for gin
	m.Use(r)

	// server status check route
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, model.Response{
			Code:    http.StatusOK,
			Message: "Service Up",
			Data:    make([]interface{}, 0),
		})
	})

	return r
}
