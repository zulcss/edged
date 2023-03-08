package client

import (
	"github.com/imroc/req/v3"
)

type Client struct {
	*req.Client
}

func NewClient() *Client {
	return &Client{
		Client: req.C(),
	}
}
