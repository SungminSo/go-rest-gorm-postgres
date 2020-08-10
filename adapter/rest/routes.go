package rest

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (ps *ProjectService) registerRoutes() {
	ps.server.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
		return
	})

	ps.server.GET("version", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"version": ps.config.Version,
		})
		return
	})
}