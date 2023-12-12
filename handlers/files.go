package handlers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"os/exec"
)

func GetFiles(ctx *fiber.Ctx) error {
	// OS Command Injection
	find := ctx.Query("find")
	cmd := exec.Command("powershell.exe", "ls", "public")
	if find != "" {
		cmd = exec.Command("powershell.exe", "ls", "public", find+"*")
	}
	out, err := cmd.Output()

	if err != nil {
		fmt.Println(err)
		return ctx.Status(fiber.StatusInternalServerError).SendString("Error executing ls command")
	}

	return ctx.SendString(string(out))
}

func DownloadFile(ctx *fiber.Ctx) error {
	fileName := ctx.Query("file")

	if fileName == "" {
		return ctx.Status(fiber.StatusBadRequest).SendString("File name not provided in the query parameter")
	}
	// Path Traversal
	filePath := "public/" + fileName

	return ctx.SendFile(filePath)
}
