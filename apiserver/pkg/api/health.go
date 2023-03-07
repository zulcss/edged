package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Health(c *gin.Context) {
	c.IndentedJSON(http.StatusCreated, gin.H{
		"status": "ok"})
}