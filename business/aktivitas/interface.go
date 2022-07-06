package aktivitas

import "github.com/arifbugaresa/go-hexa/api/v1/aktivitas/dto"

type Service interface {
	CreateAktivitas(request dto.AktivitasRequest) (err error)
	UpdateAktivitas(request dto.AktivitasRequest) (err error)
}

type Repository interface {
	InsertAktivitas(aktivitas Aktivitas) (err error)
	UpdateAktivitas(aktivitas Aktivitas) (err error)
	FindAktivitasByID(id int64) (aktivitasOnDB Aktivitas, err error)
}
