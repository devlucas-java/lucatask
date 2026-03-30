package domain

import (
	"time"

	"github.com/devlucas-java/lucatask/pkg/idgen"
)

type Task struct {
	ID          idgen.ID
	Name        string
	Description string
	Completed   bool
	CreatedAt   time.Time
}

func NewTask(name string, description string) *Task {
	return &Task{
		ID:          idgen.NewID(),
		Name:        name,
		Description: description,
		Completed:   false,
		CreatedAt:   time.Now(),
	}
}
