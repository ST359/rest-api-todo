package service

import (
	"context"
	"errors"
	"fmt"
	"time"

	models "github.com/ST359/rest-api-todo/internal"
	"github.com/gofiber/fiber/v2"
)

var ErrInvalidTaskStatus = errors.New("invalid task status")

type Storage interface {
	GetTask(ctx context.Context, id int) (*models.Task, error)
	GetAllTasks(ctx context.Context) ([]*models.Task, error)
	CreateTask(ctx context.Context, task *models.TaskRequest) (int, error)
	UpdateTask(ctx context.Context, task *models.TaskRequest, id int) error
	DeleteTask(ctx context.Context, id int) error
}
type Service struct {
	s Storage
}

func isValidTaskStatus(status string) bool {
	_, ok := models.ValidStatuses[status]
	return ok
}
func New(s Storage) *Service {
	return &Service{s}
}

func (svc *Service) GetTask(ctx *fiber.Ctx, id int) (*models.Task, error) {
	const op = "service.GetTask"

	task, err := svc.s.GetTask(ctx.Context(), id)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	return task, nil
}

func (svc *Service) GetAllTasks(ctx *fiber.Ctx) ([]*models.Task, error) {
	const op = "service.GetAllTasks"

	tasks, err := svc.s.GetAllTasks(ctx.Context())
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	return tasks, nil
}

// CreateTask() sets up CreatedAt and UpdatedAt times for consistency;
//
// Returns an ID of created task
func (svc *Service) CreateTask(ctx *fiber.Ctx, task *models.TaskRequest) (int, error) {
	const op = "service.CreateTask"

	task.CreatedAt, task.UpdatedAt = time.Now(), time.Now()
	if task.Status != nil {
		if !isValidTaskStatus(*task.Status) {
			return -1, ErrInvalidTaskStatus
		}
	} else {
		task.Status = &models.StatusNew
	}
	id, err := svc.s.CreateTask(ctx.Context(), task)
	if err != nil {
		return -1, fmt.Errorf("%s: %w", op, err)
	}
	return id, nil
}

// CreateTask() sets up UpdatedAt time for consistency;
func (svc *Service) UpdateTask(ctx *fiber.Ctx, task *models.TaskRequest, id int) error {
	const op = "service.UpdateTask"

	if task.Status != nil {
		if !isValidTaskStatus(*task.Status) {
			return ErrInvalidTaskStatus
		}
	}
	task.UpdatedAt = time.Now()
	err := svc.s.UpdateTask(ctx.Context(), task, id)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}

func (svc *Service) DeleteTask(ctx *fiber.Ctx, id int) error {
	const op = "service.DeleteTask"

	err := svc.s.DeleteTask(ctx.Context(), id)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}
