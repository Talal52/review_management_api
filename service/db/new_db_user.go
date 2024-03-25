package db

import (
	"template/service/models"

	"github.com/gin-gonic/gin"
)

func (db *TemplateDBImpl) Db_new_api(c *gin.Context, new_user *models.User) (*models.User, error) {
	query := `
        INSERT INTO review (
            userid,
            username,
            email,
            message
        )
        VALUES (
            :userid,
            :username,
            :email,
            :message
        )
    `
	_, err := db.dbConn.NamedExec(query, new_user)
	if err != nil {
		return nil, err
	}
	return new_user, nil
}
