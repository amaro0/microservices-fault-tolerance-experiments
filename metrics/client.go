package main

import (
	"bytes"
	"encoding/json"
	"github.com/amaro0/micoservices-fault-tolerance-experiments/metrics/data"
	"log"
	"net/http"
)

type Client struct {
	url string
}

func NewClient(url string) *Client {
	return &Client{url}
}

func (c *Client) Request(metrics data.Metrics) {
	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(metrics)

	_, err := http.Post(c.url, "application/json", buf)

	if err != nil {
		log.Panic("Request error! ", err.Error())
	}
}
