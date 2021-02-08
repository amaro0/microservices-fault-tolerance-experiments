package main

import (
	"encoding/json"
	"github.com/amaro0/microservices-fault-tolerance-experiments/proxyserver/config"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"
	"net/http"
	"net/url"
	"time"
)

type Result struct {
	Error string            `json:"error"`
	Data  FinalServerResult `json:"data"`
}

type FinalServerResult struct {
	Hashed string `json:"hashed"`
}

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
		url, err := createExperimentUrl(serverConfig)
		if err != nil {
			c.JSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}

		resp, err := http.Get(url.String())
		defer resp.Body.Close()

		if err != nil {
			c.JSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}

		result := Result{}
		json.NewDecoder(resp.Body).Decode(&result)

		if resp.StatusCode > 299 {
			log.Println(result.Error)
			c.JSON(500, gin.H{
				"error": "Experiment returned " + resp.Status,
			})
			return
		}

		c.JSON(200, gin.H{
			"data": result.Data,
		})
	})

	r.Run(":" + serverConfig.Port)
}

func createExperimentUrl(serverConfig *config.ServerConfig) (url.URL, error) {
	base, err := url.Parse(serverConfig.FinalServerUrl)
	if err != nil {
		return *base, err
	}

	query := url.Values{}
	query.Add("stringToHash", uuid.NewString())
	base.RawQuery = query.Encode()

	return *base, nil
}
