package handle

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator"

	"github.com/devlucas-java/lucatask/internal/delivery/dto"
	"github.com/devlucas-java/lucatask/internal/usecase"
	"github.com/go-chi/chi"
)

type TaskHandle struct {
	TaskUseCase usecase.TaskUseCase
}

func NewTaskHandle(t usecase.TaskUseCase) *TaskHandle {
	return &TaskHandle{
		TaskUseCase: t,
	}
}

// CreateTask godoc
// @Summary Create task
// @Description Create a new task
// @Tags tasks
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body dto.TaskDTO true "Task payload"
// @Success 201
// @Failure 400 {string} string "invalid request"
// @Router /tasks/ [post]
func (th *TaskHandle) CreateTask(w http.ResponseWriter, r *http.Request) {
	var dto dto.TaskDTO
	err := json.NewDecoder(r.Body).Decode(&dto)

	validate := validator.New()
	err = validate.Struct(dto)

	if err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}
	th.TaskUseCase.CreateTask(&dto)
	w.WriteHeader(http.StatusCreated)
	return
}

// UpdateTask godoc
// @Summary Update task
// @Description Update task name and description
// @Tags tasks
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Task ID"
// @Param request body dto.TaskDTO true "Task payload"
// @Success 204
// @Failure 400 {string} string "invalid request"
// @Failure 500 {string} string "error updating task"
// @Router /tasks/{id} [put]
func (th *TaskHandle) UpdateTask(w http.ResponseWriter, r *http.Request) {
	var dto dto.TaskDTO
	err := json.NewDecoder(r.Body).Decode(&dto)

	validate := validator.New()
	err = validate.Struct(dto)

	if err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}
	id := chi.URLParam(r, "id")
	err = th.TaskUseCase.UpdateTask(id, &dto)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("error updating task"))
		return
	}
	w.WriteHeader(http.StatusNoContent)
	return
}

// CompletedTask godoc
// @Summary Complete task
// @Description Mark task as completed or incomplete
// @Tags tasks
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Task ID"
// @Param request body dto.TaskCompletedDTO true "Completed payload"
// @Success 204
// @Failure 400 {string} string "invalid request"
// @Failure 500 {string} string "error updating task"
// @Router /tasks/{id}/complete [patch]
func (th *TaskHandle) CompletedTask(w http.ResponseWriter, r *http.Request) {
	var dto dto.TaskCompletedDTO
	err := json.NewDecoder(r.Body).Decode(&dto)

	validate := validator.New()
	err = validate.Struct(dto)

	if err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}
	id := chi.URLParam(r, "id")
	err = th.TaskUseCase.CompletedTask(id, &dto)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("error updating task"))
		return
	}
	w.WriteHeader(http.StatusNoContent)
	return
}

// DeleteTask godoc
// @Summary Delete task
// @Description Delete a task by ID
// @Tags tasks
// @Produce json
// @Security BearerAuth
// @Param id path string true "Task ID"
// @Success 204
// @Failure 500 {string} string "error deleting task"
// @Router /tasks/{id} [delete]
func (th *TaskHandle) DeleteTask(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	err := th.TaskUseCase.DeleteTask(id)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("error deleting task"))
		return
	}
	w.WriteHeader(http.StatusNoContent)
	return
}

// GetTask godoc
// @Summary Get task by ID
// @Description Retrieve a single task
// @Tags tasks
// @Produce json
// @Security BearerAuth
// @Param id path string true "Task ID"
// @Success 200 {object} dto.TaskResponseDTO
// @Failure 500 {string} string "error getting task"
// @Router /tasks/{id} [get]
func (th *TaskHandle) GetTask(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	task, err := th.TaskUseCase.GetTask(id)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("error getting task"))
		return
	}
	json.NewEncoder(w).Encode(task)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	return
}

// GetAllTasks godoc
// @Summary List all tasks
// @Description Get all tasks
// @Tags tasks
// @Produce json
// @Security BearerAuth
// @Success 200 {array} dto.TaskResponseDTO
// @Failure 500 {string} string "error getting tasks"
// @Router /tasks/ [get]
func (th *TaskHandle) GetAllTasks(w http.ResponseWriter, r *http.Request) {

	tasks, err := th.TaskUseCase.ListTasks()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("error getting tasks"))
		return
	}
	json.NewEncoder(w).Encode(tasks)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	return
}
