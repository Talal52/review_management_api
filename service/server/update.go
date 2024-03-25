package server

// import (
// 	"template/errors"
// 	"template/service/models"

// 	"github.com/gin-gonic/gin"
// )

// func (s *Server) UpdateUser(c *gin.Context) {
// 	var user *models.UpdateUsr
// 	if err := c.ShouldBindJSON(&user); err != nil {
// 		errors.BadRequest(c, err)
// 	}
// 	updatedUser, err := s.api.Updateuser(c, user)
// 	if err != nil {
// 		errors.BadRequest(c, err)
// 	}
// 	errors.SuccessResponse(c, updatedUser)

// }
