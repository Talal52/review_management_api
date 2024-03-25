package api

// import (
// 	"template/errors"
// 	"template/service/models"

// 	"github.com/gin-gonic/gin"
// )

// func (api *TemplateAPIImpl) Updateuser(c *gin.Context, user *models.UpdateUsr) (*models.UpdateUsr, error) {
// 	userName := c.MustGet("userName").(string)
// 	user.UserName = userName
// 	updatedUser, err := api.db.UpdteUser(c, user)
// 	if err != nil {
// 		errors.BadRequest(c, err)
// 	}
// 	if !updatedUser {
// 		errors.BadRequest(c, err)

// 	}
// 	return user, nil

// }
