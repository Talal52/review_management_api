package api

import (
	"fmt"
	"template/service/models"

	"github.com/gin-gonic/gin"
)

func (api *TemplateAPIImpl) GetReviewByid(c *gin.Context, reviewID int) (*models.User, error) {
	// Call the Db_getreview_api method from the db package
	reviews, err := api.db.Db_getreview_api(c, reviewID)
	if err != nil {
		return nil, err
	}

	// Check if there are any reviews returned
	if len(reviews) == 0 {
		return nil, fmt.Errorf("review not found")
	}

	// Assuming you want to return the first review found
	return reviews[0], nil
}
