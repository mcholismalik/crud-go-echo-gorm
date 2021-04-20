package route

import (
	_ "crud-go-echo-gorm/docs"
	"crud-go-echo-gorm/internal/app/auth"
	"crud-go-echo-gorm/internal/app/sample"
	"fmt"
	"net/http"
	"os"

	echoSwagger "github.com/swaggo/echo-swagger"

	"github.com/labstack/echo/v4"
)

func Init(g *echo.Group) {
	var (
		APP     = os.Getenv("APP")
		VERSION = os.Getenv("VERSION")
	)

	// Index
	g.GET("/", func(c echo.Context) error {
		message := fmt.Sprintf("Welcome to %s version %s", APP, VERSION)
		return c.String(http.StatusOK, message)
	})

	g.GET("/swagger/*", echoSwagger.WrapHandler)
	// Routes
	auth.NewHandler().Route(g.Group("/auth"))
	sample.NewHandler().Route(g.Group("/samples"))
}
