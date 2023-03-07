package api

import (
	"github.com/gin-gonic/gin"
	"os"
	"net/http"
)

func GetInfo(c *gin.Context) {
	host, _ := os.Hostname()
	c.IndentedJSON(http.StatusCreated,
		gin.H{"hostname": host})
}