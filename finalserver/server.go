package finalserver

import (
	"github.com/amaro0/microservices-fault-tolerance-experiments/finalserver/config"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"time"
)

type ErrorType string

const (
	TimeoutError   ErrorType = "timeout"
	UnhandledError ErrorType = "unhandled"
)

type Experiment struct {
	StringToHash     string    `form:"stringToHash" json:"stringToHash" binding:"required"`
	RequestId        string    `form:"requestId" json:"requestId" binding:"required"`
	ErrorRatio       int       `json:"errorRatio" json:"errorRatio" validate:"min=0,max=100"`
	ErrorType        ErrorType `json:"errorType" json:"errorType" validate:"oneof=timeout"'`
	TimeoutLengthInS int       `json:"timeoutLengthInS" json:"timeoutLengthInS"`
}

func RunServer() {
	serverConfig := config.GetServerConfig()

	failServerIfRequired(*serverConfig)

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/experiment", func(c *gin.Context) {
		query := Experiment{ErrorRatio: 50, ErrorType: UnhandledError, TimeoutLengthInS: 30}

		if err := c.Bind(&query); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		hashed, err := runExperiment(query)
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

func failServerIfRequired(serverConfig config.ServerConfig) {
	if serverConfig.ShouldServerFail {
		go func() {
			time.Sleep(30 * time.Second)
			log.Println("SERVER FAIL SIMULATION")
			os.Exit(1)
		}()
	}
}
