package main

import (
	"github.com/amaro0/microservices-fault-tolerance-experiments/proxyserver/config"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
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
		time.Sleep(300 * time.Millisecond)

		resp, err := http.Get(serverConfig.FinalServerUrl)

		if err != nil {
			log.Println("Request error! ", err.Error())
		}

		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			c.JSON(500, gin.H{
				"error": err,
			})
		}

		c.JSON(200, gin.H{
			"data": body,
		})
	})

	r.Run(":" + serverConfig.Port)
}
