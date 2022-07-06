package absensi

import "github.com/arifbugaresa/go-hexa/api/v1/absensi/dto"

type Service interface {
	CheckIn(User int) (err error)
	CheckOut(UserID int) (err error)
	GetListAbsensi(UserID int) (listAbsensi []dto.GetListAbsensi, err error)
}

type Repository interface {
	FindAbsensiByUserID(ID int) (absensi Absensi, err error)
	InsertCheckInAbsensi(absensi Absensi) (err error)
	UpdateCheckOutAbsensi(absensi Absensi) (err error)
	FindAllAbsensiByIDUser(idUser int) (absensi []Absensi,err error)
}
