package server

import (
	"go-amba/config"
	"go-amba/internal/db"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type Server struct {
	app  *fiber.App
	port string
	cfg  *config.Config
	db   db.DB
}

func NewServer(cfg *config.Config) *Server {
	app := fiber.New(fiber.Config{
		ErrorHandler: fiber.DefaultErrorHandler,
	})

	app.Use(cors.New())

	port := ":" + cfg.Port

	sqlxDB := &db.SqlxDB{DB: cfg.DB}

	return &Server{
		app:  app,
		port: port,
		cfg:  cfg,
		db:   sqlxDB,
	}
}

func (s *Server) Start() error {

	s.SetupRoutes()
	return s.app.Listen(s.port)
}

func (s *Server) Stop() error {

	return s.app.Shutdown()
}
