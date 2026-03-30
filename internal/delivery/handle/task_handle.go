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
