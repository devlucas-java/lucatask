package usecase_test

import (
	"testing"

	"github.com/devlucas-java/lucatask/internal/delivery/dto"
	"github.com/devlucas-java/lucatask/internal/domain"
	"github.com/devlucas-java/lucatask/internal/infra/database"
	"github.com/devlucas-java/lucatask/internal/usecase"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestUpdateUserUseCase(t *testing.T) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	assert.NoError(t, err)

	db.AutoMigrate(&domain.User{})

	repo := database.NewUserDB(db)
	usecase := usecase.NewUserUseCase(repo)

	user := domain.User{
		Name:     "name",
		Email:    "email",
		Password: "password",
	}
	err = repo.Create(&user)
	assert.NoError(t, err)

	dto := dto.UserUpdateDTO{
		Name:     "new name",
		Email:    "new email",
		Password: "new password",
	}
	err = usecase.UpdateUser(user.ID.String(), &dto)
	assert.NoError(t, err)

	userFound, err := repo.FindByID(user.ID)
	assert.NoError(t, err)
	assert.NotNil(t, userFound)
	assert.Equal(t, userFound.ID, user.ID)
	assert.Equal(t, userFound.Name, dto.Name)
	assert.Equal(t, userFound.Email, dto.Email)
}
