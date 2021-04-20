package main

import (
	"crud-go-echo-gorm/db"
	cmiddleware "crud-go-echo-gorm/internal/middleware"
	"crud-go-echo-gorm/internal/route"
	"crud-go-echo-gorm/pkg/elastic"
	"crud-go-echo-gorm/pkg/log"
	"crud-go-echo-gorm/pkg/util"
	"fmt"
	"net/http"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
)

func init() {
	ENV := os.Getenv("ENV")
	env := util.NewEnv()
	env.Load(ENV)
	logrus.Info("choose environment " + ENV)
}

// @title crud-go-echo-gorm
// @version 0.0.1
// @description This is a doc for crud-go-echo-gorm.

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @host localhost
// @BasePath /boilerplate
func main() {
	var (
		APP  = os.Getenv("APP")
		ENV  = os.Getenv("ENV")
		PORT = os.Getenv("PORT")
		NAME = fmt.Sprintf("%s-%s", APP, ENV)
	)

	// Init
	db.Init()
	elastic.Init()
	log.Init()
	e := echo.New()

	// Middleware
	e.Use(
		middleware.Logger(),
		middleware.Recover(),
		middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins: []string{"*"},
			AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
		}),
		middleware.LoggerWithConfig(middleware.LoggerConfig{
			Format: fmt.Sprintf("\n%s | ${host} | ${time_custom} | ${status} | ${latency_human} | ${remote_ip} | ${method} | ${uri} ",
				NAME,
			),
			CustomTimeFormat: "2006/01/02 15:04:05",
			Output:           os.Stdout,
		}),
	)
	e.HTTPErrorHandler = cmiddleware.ErrorHandler
	e.Validator = &util.CustomValidator{Validator: validator.New()}

	// Routes
	route.Init(e.Group(""))

	// Start
	e.Logger.Fatal(e.Start(":" + PORT))
}
