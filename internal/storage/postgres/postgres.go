package postgres

import (
	"context"
	"fmt"

	models "github.com/ST359/rest-api-todo/internal"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Storage struct {
	db *pgxpool.Pool
}

func New(dbURL string) (*Storage, error) {
	const op = "storage.postgres.New"
	db, err := pgxpool.New(context.Background(), dbURL)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	err = db.Ping(context.Background())
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	return &Storage{db: db}, nil
}

func (s *Storage) CreateTask(models.Task) (*models.Task, error) {
	return nil, nil
}
func (s *Storage) UpdateTask(id int) (*models.Task, error) {
	return nil, nil
}
func (s *Storage) DeleteTask(id int) (*models.Task, error) {
	return nil, nil
}
func (s *Storage) GetTask(id int) (*models.Task, error) {
	return nil, nil
}
func (s *Storage) GetAllTasks() ([]*models.Task, error) {
	return nil, nil
}
