package db

import (
	"crud-go-echo-gorm/internal/model"
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

func Init() {
	var (
		DB_HOST    = os.Getenv("DB_HOST")
		DB_USER    = os.Getenv("DB_USER")
		DB_PASS    = os.Getenv("DB_PASS")
		DB_NAME    = os.Getenv("DB_NAME")
		DB_PORT    = os.Getenv("DB_PORT")
		DB_SSLMODE = os.Getenv("DB_SSLMODE")
		DB_TZ      = os.Getenv("DB_TZ")
	)

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s", DB_HOST, DB_USER, DB_PASS, DB_NAME, DB_PORT, DB_SSLMODE, DB_TZ)
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	// Migration
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Sample{})
}

func DbManager() *gorm.DB {
	return db
}
