package loadgen

import (
	"github.com/amaro0/microservices-fault-tolerance-experiments/loadgen/config"
	"github.com/amaro0/microservices-fault-tolerance-experiments/metrics"
	"github.com/amaro0/microservices-fault-tolerance-experiments/proxyserver/finalclient"
	"github.com/google/uuid"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"
)

var conf *config.ExperimentConfig
var metricsClient *metrics.Client

func Run() {
	conf = config.GetExperimentConfig()
	metricsClient = metrics.NewClient(conf.MetricsServerUrl)

	tickerDone := make(chan bool)
	allRequestsDone := make(chan bool, conf.RequestBatch)
	end := make(chan bool)
	// ticks at Rate per s
	ticker := time.NewTicker(time.Duration(1000/conf.Rate) * time.Millisecond)

	var i int
	go func() {
		for {
			select {
			case <-ticker.C:
				go request(allRequestsDone)
				i++
				if i == conf.RequestBatch {
					tickerDone <- true
				}
			case <-tickerDone:
				ticker.Stop()
				return
			}
		}
	}()

	var endCounter int
	go func() {
		for {
			select {
			case <-allRequestsDone:
				endCounter++
				if endCounter == conf.RequestBatch {
					end <- true
					return
				}
			}
		}
	}()

	<-end
	log.Println("END")
}

func request(done chan bool) {
	startTime := time.Now()
	requestId := uuid.NewString()
	metric := metrics.Model{
		Server:    metrics.LoadGen,
		RequestId: requestId,
		WasError:  false,
	}
	base, err := url.Parse(conf.ProxyServerUrl)
	if err != nil {
		log.Println("Url parsing error")
		return
	}

	query := url.Values{}
	query.Add("requestId", requestId)
	base.RawQuery = query.Encode()

	resp, err := http.Get(base.String())

	if err != nil {
		log.Println("Request error! ", err.Error())
		metric.WasError = true
		metric.ErrorTime = int(time.Since(startTime) / time.Millisecond)
		metric.ErrorType = finalclient.UnknownError
		metricsClient.SendMetric(metric)

		done <- true
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Println("Body parsing error! ", err.Error())
		done <- true
		return
	}

	log.Println(string(body))

	if resp.StatusCode == 502 {
		metric.WasError = true
		metric.ErrorTime = int(time.Since(startTime) / time.Millisecond)
		metric.ErrorType = finalclient.UnexpectedError
		metricsClient.SendMetric(metric)

		done <- true
		return
	}

	metric.SuccessTime = int(time.Since(startTime) / time.Millisecond)
	metricsClient.SendMetric(metric)

	done <- true
}
