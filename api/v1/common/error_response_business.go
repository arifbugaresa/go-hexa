package common

import (
	"github.com/arifbugaresa/go-hexa/business"
	"net/http"
)

type errorBusinessResponseCode string

const (
	errInternalServerError errorBusinessResponseCode = "internal_server_error"
	errHasBeenModified     errorBusinessResponseCode = "data_has_been modified"
	errNotFound            errorBusinessResponseCode = "data_not_found"
	errUnauthorized        errorBusinessResponseCode = "unauthorized"
	errDuplicateCheckIn    errorBusinessResponseCode = "duplicate_check_in"
)

//BusinessResponse default payload response
type BusinessResponse struct {
	Code    errorBusinessResponseCode `json:"code"`
	Message string                    `json:"message"`
	Data    interface{}               `json:"data"`
}

//NewErrorBusinessResponse Response return choosen http status like 400 bad request 422 unprocessable entity, ETC, based on responseCode
func NewErrorBusinessResponse(err error) (int, BusinessResponse) {
	return errorMapping(err)
}

//errorMapping error for missing header key with given value
func errorMapping(err error) (int, BusinessResponse) {
	switch err {
	default:
		return newInternalServerErrorResponse()
	case business.ErrDataNotFound:
		return newNotFoundResponse()
	case business.ErrDatabase:
		return newInternalServerErrorResponse()
	case business.ErrJwt:
		return NewUnauthorizedErrorResponse()
	case business.ErrDatabase:
		return newInternalServerErrorResponse()
	case business.ErrDuplicateCheckIn:
		return NewDuplicateCheckInErrorResponse()
	}
}

//newInternalServerErrorResponse default internal server error response
func newInternalServerErrorResponse() (int, BusinessResponse) {
	return http.StatusInternalServerError, BusinessResponse{
		errInternalServerError,
		"Internal server error",
		map[string]interface{}{},
	}
}

func NewUnauthorizedErrorResponse() (int, BusinessResponse) {
	return http.StatusUnauthorized, BusinessResponse{
		errUnauthorized,
		"Unauthorized",
		map[string]interface{}{},
	}
}

func NewDuplicateCheckInErrorResponse() (int, BusinessResponse) {
	return http.StatusBadRequest, BusinessResponse{
		errDuplicateCheckIn,
		"Duplicate Check In",
		map[string]interface{}{},
	}
}

//newHasBeedModifiedResponse failed to validate request payload
func newHasBeedModifiedResponse() (int, BusinessResponse) {
	return http.StatusBadRequest, BusinessResponse{
		errHasBeenModified,
		"Data has been modified",
		map[string]interface{}{},
	}
}

//newNotFoundResponse default not found error response
func newNotFoundResponse() (int, BusinessResponse) {
	return http.StatusNotFound, BusinessResponse{
		errNotFound,
		"Terjadi kesalahan, data tidak ditemukan",
		map[string]interface{}{},
	}
}
