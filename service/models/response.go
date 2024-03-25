package models

import "github.com/dgrijalva/jwt-go"

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
type UserDataResponse struct {
	UserId      *int64  `json:"user_id" db:"user_id"`
	Name        *string `json:"name" db:"user_name"`
	Email       *string `json:"email" db:"email"`
	PhoneNumber *string `json:"phone_number" db:"phone_number"`
}

// my struct
type User struct {
	UserID   int    `json:"user_id" db:"userid"`
	UserName string `json:"username" db:"username"`
	Email    string `json:"email" db:"email"`
	Message  string `json:"message" db:"message"`
}

type LoginUser struct {
	UserName     string `json:"username" db:"username"`
	UserPassword string `json:"userpassword" db:"userpassword"`
}
type UpdateUsr struct {
	UserName  string `json:"username" db:"username"`
	FirstName string `json:"firstname" db:"firstname"`
	LastName  string `json:"lastname" db:"lastname"`
}

func NewUserDataResponse() *UserDataResponse {
	return &UserDataResponse{
		UserId:      new(int64),
		Name:        new(string),
		Email:       new(string),
		PhoneNumber: new(string),
	}
}

// private key
var JwtKey = []byte("secretkay")

// payload of jwt token
type JwtClaim struct {
	User_name string `json:"username"`
	UserId    int    `json:"userid"`
	jwt.StandardClaims
}

type HttpResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Token   string      `json:"token"`
	Data    interface{} `json:"data"`
}

type ErrorResponse struct {
	Status int    `json:"status"`
	Err    string `json:"error"`
}
