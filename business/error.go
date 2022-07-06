package business

import (
	"errors"
	"github.com/labstack/gommon/log"
)

var (
	ErrInvalidBody              = errors.New("given body cannot be parsed to struct")
	ErrGetDataFromDB            = errors.New("error from db")
	ErrDataNotFound             = errors.New("error data not found")
	ErrEmailAndPasswordMismatch = errors.New("error email and password mismatch")
	ErrInsertData               = errors.New("error insert data")
	ErrDeleteData               = errors.New("error delete data")
	ErrDatabase                 = errors.New("error from database")
	ErrJwt                      = errors.New("error from token jwt invalid")
	ErrUnauthorized             = errors.New("error from token jwt unauthorized")
	ErrDuplicateCheckIn         = errors.New("duplicate check in")
	ErrDuplicateCheckOut        = errors.New("duplicate check Out")
	ErrForbiddenCheckOut        = errors.New("forbidden check Out")
	ErrForbiddenAccess          = errors.New("forbidden access")
	ErrForbiddenCreateAktivitas = errors.New("akses tambah aktivitas tidak diperbolehkan")
	ErrForbiddenUpdateAktivitas = errors.New("akses ubah aktivitas tidak diperbolehkan")
	ErrForbiddenDeleteAktivitas = errors.New("akses menghapus aktivitas tidak diperbolehkan")
)

func GenerateErrorEmailAndPasswordMissmatch() (err error) {
	log.Info(ErrEmailAndPasswordMismatch)
	return ErrEmailAndPasswordMismatch
}

func GenerateErrorDataUserNotFound() (err error) {
	log.Info("Data User Not Found")
	return ErrDataNotFound
}

func GenerateErrorQueryDatabase(pureError error) (err error) {
	log.Error(pureError)
	return ErrDatabase
}
