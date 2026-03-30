package dto

type TaskDTO struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
}

type TaskCompletedDTO struct {
	Completed bool `json:"completed" validate:"required"`
}

type TaskResponseDTO struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
	CreatedAt   string `json:"created_at"`
}
