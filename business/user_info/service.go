package user_info

import (
	"github.com/arifbugaresa/go-hexa/api/v1/user_info/dto"
	"github.com/arifbugaresa/go-hexa/business"
)

type service struct {
	repository Repository
}

// NewService is used to inject repo product to service
func NewService(repository Repository) Service {
	return &service{
		repository,
	}
}

func (s *service) Login(request dto.UserLoginRequest) (userOnDB UserInfo, err error) {
	email := request.Email
	password := request.Password

	userOnDB, err = s.repository.FindUserInfoByEmail(email)
	if err != nil {
		err = business.GenerateErrorQueryDatabase(err)
		return
	}

	if userOnDB.ID == 0 {
		err = business.GenerateErrorDataUserNotFound()
		return
	}

	if password != userOnDB.Password {
		err = business.GenerateErrorEmailAndPasswordMissmatch()
		return
	}

	// todo : jalankan ketika password di db sudah di hash
	//err = bcrypt.CompareHashAndPassword([]byte(userOnDB.Password), []byte(password))
	//if err != nil {
	//	err = business.GenerateErrorEmailAndPasswordMissmatch()
	//	return
	//}

	return
}

func (s *service) GetListUserInfo() (listUserInfo []dto.GetListUserResponse, err error) {

	listUserInfoOnDB, err := s.repository.FindAllUserInfo()
	if err != nil {
		return nil, business.GenerateErrorQueryDatabase(err)
	}

	listUserInfo = s.convertModelToDTOOutForGetList(listUserInfoOnDB)

	return
}

func (s *service) convertModelToDTOOutForGetList(listUserInfoOnDB []UserInfoModel) (listUserInfoResponse []dto.GetListUserResponse) {
	for _, userInfoOnDB := range listUserInfoOnDB {
		userInfoResponse := dto.GetListUserResponse{
			ID:       userInfoOnDB.ID,
			Username: userInfoOnDB.Username,
			Email:    userInfoOnDB.Email,
		}
		listUserInfoResponse = append(listUserInfoResponse, userInfoResponse)
	}
	return
}

