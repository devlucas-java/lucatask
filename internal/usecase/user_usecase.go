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

func (u *UserUseCase) UpdateUser(idRequest string, dto *dto.UserUpdateDTO) error {
	id, err := idgen.ParseID(idRequest)
	if err != nil {
		return err
	}
	var user *domain.User
	userFound, err := u.UserRepository.FindByID(id)
	if err != nil || userFound == nil {
		return err
	}
	user.ID = id
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
