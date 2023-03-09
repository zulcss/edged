package client

import (
	"fmt"
	"os"
	"time"
	"text/tabwriter"

	"github.com/zulcss/edged/shared/model"
)

func (c *Client) ListSites() {
	sites := []*model.Site{}

	resp, err := c.R().
		 SetSuccessResult(&sites). 
		 Get(fmt.Sprintf("%s/site", c.Host))
	if err != nil {
		return 
	}

	if !resp.IsSuccessState() {
		return 
	}
	w := tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', tabwriter.AlignRight)
	fmt.Fprintf(w, "SITE\t\tCREATED\n")
	for _, name := range sites {
		fmt.Fprintf(w,"%s %s\n", name.Name, name.CreatedAt)
	}
}

func (c *Client) GetSite(site string) {
	var sites *model.Site
	resp, err := c.R().
		     SetPathParam("site", site).
		     SetSuccessResult(&sites).
		     Get(fmt.Sprintf("%s/site/{site}", c.Host))
	if err != nil {
		return 
	}
	if !resp.IsSuccessState() {
		return
	}

	fmt.Printf("%s %s", sites.Name, sites.CreatedAt)
}

func (c *Client) DeleteSite(site string) {
	var sites *model.Site

	resp, err := c.R().
		     SetPathParam("site", site).
		     SetSuccessResult(&sites).
		     Delete(fmt.Sprintf("%s/site/{site}", c.Host))
	if err != nil {
		return 
	}
	if !resp.IsSuccessState() {
		return
	}
	fmt.Printf("Removed %s", site)
}


func (c *Client) CreateSite(site string) {
	var result model.Result
	t := time.Now()
	timestamp := string(t.Format("2006-01-02 15:04:05"))
	resp, err := c.R().
		SetSuccessResult(&result).
		SetBody(&model.Site{Name: site, CreatedAt: timestamp}).
		Post(fmt.Sprintf("%s/site", c.Host))
	if err != nil {
		return 
	}
	if !resp.IsSuccessState() {
		return
	}
	fmt.Printf("Added %s", site)
}
