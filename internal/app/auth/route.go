package auth

import (
	"github.com/labstack/echo/v4"
)

func (h *handler) Route(g *echo.Group) {
	g.POST("/register", h.Register)
	g.POST("/login", h.Login)
}
