package service

import (
	"context"
	"fmt"

	models "github.com/ST359/rest-api-todo/internal"
	"github.com/gofiber/fiber/v2"
)

type Storage interface {
	GetTask(ctx context.Context, id int) (*models.Task, error)
	GetAllTasks(ctx context.Context) ([]*models.Task, error)
	CreateTask(ctx context.Context, task models.Task) (int, error)
	UpdateTask(ctx context.Context, task models.Task, id int) error
	DeleteTask(ctx context.Context, id int) error
}
type Service struct {
	s Storage
}

func New(s Storage) *Service {
	return &Service{s}
}

func (svc *Service) GetTask(ctx *fiber.Ctx, id int) (*models.Task, error) {
	const op = "api.GetTask"

	task, err := svc.s.GetTask(ctx.Context(), id)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	return task, nil
}

func (svc *Service) GetAllTasks(ctx *fiber.Ctx) ([]*models.Task, error) {
	const op = "api.GetAllTasks"

	tasks, err := svc.s.GetAllTasks(ctx.Context())
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	return tasks, nil
}

func (svc *Service) CreateTask(ctx *fiber.Ctx, task models.Task) (int, error) {
	const op = "api.CreateTask"

	id, err := svc.s.CreateTask(ctx.Context(), task)
	if err != nil {
		return -1, fmt.Errorf("%s: %w", op, err)
	}
	return id, nil
}

func (svc *Service) UpdateTask(ctx *fiber.Ctx, task models.Task, id int) error {
	const op = "api.UpdateTask"

	err := svc.s.UpdateTask(ctx.Context(), task, id)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}

func (svc *Service) DeleteTask(ctx *fiber.Ctx, id int) error {
	const op = "api.DeleteTask"

	err := svc.s.DeleteTask(ctx.Context(), id)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}
