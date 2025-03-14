package service

import (
	"context"
	"testing"

	models "github.com/ST359/rest-api-todo/internal"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/valyala/fasthttp"
)

type MockStorage struct {
	mock.Mock
}

func (m *MockStorage) GetTask(ctx context.Context, id int) (*models.Task, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*models.Task), args.Error(1)
}

func (m *MockStorage) GetAllTasks(ctx context.Context) ([]*models.Task, error) {
	args := m.Called(ctx)
	return args.Get(0).([]*models.Task), args.Error(1)
}

func (m *MockStorage) CreateTask(ctx context.Context, task *models.TaskRequest) (int, error) {
	args := m.Called(ctx, task)
	return args.Int(0), args.Error(1)
}

func (m *MockStorage) UpdateTask(ctx context.Context, task *models.TaskRequest, id int) error {
	args := m.Called(ctx, task, id)
	return args.Error(0)
}

func (m *MockStorage) DeleteTask(ctx context.Context, id int) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func TestServiceGetTask(t *testing.T) {
	mockStorage := new(MockStorage)
	svc := New(mockStorage)
	ctx := fiber.New().AcquireCtx(&fasthttp.RequestCtx{})

	task := &models.Task{ID: 1, Title: "Test Task"}
	mockStorage.On("GetTask", ctx.Context(), 1).Return(task, nil)

	result, err := svc.GetTask(ctx, 1)

	assert.NoError(t, err)
	assert.Equal(t, task, result)
	mockStorage.AssertExpectations(t)
}

func TestServiceGetAllTasks(t *testing.T) {
	mockStorage := new(MockStorage)
	svc := New(mockStorage)
	ctx := fiber.New().AcquireCtx(&fasthttp.RequestCtx{})

	tasks := []*models.Task{{ID: 1, Title: "Test Task 1"}, {ID: 2, Title: "Test Task 2"}}
	mockStorage.On("GetAllTasks", ctx.Context()).Return(tasks, nil)

	result, err := svc.GetAllTasks(ctx)

	assert.NoError(t, err)
	assert.Equal(t, tasks, result)
	mockStorage.AssertExpectations(t)
}

func TestServiceCreateTask(t *testing.T) {
	mockStorage := new(MockStorage)
	svc := New(mockStorage)
	ctx := fiber.New().AcquireCtx(&fasthttp.RequestCtx{})
	title := "New Task"
	taskRequest := &models.TaskRequest{Title: &title}
	mockStorage.On("CreateTask", ctx.Context(), taskRequest).Return(1, nil)

	id, err := svc.CreateTask(ctx, taskRequest)

	assert.NoError(t, err)
	assert.Equal(t, 1, id)
	mockStorage.AssertExpectations(t)
}

func TestServiceCreateTaskInvalidStatus(t *testing.T) {
	mockStorage := new(MockStorage)
	svc := New(mockStorage)
	ctx := fiber.New().AcquireCtx(&fasthttp.RequestCtx{})
	title := "New Task"
	taskRequest := &models.TaskRequest{Title: &title, Status: new(string)}
	*taskRequest.Status = "invalid_status"

	id, err := svc.CreateTask(ctx, taskRequest)

	assert.Error(t, err)
	assert.Equal(t, ErrInvalidTaskStatus, err)
	assert.Equal(t, -1, id)
}

func TestServiceUpdateTask(t *testing.T) {
	mockStorage := new(MockStorage)
	svc := New(mockStorage)
	ctx := fiber.New().AcquireCtx(&fasthttp.RequestCtx{})
	title := "Updated Task"
	taskRequest := &models.TaskRequest{Title: &title}
	mockStorage.On("UpdateTask", ctx.Context(), taskRequest, 1).Return(nil)

	err := svc.UpdateTask(ctx, taskRequest, 1)

	assert.NoError(t, err)
	mockStorage.AssertExpectations(t)
}

func TestServiceUpdateTaskInvalidStatus(t *testing.T) {
	mockStorage := new(MockStorage)
	svc := New(mockStorage)
	ctx := fiber.New().AcquireCtx(&fasthttp.RequestCtx{})
	title := "Updated Task"
	taskRequest := &models.TaskRequest{Title: &title, Status: new(string)}
	*taskRequest.Status = "invalid_status"

	err := svc.UpdateTask(ctx, taskRequest, 1)

	assert.Error(t, err)
	assert.Equal(t, ErrInvalidTaskStatus, err)
}
