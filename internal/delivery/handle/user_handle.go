package handle

import (
	"encoding/json"
	"net/http"

	"github.com/devlucas-java/lucatask/internal/delivery/dto"
	"github.com/devlucas-java/lucatask/internal/delivery/middleware"
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
	if err != nil {
		http.Error(w, "Illegal argument", http.StatusBadRequest)
		return
	}
	validate := validator.New()
	err = validate.Struct(dto)
	if err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}
	ctx := r.Context().Value(middleware.AuthKey).(middleware.AuthContext)
	err = h.UserUseCase.UpdateUser(ctx.UserID, &dto)
	if err != nil {
		http.Error(w, "failed to update user", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (h *UserHandle) DeleteMe(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context().Value(middleware.AuthKey).(middleware.AuthContext)
	err := h.UserUseCase.DeleteUser(ctx.UserID)
	if err != nil {
		http.Error(w, "failed to delete user", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (h *UserHandle) GetMe(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context().Value(middleware.AuthKey).(middleware.AuthContext)
	dto, err := h.UserUseCase.GetUser(ctx.UserID)
	if err != nil || dto == nil {
		http.Error(w, "failed to get user", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(dto)
}
