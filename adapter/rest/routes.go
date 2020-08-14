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
		admins.POST("/register", ps.AdminRegister)
		// 로그인
		admins.POST("/login", ps.Login)

		management := admins.Group("/management")
		management.Use(ps.app.AdminMiddleware)
		{
			// 사용자 가입 목록 조회
			management.GET("", ps.GetUserList)
			// 사용자 가입 관리 - 승인
			management.POST("", ps.ApproveRegistration)
			// 사용자 가입 관리 - 거절
			management.PATCH("", ps.RejectRegistration)
		}
	}

	users := ps.server.Group("/users")
	{
		// 회원가입
		users.POST("/register", ps.UserRegister)
		// 로그인
		users.POST("/sign-in", ps.SignIn)
		// 사용자 정보 조회
		users.GET("/info", )
		// 사용자 정보 수정
		users.PATCH("/info", )
	}
}