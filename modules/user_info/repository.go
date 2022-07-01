package user_info

import (
	"fmt"
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

func (r *repository) FindAllUserInfo() (listUserInfo []user_info.UserInfoModel, err error) {
	var temp user_info.UserInfoModel

	query := fmt.Sprintf(`SELECT id, username, email from user_infos`)

	rows, err := r.db.Raw(query).Rows()
	if err != nil {
		return
	}

	defer rows.Close()

	for rows.Next() {
		rows.Scan(&temp.ID, &temp.Username, &temp.Email)
		listUserInfo = append(listUserInfo, temp)
	}

	return
}
