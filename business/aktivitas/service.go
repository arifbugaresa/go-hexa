package aktivitas

import (
	"github.com/arifbugaresa/go-hexa/api/v1/aktivitas/dto"
	"github.com/arifbugaresa/go-hexa/business"
	"github.com/arifbugaresa/go-hexa/business/absensi"
	"time"
)

type service struct {
	repository        Repository
	absensiRepository absensi.Repository
}

// NewService is used to inject repo product to service
func NewService(repository Repository, absensiRepository absensi.Repository) Service {
	return &service{
		repository,
		absensiRepository,
	}
}

func (s *service) CreateAktivitas(request dto.AktivitasRequest) (err error) {

	absensiOnDB, err := s.absensiRepository.FindAbsensiByUserID(int(request.UserInfoId))

	if !s.IsCheckInToday(absensiOnDB) {
		return business.ErrForbiddenCreateAktivitas
	}

	aktivitasModel := s.convertDTOToModelForCreateAktivitas(request)

	err = s.repository.InsertAktivitas(aktivitasModel)
	if err != nil {
		return business.ErrDatabase
	}

	return
}

func (s *service) convertDTOToModelForCreateAktivitas(request dto.AktivitasRequest) Aktivitas {
	return Aktivitas{
		UserInfoID: request.UserInfoId,
		Name:       request.Name,
		CreatedBy:  "",
		CreatedAt:  time.Now(),
		UpdatedBy:  "",
		UpdatedAt:  time.Now(),
		Deleted:    false,
	}
}

func (s *service) IsCheckInToday(absensi absensi.Absensi) bool {
	absensiDate := absensi.CheckInTime.Format("2006-01-02")
	today := time.Now().Format("2006-01-02")

	if absensiDate == today {
		return true
	}

	return false
}
