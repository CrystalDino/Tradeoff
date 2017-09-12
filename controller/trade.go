package controller

import (
	"Tradeoff/core"

	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {
	core.WebEngine.POST("/add", AddTrade)
	core.WebEngine.POST("/cancel", CancelTrade)
}

func AddTrade(c *gin.Context) {
	c.String(http.StatusOK, "-------------")
}

func CancelTrade(c *gin.Context) {
	c.String(http.StatusOK, "-------------")
}

func Run() {
}
