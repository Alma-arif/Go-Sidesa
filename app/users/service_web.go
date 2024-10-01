package users

import (
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type ServiceWeb interface {
}

type serviceWeb struct {
	repository Repository
	validate   *validator.Validate
	db         *gorm.DB
}

func NewServiceWeb(repository Repository, validate *validator.Validate, db *gorm.DB) *serviceWeb {
	return &serviceWeb{repository, validate, db}
}
