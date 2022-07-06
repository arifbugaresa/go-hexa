package aktivitas

import "github.com/arifbugaresa/go-hexa/api/v1/aktivitas/dto"

type Service interface {
	CreateAktivitas(request dto.AktivitasRequest) (err error)
	UpdateAktivitas(request dto.AktivitasRequest) (err error)
	DeleteAktivitas(request dto.AktivitasRequest) (err error)
	GetListAktivitas(UserID int) (listAktivitas []dto.GetListAktivitas, err error)
}

type Repository interface {
	InsertAktivitas(aktivitas Aktivitas) (err error)
	UpdateAktivitas(aktivitas Aktivitas) (err error)
	FindAktivitasByID(id int64) (aktivitasOnDB Aktivitas, err error)
	DeleteAktivitas(aktivitas Aktivitas) (err error)
	FindAllAktivitasByIDUser(idUser int) (aktivitas []Aktivitas, err error)
}
