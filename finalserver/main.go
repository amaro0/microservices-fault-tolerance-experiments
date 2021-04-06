package main

import (
	"github.com/amaro0/microservices-fault-tolerance-experiments/finalserver/config"
	"github.com/gin-gonic/gin"
)

var (
	ExperimentErrors = struct {
		TimeoutError string
	}{
		TimeoutError: "timeout",
	}
)

type Experiment struct {
	StringToHash     string `form:"stringToHash" json:"stringToHash" binding:"required"`
	RequestId        string `form:"requestId" json:"requestId" binding:"required"`
	ErrorRatio       int    `json:"errorRatio" json:"errorRatio"`
	ErrorType        string `json:"errorType" json:"errorType" validate:"oneof=timeout"'`
	TimeoutLengthInS int    `json:"timeoutLengthInS" json:"timeoutLengthInS"`
}

func main() {
	serverConfig := config.GetServerConfig()

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/experiment", func(c *gin.Context) {
		query := Experiment{ErrorRatio: 50, ErrorType: "timeout", TimeoutLengthInS: 30}

		if err := c.Bind(&query); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		hashed, err := RunExperiment(query)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, gin.H{
			"data": gin.H{
				"hashed": hashed,
			},
		})
	})

	r.Run(":" + serverConfig.Port)
}
