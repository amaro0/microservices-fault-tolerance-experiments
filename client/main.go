package main

import (
	"github.com/amaro0/microservices-fault-tolerance-experiments/client/config"
	"io/ioutil"
	"log"
	"net/http"
)

var conf *config.ExperimentConfig

func main() {
	conf = config.GetExperimentConfig()

	done := make(chan bool, conf.ConcurrentRequests)

	for i := 0; i <= conf.ConcurrentRequests; i++ {
		go requestContinuously(done)
	}

	<-done
}

func requestContinuously(done chan bool) {
	for i := 0; i <= conf.RequestBatch; i++ {
		resp, err := http.Get(conf.ProxyServerUrl)
		defer resp.Body.Close()

		if err != nil {
			log.Println("Request error! ", err.Error())
		}

		body, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			log.Println("Body parsing error! ", err.Error())
		}

		log.Print(string(body))
	}

	done <- true
}
