package dto

type UserRegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type UserRegisterResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data, omitempty"`
}

type UserLoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
