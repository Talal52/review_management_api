package cookie

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Store_jwt(c *gin.Context, token string) {
	cookie := http.Cookie{
		Name:     "jwtToken",
		Value:    token,
		HttpOnly: true,
	}
	http.SetCookie(c.Writer, &cookie)
}
