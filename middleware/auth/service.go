package auth

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type jwtService struct {
}

type MyCustomClaims struct {
	UserID int64
	jwt.StandardClaims
}

func NewService() *jwtService {
	return &jwtService{}
}

var SECRET = []byte("super-secret-auth-key")

func (s *jwtService) GenerateToken(userID int64) (tokenJwt string, err error) {

	// generate (header/signing method) + (payload/claim) dulu
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour).Unix()
	claims["user_id"] = userID

	tokenJwt, err = token.SignedString(SECRET)
	fmt.Println(tokenJwt)
	if err != nil {
		return
	}

	return
}

func (s *jwtService) ValidateToken(token string) (jwtToken *jwt.Token, err error) {
	jwtToken, err = jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("invalid token")
		}
		return []byte(SECRET), nil
	})

	if err != nil {
		return
	}

	return
}
