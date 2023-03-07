package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/zulcss/edged/apiserver/pkg/api"
)

func Setup() *gin.Engine {
	r := gin.Default()
	
	r.GET("/", api.GetInfo)
	r.GET("/health", api.Health)

	return r
}
