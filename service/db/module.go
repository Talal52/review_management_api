package db

import (
	"fmt"
	"template/config"
	"template/service/models"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type TemplateDB interface {
	Db_new_api(c *gin.Context, new_user *models.User) (*models.User, error)
	Db_get_api(c *gin.Context) ([]*models.User, error)
	Db_getreview_api(c *gin.Context, userId int) ([]*models.User, error)
	Db_delete_api(c *gin.Context, userId int) (*models.User, error)
}

type TemplateDBImpl struct {
	dbConn *sqlx.DB
}

func NewTemplateDBImpl() *TemplateDBImpl {
	cfg := config.Cfg
	dbconnection := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)

	db, err := sqlx.Open("postgres", dbconnection)
	if err != nil {
		panic(err)
	}
	// defer db.Close()
	err = db.Ping()
	if err != nil {
		fmt.Println("error connecting Database", err)
	}
	return &TemplateDBImpl{
		dbConn: db,
	}
}

var _ TemplateDB = &TemplateDBImpl{}
