package db

import (
	"fmt"
	"template/service/models"

	"github.com/gin-gonic/gin"
)

func (db *TemplateDBImpl) Db_get_api(c *gin.Context) ([]*models.User, error) {
	var users []*models.User
	err := db.dbConn.Select(&users, "SELECT userid, username, email, message FROM review")
	if err != nil {
		fmt.Println("Error fetching user details:", err)
		return nil, err
	}
	return users, nil
}
