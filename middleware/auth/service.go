package auth

import (
	"github.com/arifbugaresa/go-hexa/api/v1/common"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
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


func (s *jwtService) GenerateToken(userID int64) (tokenJwt string, err error) {

	// generate (header/signing method) + (payload/claim) dulu
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour).Unix()
	claims["user_id"] = userID

	tokenJwt, err = token.SignedString([]byte(common.SECRET))
	if err != nil {
		return
	}

	return
}

func ValidateJwtMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		authHeader := c.Request().Header.Get("Authorization")

		if !strings.Contains(authHeader, "Bearer") {
			return echo.NewHTTPError(http.StatusForbidden, "Not authorized")
		}

		var tokenString string
		arrToken := strings.Split(authHeader, " ")
		if len(arrToken) == 2 {
			tokenString = arrToken[1]
		}

		// validate token
		jwtToken, err := validateToken(tokenString)

		// bongkar jwt
		claim, ok := jwtToken.Claims.(jwt.MapClaims)
		if !ok || !jwtToken.Valid {
			return echo.NewHTTPError(http.StatusForbidden, "Not authorized")
		}

		userID := int(claim["user_id"].(float64))
		c.Set("currentUser", userID)

		return next(c)
	}
}

func validateToken(token string) (jwtToken *jwt.Token, err error) {
	jwtToken, err = jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, echo.NewHTTPError(http.StatusInternalServerError, "Unable to parse toke")
		}
		return []byte(common.SECRET), nil
	})

	if err != nil {
		return
	}

	return
}
