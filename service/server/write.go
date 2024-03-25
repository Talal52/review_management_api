package server

import (
	"net/http"
	"template/errors"
	"template/service/models"

	"github.com/gin-gonic/gin"
)

func (s *Server) NewReviewAPI(c *gin.Context) {
	user := models.User{}
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}
	result, err := s.api.PostAPI(c, user)
	if err != nil {
		errors.BadRequest(c, err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	errors.SuccessResponse(c, result)
}
