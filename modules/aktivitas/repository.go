package aktivitas

import (
	"github.com/arifbugaresa/go-hexa/business/aktivitas"
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

func (r *repository) InsertAktivitas(aktivitas aktivitas.Aktivitas) (err error) {
	return r.db.Create(&aktivitas).Error
}

func (r *repository) UpdateAktivitas(aktivitas aktivitas.Aktivitas) (err error) {
	return r.db.Save(&aktivitas).Error
}

func (r *repository) FindAktivitasByID(id int64) (aktivitasOnDB aktivitas.Aktivitas, err error) {
	err = r.db.Find(&aktivitasOnDB, id).Error
	return aktivitasOnDB, err
}