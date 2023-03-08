package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/zulcss/edged/apiserver/pkg/api"
)

func Setup() *gin.Engine {
	r := gin.Default()
	
	r.GET("/", api.GetInfo)
	r.GET("/health", api.Health)
	
	r.POST("/site", api.CreateSite)
	r.GET("/site", api.ListSites)
	r.GET("/site:id", api.GetSite)
	return r
}
