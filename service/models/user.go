package models

type UserData struct {
	UserId   *int64  `json:"user_id" db:"user_id"`
	UserName *string `json:"username" db:"username"`
	Email    *string `json:"email" db:"email"`
	Messages *string `json:"messages" db:"messages"`
}

func NewUserData() *UserData {
	return &UserData{
		UserId:   new(int64),
		UserName: new(string),
		Email:    new(string),
		Messages: new(string),
	}
}
