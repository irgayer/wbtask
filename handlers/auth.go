package handlers

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"strconv"
	"time"
	"wb/database"
	"wb/models"
)

func Login(ctx *fiber.Ctx) error {
	// check cookie session
	// if session exists, redirect to /files
	// else, render login page
	username := ctx.Cookies("session")
	println("COOKIE", username)
	if username != "" {
		return ctx.Redirect("/files")
	}

	db := database.DB
	type LoginInput struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	data := new(LoginInput)
	err := ctx.BodyParser(data)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"code": 500,
		})
	}

	found := models.User{}
	query := models.User{UserName: data.Username}

	err = db.First(&found, &query).Error
	err = db.Where(fmt.Sprintf("user_name = %v", data.Username)).First(&found).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    400,
			"message": "User does not exist",
		})
	}

	if !comparePasswords(found.Password, []byte(data.Password)) {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    400,
			"message": "Incorrect password",
		})
	}

	ctx.Cookie(&fiber.Cookie{
		Name:        "session",
		Value:       strconv.Itoa(int(found.ID)),
		SessionOnly: true,
	})

	return ctx.Status(200).Redirect("/files")
}

func Register(ctx *fiber.Ctx) error {
	db := database.DB
	type RegisterInput struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	data := new(RegisterInput)
	err := ctx.BodyParser(data)
	if err != nil {
		ctx.SendString("Error")
	}

	found := models.User{}
	query := models.User{UserName: data.Username}
	err = db.First(&found, &query).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    400,
			"message": "User already exists",
		})
	}

	new := models.User{
		UserName: data.Username,
		Password: hashAndSalt([]byte(data.Password)),
	}
	db.Create(&new)

	return ctx.Redirect("/login")
}

func hashAndSalt(pwd []byte) string {
	hash, _ := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	return string(hash)
}
func comparePasswords(hashedPwd string, plainPwd []byte) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	return err == nil
}

func SessionExpires() time.Time {
	return time.Now().Add(5 * 24 * time.Hour)
}
