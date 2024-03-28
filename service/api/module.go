// template/service/api/module.go
package api

import (
	"template/service/db"
	"template/service/models"

	"github.com/gin-gonic/gin"
)

// Import the TemplateDB interface from the db package
type TemplateDB = db.TemplateDB

type TemplateAPI interface {
	PostAPI(c *gin.Context, user models.User) (*models.User, error)
	GetReviewAPI(c *gin.Context) ([]*models.User, error)
	DeleteReviewAPI(c *gin.Context, userId int) (*models.User, error)
	GetReviewByid(c *gin.Context, reviewID int) (*models.User, error)
	// Other methods...
}

type TemplateAPIImpl struct {
	db TemplateDB
}

func NewTemplateAPIImpl() *TemplateAPIImpl {
	dbImpl := db.NewTemplateDBImpl()
	return &TemplateAPIImpl{
		db: dbImpl,
	}
}

var _ TemplateAPI = &TemplateAPIImpl{}
