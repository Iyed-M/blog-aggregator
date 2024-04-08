package main

import "github.com/gofiber/fiber/v3"

func hanldeReadiness(c fiber.Ctx) error {
	return respondWithJson(c, 200, map[string]string{"status": "ok"})
}

func handleErr(c fiber.Ctx) error {
	return respondWithError(c, 500, "Internal server error")
}
