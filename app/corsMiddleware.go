package app

import "github.com/gin-gonic/gin"

func (app *ProjectApp) CORSMiddleware(c *gin.Context) {
	c.Header("Cache-Control", "no-cache")
	c.Header("Pragma", "no-cache")

	c.Next()
}
