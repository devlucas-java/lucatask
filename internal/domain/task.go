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
