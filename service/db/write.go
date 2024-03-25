package db

// import (
// 	"template/service/models"
// )

// func (db *TemplateDBImpl) CreateUserIdDB(userData *models.UserData) (*models.UserDataResponse, error) {
// 	userIdResp := models.NewUserDataResponse()
// 	tx := db.dbConn.MustBegin()
// 	_, err := tx.NamedQuery(`INSERT INTO review(user_id,username,email,messages)VALUES(:userid,:username,:email,:messages)`, userData)
// 	if err != nil {
// 		return nil, err
// 	}
// 	// err = tx.Commit()
// 	// if err != nil {
// 	// 	return nil, db_error.NewInternalServerError(err.Error())
// 	// }
// 	// err = db.dbConn.Get(userIdResp, `SELECT * FROM userdatabase WHERE user_id=?`, *userData.UserId)
// 	// if err != nil {
// 	// 	return nil, db_error.NewInternalServerError(err.Error())
// 	// }
// 	return userIdResp, nil
// }
