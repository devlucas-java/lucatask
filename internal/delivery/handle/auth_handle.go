package handle

import (
	"encoding/json"
	"net/http"

	"github.com/devlucas-java/lucatask/internal/delivery/dto"
	"github.com/devlucas-java/lucatask/internal/usecase"
)

type AuthHandle struct {
	AuthUseCase usecase.AuthUseCase
}

func NewAuthHandle(authUseCase usecase.AuthUseCase) AuthHandle {
	return AuthHandle{
		AuthUseCase: authUseCase,
	}
}

func (h *AuthHandle) Login(w http.ResponseWriter, r *http.Request) {

	var dto dto.LoginDTO
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}
	auth, err := h.AuthUseCase.Login(dto.Email, dto.Password)
	if err != nil {
		http.Error(w, "invalid credentials", http.StatusUnauthorized)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(auth)
}

func (h *AuthHandle) Register(w http.ResponseWriter, r *http.Request) {

	var dto dto.RegisterDTO
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}
	auth, err := h.AuthUseCase.Register(&dto)
	if err != nil {
		http.Error(w, "failed to register", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(auth)
}
