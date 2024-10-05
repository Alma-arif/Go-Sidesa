package users

import (
	"context"
	"errors"
	"fmt"
	"go-sidesa/helper"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type ServiceWeb interface {
	GetAllUsers(ctx context.Context) ([]UserView, error)
	GetUserByID(ctx context.Context, id uint) (User, error)

	RegisterUser(ctx context.Context, input RegisterUserInput) (User, error)
	Login(ctx context.Context, input LoginInput) (User, error)

	GetAllUsersDeleted(ctx context.Context) ([]UserView, error)
	GetUserByIDDeleted(ctx context.Context, id uint) (User, error)
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

			// Format tanggal lahir
			date, _ := helper.DateToFormatIndo(user.TanggalLahir)
			userRow.TanggalLahir = date
			userRow.Role = user.Role

			// Menentukan file profile picture
			if user.ProfileFile == "" {
				userRow.ProfileFile = "image-user-no-poto.png"
			} else {
				userRow.ProfileFile = user.ProfileFile
			}

			// Format waktu create dan update
			timeUserCreate, _ := helper.DatetimeToFormatIndo(user.CreatedAt)
			userRow.CreatedAt = timeUserCreate
			timeUserUpdate, _ := helper.DatetimeToFormatIndo(user.UpdatedAt)
			userRow.UpdatedAt = timeUserUpdate

			userResult = append(userResult, userRow)
		}

		if err := tx.Commit().Error; err != nil {
			return userResult, err
		}
	}

	return userResult, nil
}
func (s *serviceWeb) GetUserByID(ctx context.Context, id uint) (User, error) {
	var user User

	tx := s.db.WithContext(ctx).Begin()
	if err := tx.Error; err != nil {
		return user, err
	}

	select {
	case <-ctx.Done():
		tx.Rollback()
		return user, ctx.Err()
	default:

		err := s.repository.FindByID(&user, id, tx)

		if err != nil {
			return user, err
		}

		date, _ := helper.DateToFormatIndo(user.TanggalLahir)
		user.TanggalLahir = date

		if user.ProfileFile == "" {
			user.ProfileFile = "image-user-no-poto.png"
		}

		if user.ID == 0 {
			return user, errors.New("No user found no with that ID")
		}

		if err := tx.Commit().Error; err != nil {
			return user, err
		}
	}

	return user, nil
}
func (s *serviceWeb) RegisterUser(ctx context.Context, input RegisterUserInput) (User, error) {

	var user User
	var userEmailAfalabel User

	tx := s.db.WithContext(ctx).Begin()
	if err := tx.Error; err != nil {
		return user, err
	}

	select {
	case <-ctx.Done():
		tx.Rollback()
		return user, ctx.Err()
	default:

		err := s.validate.Struct(input)
		if err != nil {
			return user, errors.New("isi form dengan benar!")
		}

		if input.Password != input.PasswordRetype {
			return user, errors.New("Password yang anda masukan salah")
		}

		err = s.repository.FindByEmail(&userEmailAfalabel, input.Email, tx)
		if err != nil {
			return user, err
		}

		if userEmailAfalabel.ID != 0 {
			return user, errors.New("Email sudah pernah digunakan!")
		}

		user.Nama = input.Nama
		user.Email = input.Email
		user.NoHp = input.NoHp
		date, err := helper.StringToDate(input.TanggalLahir)
		if err != nil {
			return user, errors.New("Tangal Lahir tidak sesuai.")
		}
		user.TanggalLahir = date
		user.Password = helper.Sha1ToString(input.Password)
		user.Role = "user"

		err = s.repository.Save(&user, tx)
		if err != nil {
			return user, err
		}

		if err := tx.Commit().Error; err != nil {
			return user, err
		}
	}

	return user, nil
}
func (s *serviceWeb) Login(ctx context.Context, input LoginInput) (User, error) {
	var user User

	tx := s.db.WithContext(ctx).Begin()
	if err := tx.Error; err != nil {
		return user, err
	}

	select {
	case <-ctx.Done():
		tx.Rollback()
		return user, ctx.Err()
	default:
		err := s.validate.Struct(input)
		if err != nil {
			return user, errors.New("isi form dengan benar")
		}

		err = s.repository.FindByEmail(&user, input.Email, tx)
		if err != nil {
			return user, err
		}

		if user.ID == 0 {
			return user, errors.New("Pengguna dengan email tersebut tidak di temaukan")
		}

		ok, err := helper.VerifySHA1Hash(input.Password, user.Password)
		if err != nil {
			return user, errors.New("Password tidak sesuai")
		}

		if ok == false {
			return user, errors.New("Password tidak sesuai")
		}

		if err := tx.Commit().Error; err != nil {
			return user, err
		}
	}

	return user, nil
}
func (s *serviceWeb) GetAllUsersDeleted(ctx context.Context) ([]UserView, error) {
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
		err := s.repository.FindAllDeletedAt(&users, tx)
		if err != nil {
			return userResult, err
		}

		for i, user := range users {
			var userRow UserView
			userRow.ID = user.ID
			userRow.Index = i + 1
			userRow.Nama = user.Nama
			userRow.Email = user.Email
			userRow.NoHp = user.NoHp

			date, _ := helper.DateToFormatIndo(user.TanggalLahir)
			userRow.TanggalLahir = date

			userRow.Role = user.Role
			if user.ProfileFile == "" {
				userRow.ProfileFile = "image-user-no-poto.png"
			} else {
				userRow.ProfileFile = user.ProfileFile
			}

			timeUserCreate, _ := helper.DatetimeToFormatIndo(user.CreatedAt)
			userRow.CreatedAt = timeUserCreate
			timeUserDelete, _ := helper.StringToDateTimeIndoFormat(fmt.Sprint(user.DeletedAt))
			userRow.DeletedAt = timeUserDelete

			userResult = append(userResult, userRow)

		}

		if err := tx.Commit().Error; err != nil {
			return userResult, err
		}
	}

	return userResult, nil
}
func (s *serviceWeb) GetUserByIDDeleted(ctx context.Context, id uint) (User, error) {
	var user User

	tx := s.db.WithContext(ctx).Begin()
	if err := tx.Error; err != nil {
		return user, err
	}

	select {
	case <-ctx.Done():
		tx.Rollback()
		return user, ctx.Err()
	default:

		err := s.repository.FindByIDDeletedAt(&user, id, tx)

		if err != nil {
			return user, err
		}

		date, _ := helper.DateToFormatIndo(user.TanggalLahir)
		user.TanggalLahir = date

		if user.ID == 0 {
			return user, errors.New("No user found no with that ID")
		}

		if err := tx.Commit().Error; err != nil {
			return user, err
		}
	}

	return user, nil
}
