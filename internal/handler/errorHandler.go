package handler

import (
	"../constant"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func ErrorHandler(c *gin.Context, err error) {
	if strings.Contains(err.Error(), constant.BadRequestStr) {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	} else if strings.Contains(err.Error(), constant.UnauthorizedStr){
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": err.Error(),
		})
	} else if strings.Contains(err.Error(), constant.NotFoundStr) {
		c.JSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
		})
	} else if strings.Contains(err.Error(), constant.ConflictedStr) {
		c.JSON(http.StatusConflict, gin.H{
			"message": err.Error(),
		})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
	}
}
