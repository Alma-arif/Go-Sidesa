package main

import (
	"go-sidesa/componen"
	"go-sidesa/config"
	"go-sidesa/router"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/template/html/v2"
)

func main() {
	conf := config.Get()
	validate := validator.New()
	store := session.New()
	componen.InitDB(conf)

	engine := html.New("./assets/template/", ".html")
	app := fiber.New(fiber.Config{
		Views:     engine,
		BodyLimit: 50 * 1024 * 1024,
	})

	router.SetupRoutes(app, componen.DB, validate, store, conf)

	app.Listen(conf.Srv.Host + ":" + conf.Srv.Port)
}
