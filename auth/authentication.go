package auth

import (
	"template/service/models"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateJwtToken(username string, userId int) (jwtTokenString string, err error) {
	claims := models.JwtClaim{
		User_name: username,
		UserId:    userId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(1 * time.Hour).Unix(),
		},
	}
	//jwt signature + payload
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwtTokenString, err = token.SignedString(models.JwtKey)

	return
}
