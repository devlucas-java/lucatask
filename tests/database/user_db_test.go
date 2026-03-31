package database

import (
	"testing"

	"github.com/devlucas-java/lucatask/internal/domain"
	"github.com/devlucas-java/lucatask/internal/infra/database"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestCreateUserAndFindUserByID(t *testing.T) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	assert.NoError(t, err)

	db.AutoMigrate(&domain.User{})

	repo := database.NewUserDB(db)

	user := domain.User{
		Name:     "name",
		Email:    "email",
		Password: "password",
	}
	err = repo.Create(&user)
	assert.NoError(t, err)

	userFound, err := repo.FindByID(user.ID)
	assert.NoError(t, err)
	assert.NotNil(t, userFound)
	assert.Equal(t, userFound.ID, user.ID)
	assert.Equal(t, userFound.Name, user.Name)
	assert.Equal(t, userFound.Email, user.Email)
	assert.Equal(t, userFound.Password, user.Password)
}

func TestFindUserByEmail(t *testing.T) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	assert.NoError(t, err)

	db.AutoMigrate(&domain.User{})

	repo := database.NewUserDB(db)

	user := domain.User{
		Name:     "name",
		Email:    "email",
		Password: "password",
	}
	err = repo.Create(&user)
	assert.NoError(t, err)

	userFound, err := repo.FindByEmail("email")
	assert.NoError(t, err)
	assert.NotNil(t, userFound)
	assert.Equal(t, userFound.ID, user.ID)
	assert.Equal(t, userFound.Name, user.Name)
	assert.Equal(t, userFound.Email, user.Email)
	assert.Equal(t, userFound.Password, user.Password)
}

func TestDeleteUser(t *testing.T) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	assert.NoError(t, err)

	db.AutoMigrate(&domain.User{})

	repo := database.NewUserDB(db)

	user := domain.User{
		Name:     "name",
		Email:    "email",
		Password: "password",
	}
	err = repo.Create(&user)
	assert.NoError(t, err)

	err = repo.Delete(user.ID)
	assert.NoError(t, err)

	userFound, err := repo.FindByID(user.ID)
	assert.Error(t, err)
	assert.Nil(t, userFound)
}

func TestUpdateUser(t *testing.T) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	assert.NoError(t, err)

	db.AutoMigrate(&domain.User{})

	repo := database.NewUserDB(db)

	user := domain.User{
		Name:     "name",
		Email:    "email",
		Password: "password",
	}
	err = repo.Create(&user)
	assert.NoError(t, err)

	user.Name = "new name"
	user.Email = "new email"
	user.Password = "new password"
	err = repo.Update(&user)
	assert.NoError(t, err)

	userFound, err := repo.FindByID(user.ID)
	assert.NoError(t, err)
	assert.NotNil(t, userFound)
	assert.Equal(t, userFound.ID, user.ID)
	assert.Equal(t, userFound.Name, user.Name)
	assert.Equal(t, userFound.Email, user.Email)
	assert.Equal(t, userFound.Password, user.Password)
}
