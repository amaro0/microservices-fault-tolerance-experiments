package finalserver

import (
	"github.com/amaro0/microservices-fault-tolerance-experiments/finalserver/config"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"time"
)

type Experiment struct {
	StringToHash string `form:"stringToHash" json:"stringToHash" binding:"required"`
	RequestId    string `form:"requestId" json:"requestId" binding:"required"`
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

	var experimetStartTime time.Time
	r.GET("/experiment", func(c *gin.Context) {
		if experimetStartTime.IsZero() {
			experimetStartTime = time.Now()
		}
		query := Experiment{}

		if err := c.Bind(&query); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		hashed, err := runExperiment(query, serverConfig, experimetStartTime)
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
