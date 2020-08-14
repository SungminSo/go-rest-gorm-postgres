package token

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

const (
	SignKey = "2nvoi!efg9h0q4+3revf(sdlkf)2nf4g2o=ev3#*p3r2-1v0ew"
)

type AdminTokenClaim struct {
	UUID    string `json:"UUID,omitempty"`
	AdminID string `json:"ID,omitempty"`
	jwt.StandardClaims
}

func GenerateAdminAccessToken(adminUUID, adminID string) string {
	accessToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, AdminTokenClaim{
		UUID:    adminUUID,
		AdminID: adminID,
		StandardClaims: jwt.StandardClaims{
			Issuer:    "Project",
			ExpiresAt: time.Now().Add(time.Hour * 24 * 7 * 2).Unix(), // 2 weeks
		},
	}).SignedString([]byte(SignKey))
	if err != nil {
		return ""
	}

	return accessToken
}

type UserTokenClaim struct {
	UUID  string `json:"UUID,omitempty"`
	Name  string `json:"name,omitempty"`
	Phone string `json:"phone,omitempty"`
	jwt.StandardClaims
}

func GenerateUserAccessToken(userUUID, name, phone string) string {
	accessToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, UserTokenClaim{
		UUID:  userUUID,
		Name:  name,
		Phone: phone,
		StandardClaims: jwt.StandardClaims{
			Issuer:    "Project",
			ExpiresAt: time.Now().Add(time.Hour * 24 * 7 * 2).Unix(), // 2 weeks
		},
	}).SignedString([]byte(SignKey))
	if err != nil {
		return ""
	}

	return accessToken
}
