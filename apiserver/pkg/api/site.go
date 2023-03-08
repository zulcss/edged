package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"

	"github.com/zulcss/edged/shared/model"
	"github.com/zulcss/edged/apiserver/pkg/db"
)

func CreateSite(c *gin.Context) {
	var site model.Site
	if err := c.BindJSON(&site); err != nil {
		return
	}
	err := db.CreateSite(&site)
	if err != nil {
		fmt.Println("Failed: ", err)
	}
	c.IndentedJSON(http.StatusCreated,
		gin.H{"msg": "ok"})
}

func ListSites(c *gin.Context) {
	var site []model.Site

	err := db.ListSites(&site)
	if err != nil {
		fmt.Println(err)
		c.IndentedJSON(http.StatusCreated,
		gin.H{"msg": "failed"})
	}
	c.JSON(http.StatusCreated, site)
}

func GetSite(c *gin.Context) {
	var site model.Site
	name := c.Params.ByName("name")
	err := db.GetSite(&site, name)
	if err != nil {
		fmt.Println(err)
	}
	
	c.JSON(http.StatusCreated, site)
}
