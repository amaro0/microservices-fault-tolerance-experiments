package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
)

type MetricsData struct {
	Server      string `json:"server" binding:"required"`
	RequestId   string `json:"requestId" binding:"required"`
	WasError    string `json:"wasError" binding:"required"`
	ErrorTime   int    `json:"errorTime" validate:"min=0"`
	SuccessTime int    `json:"successTime" validate:"min=0"`
	ErrorType   string `json:"errorType"`
}

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.POST("/metrics", func(c *gin.Context) {
		data := MetricsData{}

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
