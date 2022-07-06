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

func (s *service) UpdateAktivitas(request dto.AktivitasRequest) (err error) {

	absensiOnDB, err := s.absensiRepository.FindAbsensiByUserID(int(request.UserInfoId))

	if !s.IsCheckInToday(absensiOnDB) {
		return business.ErrForbiddenUpdateAktivitas
	}

	aktivitasOnDB, err := s.repository.FindAktivitasByID(request.ID)
	if err != nil || aktivitasOnDB.ID == 0 {
		return business.ErrDataNotFound
	}

	if aktivitasOnDB.UserInfoID != request.UserInfoId {
		return business.ErrForbiddenAccess
	}

	aktivitasOnDB.UpdatedAt = time.Now()
	aktivitasOnDB.Name = request.Name
	aktivitasOnDB.ID = request.ID

	err = s.repository.UpdateAktivitas(aktivitasOnDB)
	if err != nil {
		return business.ErrDatabase
	}

	return

}

func (s *service) DeleteAktivitas(request dto.AktivitasRequest) (err error) {

	absensiOnDB, err := s.absensiRepository.FindAbsensiByUserID(int(request.UserInfoId))

	if !s.IsCheckInToday(absensiOnDB) {
		return business.ErrForbiddenDeleteAktivitas
	}

	aktivitasOnDB, err := s.repository.FindAktivitasByID(request.ID)
	if err != nil || aktivitasOnDB.ID == 0 {
		return business.ErrDataNotFound
	}

	if aktivitasOnDB.UserInfoID != request.UserInfoId {
		return business.ErrForbiddenAccess
	}

	aktivitasOnDB.UpdatedAt = time.Now()
	aktivitasOnDB.Deleted = true

	err = s.repository.DeleteAktivitas(aktivitasOnDB)
	if err != nil {
		return business.ErrDatabase
	}

	return

}

func (s *service) GetListAktivitas(UserID int) (listAktivitas []dto.GetListAktivitas, err error) {

	listAktivitasOnDB, err := s.repository.FindAllAktivitasByIDUser(UserID)

	listAktivitas = s.convertModelToDTOResponse(listAktivitasOnDB)

	return
}

func (s *service) convertModelToDTOResponse(listAktivitas []Aktivitas) (listAktivitasResponse []dto.GetListAktivitas) {
	formatDate := "2006-01-02"
	for _, aktivitasOnDB := range listAktivitas {
		aktivitasResponse := dto.GetListAktivitas{
			ID:   aktivitasOnDB.ID,
			Name: aktivitasOnDB.Name,
			Date: aktivitasOnDB.CreatedAt.Format(formatDate),
		}

		listAktivitasResponse = append(listAktivitasResponse, aktivitasResponse)
	}

	return
}
