package users

type RegisterUserInput struct {
	Nama           string `json:"nama" form:"nama" validate:"required"`
	Email          string `json:"email" form:"email" validate:"required,email"`
	Password       string `json:"password" form:"password" validate:"required"`
	PasswordRetype string `json:"password-Retype" form:"password-Retype" validate:"required"`
	NoHp           string `json:"no_hp" form:"no_hp" validate:"required"`
	TanggalLahir   string `json:"tanggal_lahir" form:"tanggal_lahir" validate:"required"`
}

type LoginInput struct {
	Email    string `json:"email" form:"email" validate:"required,email"`
	Password string `json:"password" form:"password" validate:"required"`
}
