package db

import (
	"fmt"
	"template/service/models"

	"github.com/gin-gonic/gin"
)

func (db *TemplateDBImpl) Db_getreview_api(c *gin.Context, userId int) ([]*models.User, error) {
	var users []*models.User
	err := db.dbConn.Select(&users, "SELECT userid, username, email, message FROM review WHERE userid = $1", userId)
	if err != nil {
		fmt.Println("Error fetching user details:", err)
		return nil, err
	}
	return users, nil
}
