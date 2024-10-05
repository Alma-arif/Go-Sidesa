package router

import (
	"go-sidesa/app/users"
	"go-sidesa/config"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"gorm.io/gorm"
)

func SetupRoutes(app *fiber.App, db *gorm.DB, validate *validator.Validate, store *session.Store, conf config.Config) {

	userRepository := users.NewRepository(db)

	users.NewServiceWeb(userRepository, validate, db)

}
