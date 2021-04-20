package sample

import (
	mddlwr "crud-go-echo-gorm/internal/middleware"

	"github.com/labstack/echo/v4"
)

func (h *handler) Route(g *echo.Group) {
	g.GET("", h.GetSamples, mddlwr.Authorization)
}
