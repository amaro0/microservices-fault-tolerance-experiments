package main

import (
	"github.com/amaro0/microservices-fault-tolerance-experiments/client/config"
	"github.com/google/uuid"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
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
		base, err := url.Parse(conf.ProxyServerUrl)
		if err != nil {
			log.Println("Url parsing error")
		}

		query := url.Values{}
		query.Add("requestId", uuid.NewString())
		base.RawQuery = query.Encode()

		resp, err := http.Get(base.String())
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
