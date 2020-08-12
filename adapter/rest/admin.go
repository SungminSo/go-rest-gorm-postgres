package rest

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (ps *ProjectService) Register(c *gin.Context) {
	type req struct {
		AdminID 	string 	`json:"adminID" binding:"required"`
		Password 	string 	`json:"password" binding:"required"`
	}

	reqBody := req{}
	err := c.ShouldBindJSON(reqBody)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	adminUUID, err := ps.app.Register(reqBody.AdminID, reqBody.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	if len(adminUUID) == 0 {
		c.JSON(http.StatusConflict, gin.H{
			"message": "ID already exists",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"adminUUID": adminUUID,
	})
	return
}