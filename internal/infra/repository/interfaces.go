package repository

import (
	"github.com/devlucas-java/lucatask/internal/domain"
	"github.com/devlucas-java/lucatask/pkg/idgen"
)

type TaskRepository interface {
	Create(task *domain.Task) error
	Update(task *domain.Task) error
	Delete(id idgen.ID) error
	FindByID(id idgen.ID) (*domain.Task, error)
	FindAll() ([]*domain.Task, error)
}
