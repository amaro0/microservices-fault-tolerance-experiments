package main

import (
	"amaro0/github.com/microservices-fault-tolerance-experiments/finalserver/config"
	"github.com/gin-gonic/gin"
)

func main() {
	conf := config.NewConfig()
	println(conf)
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/experiment", func(c *gin.Context) {
		c.Status(204)
	})

	r.Run(":" + conf.Port)
}
