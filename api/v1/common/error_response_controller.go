package common

import "net/http"

type errorControllerResponseCode string

const (
	ErrBadRequest      errorControllerResponseCode = "bad_request"
	ErrForbiddenAccess errorControllerResponseCode = "forbidden"
	ErrBindData        errorControllerResponseCode = "bind_data"
)

type ControllerResponse struct {
	Code    errorControllerResponseCode `json:"code"`
	Message string                      `json:"message"`
	Data    interface{}                 `json:"data"`
}

func NewBadRequestResponse() (int, ControllerResponse) {
	return http.StatusBadRequest, ControllerResponse{
		Code:    ErrBadRequest,
		Message: "Terjadi kesalahan, bad request",
		Data:    nil,
	}
}

func NewErrBindData() (int, ControllerResponse) {
	return http.StatusBadRequest, ControllerResponse{
		Code:    ErrBindData,
		Message: "Terjadi kesalahan, eror bind data karena data tidak sesuai",
		Data:    nil,
	}
}

func NewValidationResponse(data string) (int, ControllerResponse) {
	return http.StatusBadRequest, ControllerResponse{
		ErrBadRequest,
		"Terjadi kesalahan, data " + data + " tidak boleh kosong.",
		map[string]interface{}{},
	}
}