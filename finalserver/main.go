package main

import (
	"github.com/amaro0/microservices-fault-tolerance-experiments/finalserver/config"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type Experiment struct {
	StringToHash string `form:"stringToHash" json:"stringToHash" binding:"required"`
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
		var query Experiment
		if err := c.Bind(&query); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		hashed, err := hash(query.StringToHash)
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

func hash(s string) (hashed string, e error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(s), 9)

	if err != nil {
		return s, err
	}

	return string(bytes), nil
}
