package auth

type Service interface {
	GenerateToken(userID int64) (token string, err error)
}
