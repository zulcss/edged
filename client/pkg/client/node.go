package client

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
		  Get("http://0.0.0.0:8080/health")
	return node, err
}

func (c *Client) GetInfo() (*NodeInfo, error) {
	node := &NodeInfo{}

	_, err := c.R().
		  SetSuccessResult(&node).
		  Get("http://0.0.0.0:8080/")
	return node, err
}

