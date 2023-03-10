package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
        "net/http"

        "github.com/zulcss/edged/shared/model"
        "github.com/zulcss/edged/apiserver/pkg/db"
)

func CreateServer(c *gin.Context) {
	var server model.Server

	
	if err := c.BindJSON(&server); err != nil {
		return
	}
	err := db.CreateServer(&server)
	if err != nil {
		fmt.Println("Failed: %s", err)
	}
	c.IndentedJSON(http.StatusCreated,
		gin.H{"Data": "ok"})
}
