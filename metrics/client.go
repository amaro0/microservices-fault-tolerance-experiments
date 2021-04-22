package metrics

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type Client struct {
	url string
}

func NewClient(url string) *Client {
	return &Client{url}
}

func (c *Client) SendMetric(metrics Model) {
	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(metrics)

	res, err := http.Post(c.url+"/metrics", "application/json", buf)

	if err != nil {
		log.Panicln("Request error! ", err.Error())
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	if res.StatusCode != 204 {
		log.Panicln("Metric send not successful with status: " + res.Status + " Error req body: " + string(body))
	}
}
