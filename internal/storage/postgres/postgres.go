package postgres

import (
	"context"
	"errors"
	"fmt"

	models "github.com/ST359/rest-api-todo/internal"
	"github.com/ST359/rest-api-todo/internal/storage"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Storage struct {
	db *pgxpool.Pool
}

func (s *Storage) Close() {
	s.db.Close()
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

func (s *Storage) GetTask(ctx context.Context, id int) (*models.Task, error) {
	const op = "storage.postgres.GetTask"

	var task models.Task
	query := `SELECT * FROM tasks WHERE id=@id`
	args := pgx.NamedArgs{"id": id}
	row, _ := s.db.Query(ctx, query, args)
	task, err := pgx.CollectOneRow(row, pgx.RowToStructByName[models.Task])
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, storage.ErrCantFindTask
		}
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	return &task, nil
}
func (s *Storage) CreateTask(ctx context.Context, task models.Task) (int, error) {
	const op = "storage.postgres.CreateTask"

	var id int
	query := `INSERT INTO tasks (title, description, status, created_at, updated_at) VALUES (@title, @description, @status, @createdAt, @updatedAt) RETURNING id;`
	args := pgx.NamedArgs{
		"title":       task.Title,
		"description": task.Description,
		"status":      task.Status,
		"createdAt":   task.CreatedAt,
		"updatedAt":   task.UpdatedAt,
	}
	err := s.db.QueryRow(ctx, query, args).Scan(&id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return -1, fmt.Errorf("%s: %w", op, err)
		}
		return -1, fmt.Errorf("%s: %w", op, err)
	}
	return id, nil
}
func (s *Storage) UpdateTask(ctx context.Context, task models.Task, id int) error {
	const op = "storage.postgres.UpdateTask"

	query := `UPDATE tasks SET title = COALESCE(@title, title), description = COALESCE(@description, description), status = COALESCE(@status, status), updated_at = @updatedAt WHERE id = @id`
	args := pgx.NamedArgs{
		"id":          id,
		"title":       task.Title,
		"description": task.Description,
		"status":      task.Status,
		"updatedAt":   task.UpdatedAt,
	}
	_, err := s.db.Exec(ctx, query, args)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return storage.ErrCantFindTask
		}
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}
func (s *Storage) DeleteTask(ctx context.Context, id int) error {
	const op = "storage.postgres.DeleteTask"

	query := `DELETE FROM tasks WHERE id = @id`
	args := pgx.NamedArgs{
		"id": id,
	}
	_, err := s.db.Exec(ctx, query, args)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}
func (s *Storage) GetAllTasks(ctx context.Context) ([]*models.Task, error) {
	const op = "storage.postgres.GetAllTasks"

	query := `SELECT * FROM tasks`
	rows, err := s.db.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	tasks, err := pgx.CollectRows(rows, pgx.RowToAddrOfStructByName[models.Task])
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	return tasks, nil
}
