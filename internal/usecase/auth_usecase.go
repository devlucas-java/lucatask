package usecase

import (
	"errors"

	"github.com/devlucas-java/lucatask/internal/delivery/dto"
	"github.com/devlucas-java/lucatask/internal/domain"
	"github.com/devlucas-java/lucatask/internal/infra/repository"
)

type jwtService interface {
	GenerateToken(user *domain.User) (string, error)
}

type AuthUseCase struct {
	UserRepository repository.UserRepository
	JwtService     jwtService
}

func NewAuthUseCase(u repository.UserRepository, j jwtService) AuthUseCase {
	return AuthUseCase{
		UserRepository: u,
		JwtService:     j,
	}
}

func (a *AuthUseCase) Login(email, password string) (*dto.AuthDTO, error) {

	user, err := a.UserRepository.FindByEmail(email)
	if err != nil {
		return nil, err
	}
	if !user.ValidatePassword(password) {
		return nil, errors.New("Invalid credentials")
	}
	userDTO := dto.UserDTO{
		ID:    user.ID.String(),
		Name:  user.Name,
		Email: user.Email,
	}
	jwt, err := a.JwtService.GenerateToken(user)
	if err != nil {
		return nil, err
	}
	return &dto.AuthDTO{
		Token: jwt,
		User:  userDTO,
	}, nil
}

func (a *AuthUseCase) Register(dtoRequest *dto.RegisterDTO) (*dto.AuthDTO, error) {

	userExisting, err := a.UserRepository.FindByEmail(dtoRequest.Email)
	if err == nil && userExisting != nil {
		return nil, errors.New("dupicate email")
	}
	user := domain.NewUser(dtoRequest.Name, dtoRequest.Email, dtoRequest.Password)
	err = a.UserRepository.Create(user)
	if err != nil {
		return nil, err
	}
	userDTO := dto.UserDTO{
		ID:    user.ID.String(),
		Name:  user.Name,
		Email: user.Email,
	}
	jwt, err := a.JwtService.GenerateToken(user)
	if err != nil {
		return nil, err
	}
	return &dto.AuthDTO{
		Token: jwt,
		User:  userDTO,
	}, nil
}
