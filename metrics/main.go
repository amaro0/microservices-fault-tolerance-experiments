package main

import (
	"fmt"
	"github.com/amaro0/microservices-fault-tolerance-experiments/metrics/data"
	"github.com/gin-gonic/gin"
	"os"
)

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.POST("/metrics", func(c *gin.Context) {
		data := data.Metrics{}

		if err := c.BindJSON(&data); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		fmt.Println(data)

		c.Status(204)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "2000"
	}

	r.Run(":" + port)
}
