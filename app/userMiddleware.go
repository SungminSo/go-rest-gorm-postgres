package app

import (
	"../internal/constant"
	"../internal/token"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func (app *ProjectApp) UserMiddleware(c *gin.Context) {
	authorization := c.GetHeader("Authorization")
	splitedToken := strings.Split(authorization, "Bearer ")
	if len(splitedToken) != 2 {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": constant.MissingAccessTokenError,
		})
		return
	}

	accesToken := strings.TrimSpace(splitedToken[1])

	// Parse the token
	userToken, err := jwt.ParseWithClaims(accesToken, &token.UserTokenClaim{}, func(userToken *jwt.Token) (interface{},error) {
		return []byte(token.SignKey), nil
	})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": constant.CorruptedJWTError + "\n" + err.Error(),
		})
		return
	}

	claims, ok := userToken.Claims.(*token.UserTokenClaim)
	if !ok || !userToken.Valid {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": constant.InValidJWTError,
		})
		return
	}

	userUUID := claims.UUID
	_, err = app.users.Find(userUUID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.Set("userUUID", userUUID)
	c.Next()
}