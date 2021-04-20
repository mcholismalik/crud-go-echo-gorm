package middleware

import (
	res "crud-go-echo-gorm/pkg/util/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

func ErrorHandler(err error, c echo.Context) {
	report, ok := err.(*echo.HTTPError)
	if !ok {
		report = echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	switch report.Code {
	case http.StatusNotFound:
		err = res.BuildError(res.ErrDataNotFound, err)
	case http.StatusInternalServerError:
		err = res.BuildError(res.ErrServerError, err)
	default:
		err = res.BuildError(res.ErrServerError, err)
	}

	res.RespError(c, err)
}
