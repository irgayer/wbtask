package router

import (
	"github.com/gofiber/fiber/v2"
	"wb/database"
	"wb/handlers"
	"wb/models"
)

func Initialize(router *fiber.App) {
	router.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Render("index", fiber.Map{}, "layouts/main")
	})

	router.Get("/login", func(ctx *fiber.Ctx) error {
		return ctx.Render("login", fiber.Map{})
	})

	router.Get("/register", func(ctx *fiber.Ctx) error {
		return ctx.Render("register", fiber.Map{}, "layouts/main")
	})

	router.Post("/login", handlers.Login)
	router.Post("/register", handlers.Register)

	router.Get("/files", func(ctx *fiber.Ctx) error {
		return ctx.Render("files", fiber.Map{}, "layouts/main")
	})

	router.Get("/get-files", handlers.GetFiles)
	router.Get("/download-file", handlers.DownloadFile)

	router.Get("/comments", func(ctx *fiber.Ctx) error {
		db := database.DB
		var comments []models.Comment
		db.Preload("User").Select("id, text, user_id").Find(&comments)

		for i, c := range comments {
			comments[i].UserName = c.User.UserName
		}

		return ctx.Render("comments", fiber.Map{
			"comments": comments,
		}, "layouts/main")
	})
	router.Patch("/comments/:id", handlers.UpdateComment)
	router.Post("/comments", handlers.AddComment)
}
