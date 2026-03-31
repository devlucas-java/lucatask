package database

import (
	"github.com/devlucas-java/lucatask/internal/domain"
	"github.com/devlucas-java/lucatask/internal/infra/repository"
	"github.com/devlucas-java/lucatask/pkg/idgen"
	"gorm.io/gorm"
)

type TaskDB struct {
	DB *gorm.DB
}

func NewTaskDB(db *gorm.DB) repository.TaskRepository {
	return &TaskDB{DB: db}
}

func (t *TaskDB) Create(task *domain.Task) error {
	return t.DB.Create(task).Error
}

func (t *TaskDB) Update(task *domain.Task) error {
	return t.DB.Model(&domain.Task{}).Where("id = ?", task.ID).Updates(&task).Error
}

func (t *TaskDB) Delete(id idgen.ID) error {
	return t.DB.Delete(&domain.Task{}, "id = ?", id).Error
}

func (t *TaskDB) FindByID(id idgen.ID) (*domain.Task, error) {
	var task domain.Task
	err := t.DB.Where("id = ?", id).First(&task).Error
	if err != nil {
		return nil, err
	}
	return &task, nil
}

func (t *TaskDB) FindAll() ([]*domain.Task, error) {
	var tasks []*domain.Task
	err := t.DB.Find(&tasks).Error
	if err != nil {
		return nil, err
	}
	return tasks, nil
}
