package absensi

type Service interface {
	CheckIn(ID int) (err error)
}

type Repository interface {
	FindAbsensiByUserID(ID int) (absensi Absensi, err error)
	InsertCheckInAbsensi(absensi Absensi) (err error)
}
