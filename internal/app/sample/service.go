package sample

import (
	"context"
	"crud-go-echo-gorm/internal/model"
	"crud-go-echo-gorm/internal/repository"
	"crud-go-echo-gorm/pkg/log"
	res "crud-go-echo-gorm/pkg/util/response"
	"time"

	"github.com/teris-io/shortid"
)

type Service interface {
	GetSamples() ([]model.Sample, error)
}

type service struct {
	sampleRepository repository.Sample
}

func NewService() *service {
	sampleRepository := repository.NewSample()
	return &service{sampleRepository}
}

func (s *service) GetSamples() ([]model.Sample, error) {
	samples, err := s.sampleRepository.FindAll()
	if err != nil {
		return samples, res.BuildError(res.ErrUnprocessableEntity, err)
	}

	log.InsertActivityLog(context.Background(), &log.LogActivity{
		ID:        shortid.MustGenerate(),
		TableName: "users",
		Row:       "nama",
		NewData:   "misal",
		OldData:   "anu",
		CreatedAt: time.Now().Local().UTC(),
	})

	return samples, nil
}
