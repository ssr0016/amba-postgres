package server

import (
	"errors"
	"go-amba/internal/db"

	"github.com/gofiber/fiber/v2"
)

func healthCheck(db db.DB) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var result int
		err := db.Get(ctx.Context(), &result, "SELECT 1")
		if err != nil {
			return errors.New("database unavailable")
		}
		return ctx.JSON(fiber.Map{
			"database": "available",
		})
	}
}

func (s *Server) SetupRoutes() {
	api := s.app.Group("/api")
	api.Get("/health", healthCheck(s.db))
}
