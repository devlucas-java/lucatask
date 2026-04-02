package usecase

import (
	"github.com/devlucas-java/lucatask/internal/delivery/dto"
	"github.com/devlucas-java/lucatask/internal/infra/repository"
	"github.com/devlucas-java/lucatask/pkg/idgen"
	"golang.org/x/crypto/bcrypt"
)

type UserUseCase struct {
	UserRepository repository.UserRepository
}

func NewUserUseCase(u repository.UserRepository) UserUseCase {
	return UserUseCase{
		UserRepository: u,
	}
}

func (u *UserUseCase) UpdateUser(idRequest string, dto *dto.UserUpdateDTO) error {
	id, err := idgen.ParseID(idRequest)
	if err != nil {
		return err
	}
	userFound, err := u.UserRepository.FindByID(id)
	if err != nil || userFound == nil {
		return err
	}
	userFound.Name = dto.Name
	userFound.Email = dto.Email

	if dto.Password != "" {
		hash, err := bcrypt.GenerateFromPassword([]byte(dto.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		userFound.Password = string(hash)
	}
	err = u.UserRepository.Update(userFound)
	return nil
}

func (u *UserUseCase) DeleteUser(idRequest string) error {
	id, err := idgen.ParseID(idRequest)
	if err != nil {
		return err
	}
	user, err := u.UserRepository.FindByID(id)
	if err != nil {
		return err
	}
	return u.UserRepository.Delete(user.ID)
}

func (u *UserUseCase) GetUser(idRequest string) (*dto.UserDTO, error) {
	id, err := idgen.ParseID(idRequest)
	if err != nil {
		return nil, err
	}
	user, err := u.UserRepository.FindByID(id)
	if err != nil {
		return nil, err
	}
	userDTO := &dto.UserDTO{
		ID:    user.ID.String(),
		Name:  user.Name,
		Email: user.Email,
	}
	return userDTO, nil
}
