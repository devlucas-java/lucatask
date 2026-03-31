package database

import (
	"github.com/devlucas-java/lucatask/internal/domain"
	"github.com/devlucas-java/lucatask/internal/infra/repository"
	"github.com/devlucas-java/lucatask/pkg/idgen"
	"gorm.io/gorm"
)

type UserDB struct {
	DB *gorm.DB
}

func NewUserDB(db *gorm.DB) repository.UserRepository {
	return &UserDB{DB: db}
}

func (u *UserDB) Create(user *domain.User) error {
	return u.DB.Create(user).Error
}

func (u *UserDB) Update(user *domain.User) error {
	return u.DB.Save(user).Error
}

func (u *UserDB) Delete(id idgen.ID) error {
	return u.DB.Where("id = ?", id).Delete(&domain.User{}).Error
}

func (u *UserDB) FindByID(id idgen.ID) (*domain.User, error) {
	var user domain.User
	err := u.DB.Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *UserDB) FindByEmail(email string) (*domain.User, error) {
	var user domain.User
	err := u.DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
