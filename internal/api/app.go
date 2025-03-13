package api

import (
	"fmt"
	"log/slog"
	"time"

	"github.com/ST359/rest-api-todo/internal/config"
	"github.com/ST359/rest-api-todo/internal/service"
	"github.com/ST359/rest-api-todo/internal/storage/postgres"
	"github.com/gofiber/fiber/v2"
)

type Api struct {
	app     *fiber.App
	storage *postgres.Storage
	port    int
	logger  *slog.Logger
}

func New() (*Api, error) {
	logger := slog.New(slog.Default().Handler())
	cfg := config.MustLoad()

	storage, err := postgres.New(cfg)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	svc := service.New(storage)
	handler := NewHandler(svc, logger)

	app := fiber.New(fiber.Config{ServerHeader: "todo-api", ReadTimeout: time.Second * 5})
	setupRoutes(app, handler)

	return &Api{app, storage, cfg.Port, logger}, nil
}
func (a *Api) Run() {
	if err := a.app.Listen(fmt.Sprintf(":%d", a.port)); err != nil {
		a.logger.Error(err.Error())
		panic(err)
	}
}
func (a *Api) Shutdown() error {
	a.storage.Close()
	return a.app.Shutdown()
}
func setupRoutes(app *fiber.App, handler *Handler) {
	app.Get("/tasks", handler.GetAllTasks)
	app.Get("/tasks/:id", handler.GetTask)
	app.Post("/tasks", handler.CreateTask)
	app.Put("/tasks/:id", handler.UpdateTask)
	app.Delete("/tasks/:id", handler.DeleteTask)
}
