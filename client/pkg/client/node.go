package client

import "fmt"

type NodeHealth struct {
	Status string `json:"status"`
}

type NodeInfo struct {
	Hostname string `json:"hostname"`
}

func (c *Client) GetStatus() (*NodeHealth, error) {
	node := &NodeHealth{}

	_, err := c.R().
		  SetSuccessResult(&node).
		  Get(fmt.Sprintf("%s/health", c.Host))
	return node, err
}

func (c *Client) GetInfo() (*NodeInfo, error) {
	node := &NodeInfo{}

	_, err := c.R().
		  SetSuccessResult(&node).
		  Get(fmt.Sprintf("%s/", c.Host))
	return node, err
}

