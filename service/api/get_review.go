package api

import (
	"template/service/models"

	"github.com/gin-gonic/gin"
)

func (api *TemplateAPIImpl) GetReviewAPI(c *gin.Context) ([]*models.User, error) {
	result, err := api.db.Db_get_api(c)
	if err != nil {
		return nil, err
	}
	return result, nil
}
