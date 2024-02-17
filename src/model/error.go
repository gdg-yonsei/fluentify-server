package model

import "net/http"

type CustomHTTPError struct {
	Code    int
	Message string
}

func (e *CustomHTTPError) Error() string {
	return e.Message
}

func NewCustomHTTPError(code int, errorOrMessage interface{}) *CustomHTTPError {
	message := http.StatusText(http.StatusInternalServerError)
	switch e := errorOrMessage.(type) {
	case string:
		message = e
	case error:
		message = e.Error()
	}

	return &CustomHTTPError{
		Code:    code,
		Message: message,
	}
}
