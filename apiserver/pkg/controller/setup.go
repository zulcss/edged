package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/zulcss/edged/apiserver/pkg/api"
)

func Setup() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	
	r.GET("/", api.GetInfo)
	r.GET("/health", api.Health)
	
	r.POST("/site", api.CreateSite)
	r.GET("/site", api.ListSites)
	r.GET("/site/:name", api.GetSite)
	r.DELETE("/site/:name", api.DeleteSite)
	return r
}
