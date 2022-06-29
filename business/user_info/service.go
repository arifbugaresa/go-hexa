package user_info

type service struct {
	repository Repository
}

// NewService is used to inject repo product to service
func NewService(repository Repository) Service {
	return &service{
		repository,
	}
}