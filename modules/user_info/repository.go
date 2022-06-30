package user_info

import (
	"github.com/arifbugaresa/go-hexa/business/user_info"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{
		db,
	}
}

func (r *repository) FindUserInfoByEmail(email string) (user user_info.UserInfo, err error) {
	err = r.db.Where("email = ?", email).Find(&user).Error
	if err != nil {
		return
	}

	return
}
