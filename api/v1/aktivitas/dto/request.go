package dto

type AktivitasRequest struct {
	ID         int64
	Name       string `json:"name"`
	UserInfoId int64
}
