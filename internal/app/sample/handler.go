package sample

import (
	"crud-go-echo-gorm/internal/model"
	res "crud-go-echo-gorm/pkg/util/response"

	"github.com/labstack/echo/v4"
)

type handler struct {
	service Service
}

func NewHandler() *handler {
	service := NewService()
	return &handler{service}
}

// Sample godoc
// @Summary get sample
// @Description Get samplle
// @Tags sample
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @param Authorization header string true "Authorization"
// @Success 200 {object} res.SuccessResponse
// @Failure 400 {object} res.ErrorResponse
// @Failure 404 {object} res.ErrorResponse
// @Failure 500 {object} res.ErrorResponse
// @Router /sample [post]
func (h *handler) GetSamples(c echo.Context) (err error) {
	pagination := new(model.PaginationDTO)

	if err = c.Bind(pagination); err != nil {
		return res.RespError(c, &res.ErrBadRequest)
	}
	if err = c.Validate(pagination); err != nil {
		return res.RespError(c, &res.ErrValidation)
	}

	samples, err := h.service.GetSamples()
	if err != nil {
		return res.RespError(c, err)
	}
	return res.RespSuccess(c, "Get samples success", samples)
}
