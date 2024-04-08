package main

import (
	"database/sql"
	"os"

	"github.com/Iyed-M/blog-aggregator/internal/database"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type apiConfig struct {
	PORT string
}
type server struct {
	apiCfg *apiConfig
	app    *fiber.App
	db     *database.Queries
}

func newSrv() *server {
	godotenv.Load()
	app := fiber.New(fiber.Config{})
	app.Use(logger.New())

	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		panic(err)
	}

	return &server{
		app:    app,
		apiCfg: &apiConfig{PORT: os.Getenv("PORT")},
		db:     database.New(db),
	}
}
