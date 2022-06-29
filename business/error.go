package business

import "errors"

var (
	ErrInvalidBody   = errors.New("given body cannot be parsed to struct")
	ErrGetDataFromDB = errors.New("error get data from db")
	ErrDataNotFound  = errors.New("error data not found")
	ErrInsertData    = errors.New("error insert data")
	ErrDeleteData    = errors.New("error delete data")
)