package main

import (
	"github.com/amaro0/microservices-fault-tolerance-experiments/client/config"
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

		if err != nil {
			log.Println("Request error! ", err.Error())
		}

		err = resp.Body.Close()

		if err != nil {
			log.Fatal(err.Error())
		}
	}

	done <- true
}
