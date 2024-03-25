package db

import (
	"fmt"
	"template/service/models"

	"github.com/gin-gonic/gin"
)

func (db *TemplateDBImpl) Db_delete_api(c *gin.Context, userId int) (*models.User, error) {
	res, err := db.dbConn.Exec("DELETE FROM review WHERE userid = $1", userId)
	if err != nil {
		fmt.Println("Error deleting user:", err)
		return nil, err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		fmt.Println("Error getting rows affected:", err)
		return nil, err
	}

	if rowsAffected == 0 {
		return nil, fmt.Errorf("no rows deleted")
	}

	deletedUser := &models.User{
		UserID: userId,
	}
	return deletedUser, nil
}
