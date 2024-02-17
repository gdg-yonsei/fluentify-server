package middleware

import (
	"errors"
	"github.com/gdsc-ys/fluentify-server/src/model"
	"github.com/labstack/echo/v4"
	"net/http"
)

// ErrorDTO TODO: replace with protobuf
type ErrorDTO struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func CustomHTTPErrorHandler(err error, c echo.Context) {
	if c.Response().Committed {
		return
	}

	code := http.StatusInternalServerError
	message := http.StatusText(code)

	var customHTTPError *model.CustomHTTPError
	if errors.As(err, &customHTTPError) {
		code = customHTTPError.Code
		message = customHTTPError.Message
	}

	c.Logger().Error(err)

	if c.Request().Method == http.MethodHead { // Issue #608
		err = c.NoContent(customHTTPError.Code)
	} else {
		err = c.JSON(code, convertToErrorDTO(code, message))
	}
}

func convertToErrorDTO(code int, message string) ErrorDTO {
	return ErrorDTO{
		Code:    code,
		Message: message,
	}
}
