package user_info

import "github.com/arifbugaresa/go-hexa/api/v1/user_info/dto"

type Service interface {
	Login(request dto.UserLoginRequest) (userOnDB UserInfo, err error)
}

type Repository interface {
	FindUserInfoByEmail(email string) (user UserInfo, err error)
}