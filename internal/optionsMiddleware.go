package internal

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Options(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET,POST,DELETE,OPTIONS")

	if c.Request.Method == "OPTIONS" {
		c.Header("Access-Control-Allow-Headers", "authorization, origin, content-type, accept, access-control-allow-origin")
		c.AbortWithStatus(http.StatusOK)
		return
	}
	c.Next()
}
