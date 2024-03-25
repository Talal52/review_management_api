// package api

// import (
// 	"template/errors"

// 	"github.com/gin-gonic/gin"
// )

// func (api *TemplateAPIImpl) DelUser(c *gin.Context) error {
// 	id := c.MustGet("userId").(int)
// 	isDeleted, err := api.db.DeltUser(c, id)
// 	if err != nil {
// 		return err
// 	}
// 	if !isDeleted {
// 		errors.Del(c)
// 	}
// 	return nil

// }
package api

import (
	"template/service/models"

	"github.com/gin-gonic/gin"
)

func (api *TemplateAPIImpl) DeleteReviewAPI(c *gin.Context, userId int) (*models.User, error) {
	result, err := api.db.Db_delete_api(c, userId)
	if err != nil {
		return nil, err
	}
	return result, nil
}
