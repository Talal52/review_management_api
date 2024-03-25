package auth

import (
	"fmt"
	"strings"
	"template/errors"
	"template/service/models"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func ValideJwtToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.Request.Header.Get("Authorization")
		tokenStr := strings.Split(tokenString, " ")
		checkBearer := strings.ToUpper(tokenStr[0])
		if checkBearer != "BEARER" {
			errors.Bearer(c)
			return
		}
		isTrueTokenFormat := TokenFormat(tokenStr[1])
		if !isTrueTokenFormat {
			errors.WrongFormat(c)
			return
		}
		if _, err := getTokenFromString(tokenStr[1], c); err != nil {
			errors.UnauthorizedPersons(c)
			return
		}
	}
}

func getTokenFromString(tokenString string, c *gin.Context) (*jwt.Token, error) {
	claims := &models.JwtClaim{}
	return jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method : %v", token.Header["alg"])
		}
		c.Set("userName", claims.User_name)
		c.Set("userId", claims.UserId)
		return []byte(models.JwtKey), nil
	})

}

func TokenFormat(token string) bool {
	count := 0
	for _, char := range token {
		if char == '.' {
			count++
		}
	}
	return count == 2
}
