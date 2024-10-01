package users

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID           uint `gorm:"primaryKey"`
	Nama         string
	Email        string
	Password     string
	NoHp         string
	TanggalLahir time.Time
	Role         string
	ProfileFile  string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}
