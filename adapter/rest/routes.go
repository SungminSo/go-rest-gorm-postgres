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

	admins := ps.server.Group("/admins")
	{
		// 회원가입
		admins.POST("/registration")
		// 로그인
		admins.POST("/login", )
		// 사용자 목록 조회
		admins.GET("/user", )
		// 사용자 가입 관리 - 승인
		admins.POST("/user", )
		// 사용자 가입 관리 - 거절
		admins.PATCH("/user", )
	}

	users := ps.server.Group("/users")
	{
		// 회원가입
		users.POST("/sign-up", )
		// 로그인
		users.POST("/sign-in", )
		// 사용자 정보 조회
		users.GET("/info", )
		// 사용자 정보 수정
		users.PATCH("/info", )
	}
}