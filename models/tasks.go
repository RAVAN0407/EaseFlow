package models

import (
	"time"

	"github.com/google/uuid"
)

const TASKS = "tasks"

type Tasks struct {
	TaskID    uuid.UUID `json:"taskid"`
	UserID    uuid.UUID `json:"userid"`
	Task      string    `json:"task"`
	Status    bool      `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
