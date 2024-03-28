package errors

import (
	"net/http"
	"template/service/models"

	"github.com/gin-gonic/gin"
)

func SuccessResponse(c *gin.Context, data any) {
	Res := models.HttpResponse{
		Message: "successful",
		Status:  200,
		Data:    data,
	}
	c.AbortWithStatusJSON(http.StatusOK, Res)
}
func SuccessLoginResponse(c *gin.Context, token string, data *models.LoginUser) {
	res := models.HttpResponse{
		Message: " logged In successfully",
		Status:  200,
		Token:   token,
		Data:    data.UserName,
	}
	c.AbortWithStatusJSON(http.StatusOK, res)
}

func BadRequest(c *gin.Context, err error) {
	res := models.HttpResponse{
		Message: "The server could not understand the request due to invalid syntax.",
		Status:  400,
		Data:    err.Error(),
	}
	c.AbortWithStatusJSON(http.StatusBadRequest, res)
}

func UnauthorizedPersons(c *gin.Context) {
	res := models.ErrorResponse{
		Status: 401,
		Err:    "Unauthorized Person",
	}
	c.AbortWithStatusJSON(http.StatusUnauthorized, res)
}
func WrongFormat(c *gin.Context) {
	res := models.ErrorResponse{
		Status: 400,
		Err:    "wrong jwt token format",
	}
	c.AbortWithStatusJSON(http.StatusBadRequest, res)
}
func Bearer(c *gin.Context) {
	res := models.ErrorResponse{
		Status: 400,
		Err:    "bearer does not exist",
	}
	c.AbortWithStatusJSON(http.StatusBadRequest, res)
}

func Del(c *gin.Context) {
	res := models.ErrorResponse{
		Status: 404,
		Err:    "Entered UserName does not exist",
	}
	c.AbortWithStatusJSON(http.StatusNotFound, res)
}
func Deleted(c *gin.Context) {
	res := models.HttpResponse{
		Message: "User Deleted successfully",
		Status:  200,
		Data:    c.Param("username"),
	}
	c.AbortWithStatusJSON(http.StatusOK, res)
}

func InternalServerError(c *gin.Context, message string) {
	c.JSON(http.StatusInternalServerError, gin.H{"error": message})
}
