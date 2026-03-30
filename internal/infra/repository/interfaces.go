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

type UserRepository interface {
	Create(user *domain.User) error
	Update(user *domain.User) error
	Delete(id idgen.ID) error
	FindByID(id idgen.ID) (*domain.User, error)
	FindByEmail(email string) (*domain.User, error)
}
