package client

import (
	"fmt"
	"github.com/imroc/req/v3"
)

type Client struct {
	*req.Client
	Host string
}

func NewClient(host string) *Client {
	return &Client{
		Client: req.C(),
		Host: fmt.Sprintf("http://%s:8080", host),
	}
}
