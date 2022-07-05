package absensi

import (
	"github.com/arifbugaresa/go-hexa/business"
	"time"
)

type service struct {
	repository Repository
}

// NewService is used to inject repo product to service
func NewService(repository Repository) Service {
	return &service{
		repository,
	}
}

func (s *service) CheckIn(UserID int) (err error) {

	absensiOnDB, err := s.repository.FindAbsensiByUserID(UserID)
	if err != nil && err.Error() != "record not found" {
		return business.ErrGetDataFromDB
	}

	// validasi duplicate checkin
	yearNow, monthNow, dayNow := time.Now().Date()
	yearDB, monthDB, dayDB := absensiOnDB.CheckInTime.Date()
	if yearNow == yearDB && monthNow == monthDB && dayNow == dayDB {
		return business.ErrDuplicateCheckIn
	}

	absensiModel := Absensi{
		UserInfoID:  int64(UserID),
		CheckInTime: time.Now(),
		CheckIn:     true,
		CreatedBy:   "UserID",
		CreatedAt:   time.Now(),
		UpdatedBy:   "UserID",
		UpdatedAt:   time.Now(),
		Deleted:     false,
	}

	err = s.repository.InsertCheckInAbsensi(absensiModel)
	if err != nil {
		return business.GenerateErrorQueryDatabase(err)
	}

	return
}

func (s *service) CheckOut(UserID int) (err error) {

	absensiOnDB, err := s.repository.FindAbsensiByUserID(UserID)
	if err != nil {
		return business.ErrGetDataFromDB
	}

	// validasi checkin bukan hari ini
	yearNow, monthNow, dayNow := time.Now().Date()
	yearDB, monthDB, dayDB := absensiOnDB.CheckInTime.Date()
	if (yearNow != yearDB) || (monthNow != monthDB) || (dayNow != dayDB) {
		return business.ErrForbiddenCheckOut
	}

	absensiOnDB.CheckOutTime = time.Now()
	absensiOnDB.UpdatedAt = time.Now()
	absensiOnDB.CheckOut = true
	err = s.repository.UpdateCheckOutAbsensi(absensiOnDB)
	if err != nil {
		return business.GenerateErrorQueryDatabase(err)
	}

	return
}
