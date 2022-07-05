package absensi

type Service interface {
	CheckIn(User int) (err error)
	CheckOut(UserID int) (err error)
}

type Repository interface {
	FindAbsensiByUserID(ID int) (absensi Absensi, err error)
	InsertCheckInAbsensi(absensi Absensi) (err error)
	UpdateCheckOutAbsensi(absensi Absensi) (err error)
}
