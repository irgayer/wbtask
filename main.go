package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/encryptcookie"
	"github.com/gofiber/template/html/v2"
	"github.com/joho/godotenv"
	"wb/database"
	"wb/router"
)

func main() {
	godotenv.Load()

	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})
	app.Use(encryptcookie.New(encryptcookie.Config{
		// Можно было бы и .env вынести.
		Key: "7gv2lnwlut5WcTVeh2O0GgJOU3MeUPVt",
	}))

	database.ConnectDB()

	router.Initialize(app)
	err := app.Listen(":3000")
	if err != nil {
		return
	}
}
