package handle

import (
	"encoding/json"
	"net/http"

	"github.com/devlucas-java/lucatask/internal/delivery/dto"
	"github.com/devlucas-java/lucatask/internal/usecase"
	"github.com/go-playground/validator"
)

type UserHandle struct {
	UserUseCase usecase.UserUseCase
}

func NewUserHandle(userUseCase usecase.UserUseCase) UserHandle {
	return UserHandle{
		UserUseCase: userUseCase,
	}
}

func (h *UserHandle) UpdateMe(w http.ResponseWriter, r *http.Request) {
	var dto dto.UserUpdateDTO
	err := json.NewDecoder(r.Body).Decode(&dto)

	validate := validator.New()
	err = validate.Struct(dto)
	if err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	err = h.UserUseCase.UpdateUser("", &dto)
	if err != nil {
		http.Error(w, "failed to update user", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (h *UserHandle) DeleteMe(w http.ResponseWriter, r *http.Request) {
	err := h.UserUseCase.DeleteUser("")
	if err != nil {
		http.Error(w, "failed to delete user", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
