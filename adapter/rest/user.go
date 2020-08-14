package rest

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (ps *ProjectService) UserRegister(c *gin.Context) {
	type req struct {
		Name     string `json:"name" binding:"required"`
		Phone    string `json:"phone" binding:"required"`
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	reqBody := &req{}
	err := c.ShouldBindJSON(reqBody)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	userUUID, err := ps.app.UserRegister(reqBody.Name, reqBody.Phone, reqBody.Email, reqBody.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	if len(userUUID) == 0 {
		c.JSON(http.StatusConflict, gin.H{
			"message": "user already exists",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"userUUID": userUUID,
	})
	return
}

func (ps *ProjectService) SignIn(c *gin.Context) {
	type req struct {
		Phone    string `json:"phone" binding:"required"`
		Password string `json:"password" biniding:"required"`
	}

	reqBody := &req{}
	err := c.ShouldBindJSON(reqBody)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	accessToken, err := ps.app.SignIn(reqBody.Phone, reqBody.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"accessToken": accessToken,
	})
	return
}
