package handlers

import (
	"github.com/gofiber/fiber/v2"
	"strconv"
	"wb/database"
	"wb/models"
)

func AddComment(ctx *fiber.Ctx) error {
	db := database.DB
	type CommentInput struct {
		Text string `json:"text"`
	}
	data := new(CommentInput)
	err := ctx.BodyParser(data)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    400,
			"message": "Error",
		})
	}

	// get user id from session
	id := ctx.Cookies("session")
	userID, err := strconv.Atoi(id)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    400,
			"message": err,
		})
	}

	db.Create(&models.Comment{Text: data.Text, UserID: uint(userID)})
	return ctx.Status(201).Redirect("/comments")
}

func UpdateComment(ctx *fiber.Ctx) error {
	db := database.DB
	type CommentInput struct {
		Text string `json:"text"`
	}

	data := new(CommentInput)
	// parse request body
	err := ctx.BodyParser(data)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    400,
			"message": "Error",
		})
	}

	id := ctx.Params("id")
	commentID, err := strconv.Atoi(id)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    400,
			"message": err,
		})
	}

	db.Model(&models.Comment{}).Where("id = ?", commentID).Update("text", data.Text)
	return ctx.Status(200).JSON(fiber.Map{
		"code":    200,
		"message": "Comment updated",
	})
}
