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
		admins.POST("/register", ps.Register)
		// 로그인
		admins.POST("/login", )

		application := admins.Group("/application")
		{
			// 사용자 가입 목록 조회
			application.GET("", )
			// 사용자 가입 관리 - 승인
			application.POST("", )
			// 사용자 가입 관리 - 거절
			application.PATCH("", )
		}
	}

	users := ps.server.Group("/users")
	{
		// 회원가입
		users.POST("/register", )
		// 로그인
		users.POST("/sign-in", )
		// 사용자 정보 조회
		users.GET("/info", )
		// 사용자 정보 수정
		users.PATCH("/info", )
	}
}