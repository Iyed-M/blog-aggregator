package main

import (
	"github.com/gofiber/fiber/v3"
)

func respondWithJson(c fiber.Ctx, status int, paylaod interface{}) error {
	return c.Status(status).JSON(paylaod)
}

func respondWithError(c fiber.Ctx, status int, msg string) error {
	return respondWithJson(c, status, map[string]string{"error": msg})
}
