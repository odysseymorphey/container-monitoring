package server

import (
	"container-monitoring/internal/api/repository"
	"container-monitoring/internal/api/storage"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"os"
	"os/signal"
)

type Server struct {
	server *fiber.App
	router *Router
	repo   repository.Repository
}

func NewServer() *Server {
	db, err := storage.NewPostgresDB(
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
	)
	if err != nil {
		log.Fatal(err)
	}
	log.Fatal(db.Ping())

	app := fiber.New()

	repo := repository.NewRepository(db)
	r := NewRouter(repo)
	r.RegisterRoutes(app)

	return &Server{
		server: app,
		router: r,
		repo:   repo,
	}
}

func (s *Server) Run() {
	log.Fatal(s.server.Listen(":8382"))

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)
	go func() {
		<-sig

		s.Shutdown()
		os.Exit(0)
	}()
}

func (s *Server) Shutdown() {
	log.Fatal(s.repo.Close())
	log.Fatal(s.server.Shutdown())
}
