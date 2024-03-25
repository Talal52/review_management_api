//api/module.go
package api

import (
	"template/service/db"
	"template/service/models"

	"github.com/gin-gonic/gin"
)

type TemplateAPI interface {
	// HelloAPI() string
	PostAPI(c *gin.Context, user models.User) (*models.User, error)
	GetReviewAPI(c *gin.Context) ([]*models.User, error)
	DeleteReviewAPI(c *gin.Context, userId int) (*models.User, error)
	// GetReviewByid(c *gin.Context, userId int) (*models.User, error)
	// LoginUser(c *gin.Context, user *models.LoginUser) (*models.LoginUser, string, error)
	// Updateuser(c *gin.Context, user *models.UpdateUsr) (*models.UpdateUsr, error)
	// DelUser(c *gin.Context) error
}

type TemplateAPIImpl struct {
	db db.TemplateDB
}

func NewTemplateAPIImpl() *TemplateAPIImpl {
	dbImpl := db.NewTemplateDBImpl()
	return &TemplateAPIImpl{
		db: dbImpl,
	}
}

var _ TemplateAPI = &TemplateAPIImpl{}
