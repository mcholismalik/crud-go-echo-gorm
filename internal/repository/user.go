package repository

import (
	"crud-go-echo-gorm/db"
	"crud-go-echo-gorm/internal/model"

	"gorm.io/gorm"
)

type User interface {
	FindAll() ([]model.User, error)
	FindByUsername(username string) (model.User, error)
	Save(data model.User) (model.User, error)
}

type user struct {
	db *gorm.DB
}

func NewUser() *user {
	db := db.DbManager()
	return &user{db}
}

func (r *user) FindAll() ([]model.User, error) {
	var datas []model.User
	err := r.db.Find(&datas).Error
	if err != nil {
		return datas, err
	}
	return datas, nil
}

func (r *user) FindByUsername(username string) (model.User, error) {
	var data model.User
	err := r.db.Where("username = ?", username).First(&data).Error
	if err != nil {
		return data, err
	}

	return data, nil
}

func (r *user) Save(data model.User) (model.User, error) {
	err := r.db.Create(&data).Error
	if err != nil {
		return data, err
	}
	return data, nil
}
