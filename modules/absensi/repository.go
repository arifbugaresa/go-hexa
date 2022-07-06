package absensi

import (
	"github.com/arifbugaresa/go-hexa/business/absensi"
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

func (r *repository) FindAbsensiByUserID(ID int) (absensi absensi.Absensi, err error) {
	err = r.db.Order("id desc").Where("user_info_id = ? AND deleted = FALSE", ID).First(&absensi).Error
	return
}

func (r *repository) InsertCheckInAbsensi(absensi absensi.Absensi) (err error) {
	return r.db.Create(&absensi).Error
}

func (r *repository) UpdateCheckOutAbsensi(absensi absensi.Absensi) (err error) {
	return r.db.Save(&absensi).Error
}

func (r *repository) FindAllAbsensiByIDUser(idUser int) (absensi []absensi.Absensi, err error) {
	err = r.db.Where("user_info_id = ?", idUser).Find(&absensi).Error
	return
}
