package auth

import "github.com/dgrijalva/jwt-go"

type Service interface {
	GenerateToken(userID int64) (token string, err error)
	ValidateToken(token string) (jwt *jwt.Token, err error)
}
