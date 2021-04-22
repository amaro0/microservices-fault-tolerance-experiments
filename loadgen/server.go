package loadgen

import (
	"github.com/amaro0/microservices-fault-tolerance-experiments/loadgen/config"
	"github.com/google/uuid"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"
)

var conf *config.ExperimentConfig

// THIS should be rewritten to use two separate gorouteies requester and resolver
// requester will issue requests in constant rate of X per second
// resolver will handle results at constant rate
func Run() {
	conf = config.GetExperimentConfig()

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
	base, err := url.Parse(conf.ProxyServerUrl)
	if err != nil {
		log.Println("Url parsing error")
		return
	}

	query := url.Values{}
	query.Add("requestId", uuid.NewString())
	base.RawQuery = query.Encode()

	resp, err := http.Get(base.String())

	if err != nil {
		log.Println("Request error! ", err.Error())
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

	log.Print(string(body))

	done <- true
}
