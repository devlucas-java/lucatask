package usecase

import (
	"github.com/devlucas-java/lucatask/internal/delivery/dto"
	"github.com/devlucas-java/lucatask/internal/domain"
	"github.com/devlucas-java/lucatask/internal/infra/repository"
	"github.com/devlucas-java/lucatask/pkg/idgen"
	"golang.org/x/crypto/bcrypt"
)

type UserUseCase struct {
	UserRepository repository.UserRepository
}

func NewUserUseCase(u repository.UserRepository) *UserUseCase {
	return &UserUseCase{
		UserRepository: u,
	}
}

func (u *UserUseCase) CreateUser(name, email, password string) error {
	return u.UserRepository.Create(domain.NewUser(name, email, password))
}

func (u *UserUseCase) GetUserByEmail(email string) (*dto.UserDTO, error) {
	user, err := u.UserRepository.FindByEmail(email)
	if err != nil {
		return nil, err
	}
	dto := &dto.UserDTO{
		ID:    user.ID.String(),
		Name:  user.Name,
		Email: user.Email,
	}
	return dto, nil
}

func (u *UserUseCase) GetUserByID(idRequest string) (*dto.UserDTO, error) {
	id, err := idgen.ParseID(idRequest)
	user, err := u.UserRepository.FindByID(id)
	if err != nil {
		return nil, err
	}
	dto := &dto.UserDTO{
		ID:    user.ID.String(),
		Name:  user.Name,
		Email: user.Email,
	}
	return dto, nil
}

func (u *UserUseCase) UpdateUser(idRequest string, dto *dto.UserResquest) error {
	id, err := idgen.ParseID(idRequest)
	if err != nil {
		return err
	}
	var user *domain.User
	userFound, err := u.UserRepository.FindByID(id)
	if err != nil {
		return err
	}
	user.ID = userFound.ID
	user.Name = dto.Name
	user.Email = dto.Email

	if dto.Password != "" {
		hash, err := bcrypt.GenerateFromPassword([]byte(dto.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		user.Password = string(hash)
	}
	err = u.UserRepository.Update(user)
	return nil
}

func (u *UserUseCase) DeleteUser(idRequest string) error {
	id, err := idgen.ParseID(idRequest)
	if err != nil {
		return err
	}
	return u.UserRepository.Delete(id)
}
