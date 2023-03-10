package client

import (
//	"os"
	"fmt"
	"github.com/zulcss/edged/shared/model"
)

func (c *Client) CreateServer(Name string, IPAddress string) {
	var result model.Result

	resp, err := c.R().
		     SetSuccessResult(&result).
		     SetBody(&model.Server{Name: Name, Status: "Unregistered"}).
		     Post(fmt.Sprintf("%s/server", c.Host))
	if err != nil {
		return 
	}
	if !resp.IsSuccessState() {
		return
	}
	fmt.Println(resp)
	fmt.Printf("Created %s", Name)
}
