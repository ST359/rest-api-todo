package api

import (
	"errors"
	"log/slog"

	models "github.com/ST359/rest-api-todo/internal"
	"github.com/ST359/rest-api-todo/internal/service"
	"github.com/ST359/rest-api-todo/internal/storage"
	"github.com/gofiber/fiber/v2"
)

var (
	InvalidRequestBody  = fiber.Map{"Error": "Unable to parse request body"}
	InternalServerError = fiber.Map{"Error": "Internal server error"}
	NoTitleProvided     = fiber.Map{"Error": "Task title must be provided"}
	WrongTaskId         = fiber.Map{"Error": "Can not find task with provided id"}
	InvalidTaskStatus   = fiber.Map{"Error": "Task status is invalid, use 'new', 'in_progress' or 'done'"}
)

type Handler struct {
	svc    *service.Service
	logger *slog.Logger
}

func NewHandler(svc *service.Service, logger *slog.Logger) *Handler {
	return &Handler{svc, logger}
}

func (h *Handler) CreateTask(ctx *fiber.Ctx) error {
	var taskToCreate *models.TaskRequest
	err := ctx.BodyParser(&taskToCreate)
	if err != nil {
		if errors.Is(err, fiber.ErrUnprocessableEntity) {
			h.logger.Debug(err.Error())
			return ctx.Status(fiber.StatusBadRequest).JSON(InvalidRequestBody)
		}
		h.logger.Error(err.Error())
	}
	if taskToCreate == nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(InternalServerError)
	}
	if taskToCreate.Title == nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(NoTitleProvided)
	}
	id, err := h.svc.CreateTask(ctx, taskToCreate)
	if err != nil {
		if errors.Is(err, service.ErrInvalidTaskStatus) {
			return ctx.Status(fiber.StatusBadRequest).JSON(InvalidTaskStatus)
		}
		h.logger.Error(err.Error())
		return ctx.Status(fiber.StatusInternalServerError).JSON(InternalServerError)
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"id": id})
}

func (h *Handler) UpdateTask(ctx *fiber.Ctx) error {
	var taskToUpdate *models.TaskRequest
	id, err := ctx.ParamsInt("id")
	if err != nil {
		h.logger.Debug(err.Error())
		return ctx.Status(fiber.StatusBadRequest).JSON(WrongTaskId)
	}
	err = ctx.BodyParser(&taskToUpdate)
	if err != nil {
		h.logger.Debug(err.Error())
		return ctx.Status(fiber.StatusBadRequest).JSON(InvalidRequestBody)
	}
	if taskToUpdate == nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(InternalServerError)
	}
	err = h.svc.UpdateTask(ctx, taskToUpdate, id)
	if err != nil {
		if errors.Is(err, service.ErrInvalidTaskStatus) {
			return ctx.Status(fiber.StatusBadRequest).JSON(InvalidTaskStatus)
		}
		if errors.Is(err, storage.ErrCantFindTask) {
			return ctx.Status(fiber.StatusNotFound).JSON(WrongTaskId)
		}
		h.logger.Error(err.Error())
	}
	return ctx.SendStatus(fiber.StatusOK)
}
func (h *Handler) DeleteTask(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(WrongTaskId)
	}
	err = h.svc.DeleteTask(ctx, id)
	if err != nil {
		if errors.Is(err, storage.ErrCantFindTask) {
			return ctx.Status(fiber.StatusBadRequest).JSON(WrongTaskId)
		}
		h.logger.Error(err.Error())
	}
	return ctx.SendStatus(fiber.StatusOK)
}
func (h *Handler) GetTask(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(WrongTaskId)
	}
	task, err := h.svc.GetTask(ctx, id)
	if err != nil {
		if errors.Is(err, storage.ErrCantFindTask) {
			return ctx.Status(fiber.StatusBadRequest).JSON(WrongTaskId)
		}
		h.logger.Error(err.Error())
	}
	return ctx.Status(fiber.StatusOK).JSON(task)
}
func (h *Handler) GetAllTasks(ctx *fiber.Ctx) error {
	tasks, err := h.svc.GetAllTasks(ctx)
	if err != nil {
		if errors.Is(err, storage.ErrCantFindTask) {
			return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"tasks": nil})
		}
		h.logger.Error(err.Error())
	}
	tasksResp := models.AllTasksResponce{Tasks: tasks}
	return ctx.Status(fiber.StatusOK).JSON(tasksResp)
}
