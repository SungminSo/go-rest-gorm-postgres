package token

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

const (
	SignKey = "2nvoi!efg9h0q4+3revf(sdlkf)2nf4g2o=ev3#*p3r2-1v0ew"
)

type TokenClaim struct {
	UUID    string `json:"UUID,omitempty"`
	AdminID string `json:"ID,omitempty"`
	jwt.StandardClaims
}

func GenerateAccessToken(adminUUID, adminID string) string {
	accessToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, TokenClaim{
		UUID:    adminUUID,
		AdminID: adminID,
		StandardClaims: jwt.StandardClaims{
			Issuer:    "saenghyeob",
			ExpiresAt: time.Now().Add(time.Hour * 24 * 7 * 2).Unix(), // 2 weeks
		},
	}).SignedString([]byte(SignKey))
	if err != nil {
		return ""
	}

	return accessToken
}