package app

import (
	"../internal/constant"
	"../internal/token"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func (app *ProjectApp) AdminMiddleware(c *gin.Context) {
	authorization := c.GetHeader("Authorization")
	splitedToken := strings.Split(authorization, "Bearer ")
	if len(splitedToken) != 2 {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": constant.MissingAccessTokenError,
		})
		return
	}

	accessToken := strings.TrimSpace(splitedToken[1])
	// Parse the token
	adminToken, err := jwt.ParseWithClaims(accessToken, &token.AdminTokenClaim{}, func(adminToken *jwt.Token) (interface{}, error) {
		return []byte(token.SignKey), nil
	})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": constant.CorruptedJWTError + "\n" + err.Error(),
		})
		return
	}

	claims, ok := adminToken.Claims.(*token.AdminTokenClaim)
	if !ok || !adminToken.Valid {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": constant.InValidJWTError + "\n" + err.Error(),
		})
		return
	}

	adminUUID := claims.UUID
	_, err = app.admins.Find(adminUUID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.Set("adminUUID", adminUUID)
	c.Next()
}