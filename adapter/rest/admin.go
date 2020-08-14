package rest

import (
	"../../internal/constant"
	"../../internal/handler"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func (ps *ProjectService) AdminRegister(c *gin.Context) {
	type req struct {
		AdminID  string `json:"adminID" binding:"required"`
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

	adminUUID, err := ps.app.AdminRegister(reqBody.AdminID, reqBody.Password)
	if err != nil {
		handler.ErrorHandler(c, err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"adminUUID": adminUUID,
	})
	return
}

func (ps *ProjectService) Login(c *gin.Context) {
	type req struct {
		AdminID  string `json:"adminID" binding:"required"`
		password string `json:"password" binding:"reuiqred"`
	}

	reqBody := &req{}
	err := c.ShouldBindJSON(reqBody)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	accessToken, err := ps.app.Login(reqBody.AdminID, reqBody.password)
	if err != nil {
		handler.ErrorHandler(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"accessToken": accessToken,
	})
	return
}

func (ps *ProjectService) GetUserList(c *gin.Context) {
	page := c.Param("page")

	totalNum, users, err := ps.app.GetUserList(page)
	if err != nil {
		handler.ErrorHandler(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"totalNum": totalNum,
		"users": users,
	})
}

func (ps *ProjectService) ApproveRegistration(c *gin.Context) {
	type req struct {
		UserUUID string `json:"userUUID" binding:"required"`
	}

	reqBody := &req{}
	err := c.ShouldBindJSON(reqBody)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	userUUID, userStatus, err := ps.app.ApproveRegistration(reqBody.UserUUID)
	if err != nil {
		handler.ErrorHandler(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"userUUID": userUUID,
		"userStatus": userStatus,
	})
	return
}

func (ps *ProjectService) RejectRegistration(c *gin.Context) {
	type req struct {
		UserUUID string `json:"userUUID" binding:"required"`
	}

	reqBody := &req{}
	err := c.ShouldBindJSON(reqBody)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	userUUID, userStatus, err := ps.app.RejectRegistration(reqBody.UserUUID)
	if err != nil {
		handler.ErrorHandler(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"userUUID":   userUUID,
		"userStatus": userStatus,
	})
	return
}