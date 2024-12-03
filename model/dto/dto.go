package dto

import "time"

type UserRegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type UserRegisterResponse struct {
	UserId      int64     `json:"userId"`
	Username    string    `json:"username"`
	Email       string    `json:"email"`
	CreateTime  time.Time `json:"create_Time"`
	AccessToken string    `json:"access_token"`
}

type UserLoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserLoginResponse struct {
	UserId      int64     `json:"userId"`
	Username    string    `json:"username"`
	Email       string    `json:"email"`
	CreateTime  time.Time `json:"create_Time"`
	AccessToken string    `json:"access_token"`
}
