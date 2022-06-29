package common

import "net/http"

type SuccessResponseCode string

const Success SuccessResponseCode = "success"

type SuccessResponse struct {
	Code    SuccessResponseCode `json:"code"`
	Message string              `json:"message"`
	Data    interface{}         `json:"data"`
}

func NewSuccessResponseWithData(message string, data interface{}) (int, SuccessResponse) {
	return http.StatusOK, SuccessResponse{
		Code:    Success,
		Message: message,
		Data:    data,
	}
}

func NewSuccessResponseWithoutData(message string) (int, SuccessResponse) {
	return http.StatusOK, SuccessResponse{
		Code:    Success,
		Message: message,
		Data:    nil,
	}
}