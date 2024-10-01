package users

import (
	"gorm.io/gorm"
)

type Repository interface {
	FindAll(users *[]User, tx *gorm.DB) error
	FindByID(user *User, id uint, tx *gorm.DB) error
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
