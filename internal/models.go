package models

import "time"

var (
	StatusNew        = "new"
	StatusInProgress = "in_progress"
	StatusDone       = "done"
)

//allows to validate status faster because of o(1) map access
var ValidStatuses = map[string]struct{}{
	StatusNew: {}, StatusInProgress: {}, StatusDone: {},
}

type Task struct {
	ID          int       `json:"id" db:"id"`
	Title       string    `json:"title" db:"title"`
	Description *string   `json:"description" db:"description"`
	Status      string    `json:"status" db:"status"`
	CreatedAt   time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt   time.Time `json:"updatedAt" db:"updated_at"`
}
type TaskRequest struct {
	ID          *int      `json:"id" db:"id"`
	Title       *string   `json:"title" db:"title"`
	Description *string   `json:"description" db:"description"`
	Status      *string   `json:"status" db:"status"`
	CreatedAt   time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt   time.Time `json:"updatedAt" db:"updated_at"`
}

type AllTasksResponce struct {
	Tasks []*Task `json:"tasks"`
}
