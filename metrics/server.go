package metrics

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
)

func RunServer() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.POST("/metrics", func(c *gin.Context) {
		var data Model

		if err := c.ShouldBindJSON(&data); err != nil {
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
