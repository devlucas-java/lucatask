package database

import (
	"testing"

	"github.com/devlucas-java/lucatask/internal/domain"
	"github.com/devlucas-java/lucatask/internal/infra/database"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestCreateTask(t *testing.T) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	assert.NoError(t, err)

	db.AutoMigrate(&domain.Task{})

	repo := database.NewTaskDB(db)

	task := domain.Task{
		Name:        "title",
		Description: "body",
	}
	err = repo.Create(&task)
	assert.NoError(t, err)
}

func TestFindTaskByID(t *testing.T) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	assert.NoError(t, err)

	db.AutoMigrate(&domain.Task{})

	repo := database.NewTaskDB(db)

	task := domain.Task{
		Name:        "title",
		Description: "body",
	}
	err = repo.Create(&task)
	assert.NoError(t, err)

	taskFound, err := repo.FindByID(task.ID)
	assert.NoError(t, err)
	assert.NotNil(t, taskFound)
	assert.Equal(t, taskFound.ID, task.ID)
	assert.Equal(t, taskFound.Name, task.Name)
	assert.Equal(t, taskFound.Description, task.Description)
}

func TestFindById(t *testing.T) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	assert.NoError(t, err)

	db.AutoMigrate(&domain.Task{})

	repo := database.NewTaskDB(db)

	task := domain.Task{
		Name:        "title",
		Description: "body",
	}
	err = repo.Create(&task)
	assert.NoError(t, err)

	taskFound, err := repo.FindByID(task.ID)
	assert.NoError(t, err)
	assert.NotNil(t, taskFound)
	assert.Equal(t, taskFound.ID, task.ID)
	assert.Equal(t, taskFound.Name, task.Name)
	assert.Equal(t, taskFound.Description, task.Description)
}

func TestDeleteTask(t *testing.T) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	assert.NoError(t, err)

	db.AutoMigrate(&domain.Task{})

	repo := database.NewTaskDB(db)

	task := domain.Task{
		Name:        "title",
		Description: "body",
	}
	err = repo.Create(&task)
	assert.NoError(t, err)

	err = repo.Delete(task.ID)
	assert.NoError(t, err)

	taskFound, err := repo.FindByID(task.ID)
	assert.Error(t, err)
	assert.Nil(t, taskFound)
}
