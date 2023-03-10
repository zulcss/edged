package api

import (
	"github.com/gin-gonic/gin"
        "net/http"

        "github.com/zulcss/edged/shared/model"
       // "github.com/zulcss/edged/apiserver/pkg/db"
)

func CreateServer(c *gin.Context) {
	var server model.Server

	if err := c.BindJSON(&server); err != nil {
		return
	}

	c.IndentedJSON(http.StatusCreated,
		gin.H{"Data": "ok"})
}
