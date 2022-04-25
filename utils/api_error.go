package utils

import "fmt"

type ApiError struct {
	Status int
	Code string
	Err error
}

func (e ApiError) Error() string {
	return fmt.Sprintf("error: %v", e.Err.Error())
}

func NewApiError(status int, code string, err error) *ApiError {
	return &ApiError{
		Status: status,
		Code: code,
		Err: err,
	}
}