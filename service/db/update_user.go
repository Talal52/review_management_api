package db

// import (
// 	"template/service/models"

// 	"github.com/gin-gonic/gin"
// )

// func (db *TemplateDBImpl) UpdteUser(c *gin.Context, user *models.UpdateUsr) (bool, error) {
// 	userData := "UPDATE users_data SET FirstName =$1,LastName = $2 WHERE userName = $3"
// 	_, err := db.dbConn.Exec(userData, user.FirstName, user.LastName, user.UserName)
// 	if err != nil {
// 		return false, err
// 	}
// 	return true, nil

// }
