package user_info

import "time"

type UserInfo struct {
	ID        int64
	Username  string
	Password  string
	Email     string
	Phone     string
	CreatedBy string
	CreatedAt time.Time
	UpdatedBy string
	UpdatedAt time.Time
	Deleted   bool
}

type UserInfoModel struct {
	ID        int64
	Username  string
	Password  string
	Email     string
	Phone     string
	CreatedBy string
	CreatedAt time.Time
	UpdatedBy string
	UpdatedAt time.Time
	Deleted   bool
}
