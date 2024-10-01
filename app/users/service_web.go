package users

import (
	"context"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type ServiceWeb interface {
	GetAllUsers(ctx context.Context) ([]UserView, error)
}

type serviceWeb struct {
	repository Repository
	validate   *validator.Validate
	db         *gorm.DB
}

func NewServiceWeb(repository Repository, validate *validator.Validate, db *gorm.DB) *serviceWeb {
	return &serviceWeb{repository, validate, db}
}

func (s *serviceWeb) GetAllUsers(ctx context.Context) ([]UserView, error) {
	var userResult []UserView
	var users []User

	tx := s.db.WithContext(ctx).Begin()
	if err := tx.Error; err != nil {
		return userResult, err
	}

	select {
	case <-ctx.Done():
		tx.Rollback()
		return userResult, ctx.Err()
	default:
		err := s.repository.FindAll(&users, tx)
		if err != nil {
			tx.Rollback()
			return userResult, err
		}

		for i, user := range users {
			var userRow UserView
			userRow.ID = user.ID
			userRow.Index = i + 1
			userRow.Nama = user.Nama
			userRow.Email = user.Email
			userRow.NoHp = user.NoHp
			userRow.Role = user.Role

			if user.ProfileFile == "" {
				userRow.ProfileFile = "image-user-no-poto.png"
			} else {
				userRow.ProfileFile = user.ProfileFile
			}

			userResult = append(userResult, userRow)
		}

		if err := tx.Commit().Error; err != nil {
			return userResult, err
		}
	}

	return userResult, nil
}
