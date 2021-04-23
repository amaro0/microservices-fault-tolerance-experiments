package finalclient

import (
	"encoding/json"
	"github.com/amaro0/microservices-fault-tolerance-experiments/metrics"
	"github.com/amaro0/microservices-fault-tolerance-experiments/proxyserver/config"
	"github.com/google/uuid"
	"log"
	"net/http"
	"net/url"
	"os"
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
	serverConfig  *config.ServerConfig
	metricsClient *metrics.Client
}

func NewFinalClient(
	serverConfig *config.ServerConfig, metricsClient *metrics.Client,
) *FinalClient {
	return &FinalClient{
		serverConfig,
		metricsClient,
	}
}

func (client *FinalClient) Request(data Data) (Result, error) {
	url, err := createExperimentUrl(client.serverConfig, data)
	if err != nil {
		return Result{}, err
	}

	resp, err := http.Get(url.String())
	if err != nil {
		log.Println(err)
		if os.IsTimeout(err) {
			return Result{}, NewRequestError(TimeoutError, err)
		}

		return Result{}, NewRequestError(UnknownError, err)
	}
	defer resp.Body.Close()

	apiResp := ApiResp{}
	json.NewDecoder(resp.Body).Decode(&apiResp)

	if resp.StatusCode == 500 {
		log.Println(err)
		return Result{}, NewRequestError(UnexpectedError, err)

	}
	if resp.StatusCode >= 400 && resp.StatusCode < 500 {
		log.Println(err)
		return Result{}, NewRequestError(ClientError, err)
	}

	return apiResp.Data, nil
}

func createExperimentUrl(serverConfig *config.ServerConfig, proxyQuery Data) (url.URL, error) {
	base, err := url.Parse(serverConfig.FinalServerUrl)
	if err != nil {
		return *base, err
	}

	query := url.Values{}
	query.Add("stringToHash", uuid.NewString())
	query.Add("requestId", proxyQuery.RequestId)
	base.RawQuery = query.Encode()

	return *base, nil
}
