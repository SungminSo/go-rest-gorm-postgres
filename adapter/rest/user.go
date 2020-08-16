package rest

import (
	"../internal/constant"
	"../internal/handler"
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
		handler.ErrorHandler(c, err)
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
		handler.ErrorHandler(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"accessToken": accessToken,
	})
	return
}

func (ps *ProjectService) GetUserInfo(c *gin.Context) {
	userUUID, exists := c.Get("userUUID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": constant.InvalidAccessTokenError,
		})
		return
	}

	user, err := ps.app.GetUserInfo(userUUID.(string))
	if err != nil {
		handler.ErrorHandler(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
	return
}

func (ps *ProjectService) PatchUserInfo(c *gin.Context) {
	userUUID, exists := c.Get("userUUID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": constant.InvalidAccessTokenError,
		})
		return
	}

	type req struct {
		Name  string `json:"name"`
		Phone string `json:"phone"`
		Email string `json:"email"`
	}

	reqBody := &req{}
	err := c.ShouldBindJSON(reqBody)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	user, err := ps.app.PatchUserInfo(userUUID.(string), reqBody.Name, reqBody.Phone, reqBody.Email)
	if err != nil {
		handler.ErrorHandler(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
	return
}
