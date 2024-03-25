// api/create_user.go
package api

import (
	"template/service/models"

	"github.com/gin-gonic/gin"
)

func (api *TemplateAPIImpl) PostAPI(c *gin.Context, user models.User) (*models.User, error) {

	_, err := api.db.Db_new_api(c, &user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
