package internal

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, DELETE, PUT, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Authorization, Origin, Content-Type, Accept")

		// If it's an OPTIONS request, we respond with the headers and a 200 status
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		// Otherwise, call the next handler
		next.ServeHTTP(w, r)
	})
}

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
