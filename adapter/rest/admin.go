package rest

import (
	"../../internal/constant"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func (ps *ProjectService) Register(c *gin.Context) {
	type req struct {
		AdminID  string `json:"adminID" binding:"required"`
		Password string `json:"password" binding:"required"`
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

func (ps *ProjectService) GetUserList(c *gin.Context) {
	page := c.Param("page")

	totalNum, users, err := ps.app.GetUserList(page)
	if err != nil {
		if strings.Contains(err.Error(), constant.NotFoundStr) {
			c.JSON(http.StatusNotFound, gin.H{
				"message": err.Error(),
			})
			return
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"totalNum": totalNum,
		"users": users,
	})
}
