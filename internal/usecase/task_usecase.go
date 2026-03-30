package usecase

import (
	"github.com/devlucas-java/lucatask/internal/delivery/dto"
	"github.com/devlucas-java/lucatask/internal/domain"
	"github.com/devlucas-java/lucatask/internal/infra/repository"
	"github.com/devlucas-java/lucatask/pkg/idgen"
)

type TaskUseCase struct {
	TaskRepository repository.TaskRepository
}

func NewTaskUseCase(tr repository.TaskRepository) *TaskUseCase {
	return &TaskUseCase{
		TaskRepository: tr,
	}
}

func (t *TaskUseCase) CreateTask(dto dto.TaskDTO) error {
	task := domain.NewTask(dto.Name, dto.Description)
	return t.TaskRepository.Create(task)
}

func (t *TaskUseCase) GetTask(idRequest string) (*domain.Task, error) {
	id, err := idgen.ParseID(idRequest)
	if err != nil {
		return nil, err
	}
	return t.TaskRepository.FindByID(id)
}

func (t *TaskUseCase) UpdateTask(idRequest string, dto dto.TaskDTO) error {
	id, err := idgen.ParseID(idRequest)
	if err != nil {
		return err
	}
	task, err := t.TaskRepository.FindByID(id)
	if err != nil {
		return err
	}
	task.Name = dto.Name
	task.Description = dto.Description
	return t.TaskRepository.Update(task)
}

func (t *TaskUseCase) DeleteTask(idRequest string) error {
	id, err := idgen.ParseID(idRequest)
	if err != nil {
		return err
	}
	return t.TaskRepository.Delete(id)
}

func (t *TaskUseCase) ListTasks() ([]*domain.Task, error) {
	return t.TaskRepository.FindAll()
}
