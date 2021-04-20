package repository

import (
	"crud-go-echo-gorm/db"
	"crud-go-echo-gorm/internal/model"

	"gorm.io/gorm"
)

type Sample interface {
	FindAll() ([]model.Sample, error)
}

type sample struct {
	db *gorm.DB
}

func NewSample() *sample {
	db := db.DbManager()
	return &sample{db}
}

func (r *sample) FindAll() ([]model.Sample, error) {
	var datas []model.Sample
	err := r.db.Find(&datas).Error
	if err != nil {
		return datas, err
	}
	return datas, nil
}
