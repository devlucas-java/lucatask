package dto

type AuthDTO struct {
	Token string  `json:"token"`
	User  UserDTO `json:"user"`
}

type LoginDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterDTO struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"min=6"`
}
