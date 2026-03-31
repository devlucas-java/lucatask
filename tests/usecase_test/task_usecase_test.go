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

func TestUpdateUser(t *testing.T) {
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

	usecase.UpdateUser(user.ID.String(), &dto)

	userFound, err := repo.FindByID(user.ID)
	assert.NoError(t, err)
	assert.NotNil(t, userFound)
	assert.Equal(t, userFound.ID, user.ID)
	assert.NotEqual(t, userFound.Name, dto.Name)
	assert.Equal(t, "new email", userFound.Email)
}

func TestDeleteTask(t *testing.T) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	assert.NoError(t, err)

	db.AutoMigrate(&domain.Task{})

	repo := database.NewTaskDB(db)
	usecase := usecase.NewTaskUseCase(repo)

	task := domain.Task{
		Name:        "title",
		Description: "body",
	}
	err = repo.Create(&task)
	assert.NoError(t, err)

	usecase.DeleteTask(task.ID.String())

	taskFound, err := repo.FindByID(task.ID)
	assert.Error(t, err)
	assert.Nil(t, taskFound)
}

func TestCompletedTask(t *testing.T) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	assert.NoError(t, err)

	db.AutoMigrate(&domain.Task{})

	repo := database.NewTaskDB(db)
	usecase := usecase.NewTaskUseCase(repo)

	task := domain.Task{
		Name:        "title",
		Description: "body",
	}
	err = repo.Create(&task)
	assert.NoError(t, err)

	dto := dto.TaskCompletedDTO{
		Completed: true,
	}
	usecase.CompletedTask(task.ID.String(), &dto)

	taskFound, err := repo.FindByID(task.ID)
	assert.NoError(t, err)
	assert.NotNil(t, taskFound)
	assert.Equal(t, taskFound.ID, task.ID)
	assert.Equal(t, taskFound.Completed, true)
}
