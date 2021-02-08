package main

import (
	"github.com/amaro0/microservices-fault-tolerance-experiments/proxyserver/config"
	"github.com/gin-gonic/gin"
	"time"
)

func main() {
	serverConfig := config.GetServerConfig()

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/proxy", func(c *gin.Context) {
		time.Sleep(500 * time.Millisecond)

		c.Status(204)
	})

	r.Run(":" + serverConfig.Port)
}
