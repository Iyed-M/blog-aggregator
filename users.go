package main

import (
	"context"
	"fmt"
	"time"

	"github.com/Iyed-M/blog-aggregator/internal/database"
	fiber "github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

type req struct {
	Name string `json:"name"`
}

type resp struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
}

func (s *server) handleCreateUser(c fiber.Ctx) error {
	c.Accepts("application/json")
	r := &req{}

	if err := c.Bind().Body(r); err != nil {
		return err
	}

	user, err := s.db.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		Name:      r.Name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	if err != nil {
		return err
	}

	respondWithJson(c, fiber.StatusOK, user)

	fmt.Printf("Received request: %v\n", r.Name)
	return nil
}
