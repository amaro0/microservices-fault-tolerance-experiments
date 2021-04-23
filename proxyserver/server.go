package proxyserver

import (
	"github.com/amaro0/microservices-fault-tolerance-experiments/metrics"
	"github.com/amaro0/microservices-fault-tolerance-experiments/proxyserver/config"
	"github.com/amaro0/microservices-fault-tolerance-experiments/proxyserver/finalclient"
	"github.com/gin-gonic/gin"
	"time"
)

type Result struct {
	Error string             `json:"error"`
	Data  finalclient.Result `json:"data"`
}

func RunServer() {
	serverConfig := config.GetServerConfig()
	metricsClient := metrics.NewClient(serverConfig.MetricsServerUrl)
	finalServerClient := finalclient.NewFinalClient(serverConfig, metricsClient)

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/proxy", func(c *gin.Context) {
		startTime := time.Now()
		var query finalclient.Data
		if err := c.Bind(&query); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		time.Sleep(300 * time.Millisecond)

		metric := metrics.Model{
			Server:    metrics.ProxyServer,
			RequestId: query.RequestId,
			WasError:  false,
		}

		result, err := finalServerClient.Request(query)
		if err != nil {
			re, ok := err.(*finalclient.RequestError)
			if ok {
				metric.ErrorType = re.ErrorType
			}
			metric.WasError = true
			metric.ErrorTime = int(time.Since(startTime) / time.Millisecond)
			metricsClient.SendMetric(metric)

			c.JSON(502, gin.H{
				"error": Result{Error: err.Error()},
			})
			return
		}

		metric.SuccessTime = int(time.Since(startTime) / time.Millisecond)
		metricsClient.SendMetric(metric)
		c.JSON(200, gin.H{
			"data": Result{Data: result},
		})
	})

	r.Run(":" + serverConfig.Port)
}
