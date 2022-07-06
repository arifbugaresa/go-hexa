package user_info

import "github.com/arifbugaresa/go-hexa/api/v1/user_info/dto"

type Service interface {
	Login(request dto.UserLoginRequest) (userOnDB UserInfo, err error)
	Logout(request dto.UserLogoutRequest) (err error)
	GetListUserInfo() (listUserInfo []dto.GetListUserResponse, err error)
	FindUserInfoByID(ID int) (userOnDB UserInfoModel, err error)
}

type Repository interface {
	FindUserInfoByEmail(email string) (user UserInfo, err error)
	FindAllUserInfo() (listUserInfo []UserInfoModel, err error)
	FindUserInfoByID(ID int) (userOnDB UserInfoModel, err error)
	InsertUser(user UserInfo) (err error)
}
