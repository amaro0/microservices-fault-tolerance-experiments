package finalclient

import (
	"encoding/json"
	"github.com/amaro0/microservices-fault-tolerance-experiments/metrics"
	"github.com/amaro0/microservices-fault-tolerance-experiments/proxyserver/config"
	"github.com/amaro0/microservices-fault-tolerance-experiments/roundrobin"
	"github.com/google/uuid"
	"github.com/sony/gobreaker"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"
)

type ApiResp struct {
	Error string `json:"error"`
	Data  Result `json:"data"`
}
type Result struct {
	Hashed string `json:"hashed"`
}

type Data struct {
	RequestId string `form:"requestId" json:"requestId" binding:"required"`
}

type FinalClient struct {
	serverConfig            *config.ServerConfig
	metricsClient           *metrics.Client
	roundRobinFinalSelector *roundrobin.Selector
	circuitBreaker          *gobreaker.CircuitBreaker
}

func NewFinalClient(
	serverConfig *config.ServerConfig, metricsClient *metrics.Client,
) *FinalClient {
	validUrls := roundrobin.GetUrlsWithNextPorts(
		serverConfig.FinalServerUrl, serverConfig.FinalInstancesCount,
	)

	return &FinalClient{
		serverConfig,
		metricsClient,
		roundrobin.NewSelector(validUrls),
		gobreaker.NewCircuitBreaker(gobreaker.Settings{
			Name: "Experiment circuit breaker",
		}),
	}
}

func (client *FinalClient) RequestWithStrategy(data Data) (Result, error) {
	if client.serverConfig.ProtectionIncluded(config.CircuitBreaker) {
		// FailRatio probably will need some adjustment to even trigger some circuit breaks in random error generation

		result, err := client.circuitBreaker.Execute(func() (interface{}, error) {
			return client.request(data)
		})

		log.Println("Circuit breaker state: ", client.circuitBreaker.State())

		if err != nil {
			log.Println("Err after circuit breaker: ", err)
			if err == gobreaker.ErrOpenState || err == gobreaker.ErrTooManyRequests {
				log.Println("Circuit breaker early error")
			}
			return Result{}, err
		}

		return result.(Result), nil
	}
	return client.request(data)
}

func (client *FinalClient) request(data Data) (Result, error) {
	url, err := createExperimentUrl(client.roundRobinFinalSelector, data)
	if err != nil {
		return Result{}, err
	}

	httpClient := &http.Client{}
	if client.serverConfig.ProtectionIncluded(config.Timeout) || client.serverConfig.ProtectionIncluded(config.CircuitBreaker) {
		log.Println("dodal timeout")
		httpClient.Timeout = 5 * time.Second
	}

	resp, err := httpClient.Get(url.String())
	if err != nil {
		log.Println("err:  ", err)
		if os.IsTimeout(err) {
			reqErr := NewRequestError(TimeoutError)
			reqErr.AttachError(err)
			return Result{}, reqErr
		}

		reqErr := NewRequestError(UnknownError)
		reqErr.AttachError(err)
		return Result{}, reqErr
	}
	defer resp.Body.Close()

	apiResp := ApiResp{}
	json.NewDecoder(resp.Body).Decode(&apiResp)

	if resp.StatusCode == 500 {
		log.Println("500")
		return Result{}, NewRequestError(UnexpectedError)

	}
	if resp.StatusCode >= 400 && resp.StatusCode < 500 {
		log.Println("Client err with status: ", resp.StatusCode)
		return Result{}, NewRequestError(ClientError)
	}

	return apiResp.Data, nil
}

func createExperimentUrl(rrFinalServerSelector *roundrobin.Selector, proxyQuery Data) (url.URL, error) {
	base, err := url.Parse(rrFinalServerSelector.Get())
	if err != nil {
		return *base, err
	}

	query := url.Values{}
	query.Add("stringToHash", uuid.NewString())
	query.Add("requestId", proxyQuery.RequestId)
	base.RawQuery = query.Encode()

	return *base, nil
}
