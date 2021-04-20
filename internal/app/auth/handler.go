package auth

import (
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

// Register godoc
// @Summary Register accounts
// @Description Register accounts
// @Tags auth
// @Accept  json
// @Produce  json
// @param register body RegisterDTO true "request body register"
// @Success 200 {object} res.SuccessResponse
// @Failure 400 {object} res.ErrorResponse
// @Failure 404 {object} res.ErrorResponse
// @Failure 500 {object} res.ErrorResponse
// @Router /auth/register [post]
func (h *handler) Register(c echo.Context) (err error) {
	dto := new(RegisterDTO)

	if err = c.Bind(dto); err != nil {
		return res.RespError(c, &res.ErrBadRequest)
	}
	if err = c.Validate(dto); err != nil {

		return res.RespError(c, &res.ErrValidation)
	}

	user, err := h.service.Register(dto)
	if err != nil {

		return res.RespError(c, err)
	}

	return res.RespSuccess(c, "Register success", user)
}

// Login godoc
// @Summary Login accounts
// @Description Login accounts
// @Tags auth
// @Accept  json
// @Produce  json
// @param register body LoginDTO true "request body login"
// @Success 200 {object} LoginResponseDOC
// @Failure 400 {object} res.ErrorResponse
// @Failure 404 {object} res.ErrorResponse
// @Failure 500 {object} res.ErrorResponse
// @Router /auth/login [post]
func (h *handler) Login(c echo.Context) (err error) {
	dto := new(LoginDTO)

	if err = c.Bind(dto); err != nil {

		return res.RespError(c, &res.ErrBadRequest)
	}
	if err = c.Validate(dto); err != nil {

		return res.RespError(c, &res.ErrValidation)
	}

	data, err := h.service.Login(dto)
	if err != nil {

		return res.RespError(c, err)
	}

	return res.RespSuccess(c, "Login success", data)
}
