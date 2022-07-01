package dto

type UserLoginResponse struct {
	Token string `json:"token"`
}

type GetListUserResponse struct {
	ID       int64
	Username string
	Email    string
}
