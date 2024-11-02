package models

import (
	"time"

	"github.com/google/uuid"
)

const TASKS = "tasks"

type tasks struct {
	TaskID    uuid.UUID `json:"taskid"`
	UsetID    uuid.UUID `json:"userid"`
	Task      string    `json:"task"`
	Status    bool      `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
