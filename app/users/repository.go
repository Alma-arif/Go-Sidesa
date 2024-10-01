package users

import (
	"errors"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll(users *[]User, tx *gorm.DB) error
	FindByID(user *User, id uint, tx *gorm.DB) error
	FindByEmail(user *User, email string, tx *gorm.DB) error

	Save(user *User, tx *gorm.DB) error
	Update(user *User, tx *gorm.DB) error
	DeleteSoft(user *User, id uint, tx *gorm.DB) error
	Delete(user *User, id uint, tx *gorm.DB) error

	FindAllDeletedAt(users *[]User, tx *gorm.DB) error
	FindByIDDeletedAt(user *User, id uint, tx *gorm.DB) error
	UpdateDeletedAt(user *User, id uint, tx *gorm.DB) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll(users *[]User, tx *gorm.DB) error {

	return tx.Find(&users).Error
}

func (r *repository) FindByID(user *User, id uint, tx *gorm.DB) error {

	return tx.Find(&user).Where("id = ?", id).Error
}

func (r *repository) FindByEmail(user *User, email string, tx *gorm.DB) error {

	err := tx.Where("email = ?", email).Find(&user).Error

	if user.ID == 0 {
		return errors.New("Data tidak di temukan.")
	}

	return err
}

func (r *repository) Save(user *User, tx *gorm.DB) error {

	return tx.Create(&user).Error
}

func (r *repository) Update(user *User, tx *gorm.DB) error {

	return tx.Save(&user).Error
}

func (r *repository) DeleteSoft(user *User, id uint, tx *gorm.DB) error {

	return tx.Where("id = ?", id).Delete(&user).Error
}

func (r *repository) Delete(user *User, id uint, tx *gorm.DB) error {

	return tx.Unscoped().Where("id = ?", id).Delete(&user).Error
}

func (r *repository) FindAllDeletedAt(users *[]User, tx *gorm.DB) error {

	return tx.Unscoped().Where("deleted_at > 0").Find(&users).Error
}

func (r *repository) FindByIDDeletedAt(user *User, id uint, tx *gorm.DB) error {

	err := tx.Unscoped().Where("id = ?", id).Find(&user).Error

	if user.ID == 0 {
		return errors.New("Data tidak di temukan.")
	}

	return err
}

func (r *repository) UpdateDeletedAt(user *User, id uint, tx *gorm.DB) error {

	return tx.Unscoped().Model(&user).Where("id", id).Update("deleted_at", nil).Error
}
