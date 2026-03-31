package usecase_test

import (
	"testing"

	"github.com/devlucas-java/lucatask/internal/delivery/dto"
	"github.com/devlucas-java/lucatask/internal/domain"
	"github.com/devlucas-java/lucatask/internal/infra/database"
	"github.com/devlucas-java/lucatask/internal/usecase"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestLogin(t *testing.T) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	assert.NoError(t, err)

	db.AutoMigrate(&domain.User{})

	repo := database.NewUserDB(db)
	usecase := usecase.NewAuthUseCase(repo)

	hash, err := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
	assert.NoError(t, err)

	user := domain.User{
		Name:     "name",
		Email:    "login",
		Password: string(hash),
	}
	_ = repo.Create(&user)

	dto, err := usecase.Login("login", "password")
	assert.NoError(t, err)
	assert.NotNil(t, dto)
	assert.Equal(t, dto.User.Email, user.Email)
	assert.Equal(t, dto.User.Name, user.Name)
	assert.NotEmpty(t, dto.Token)
}

func TestLoginInvalidCredentials(t *testing.T) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	assert.NoError(t, err)

	db.AutoMigrate(&domain.User{})

	repo := database.NewUserDB(db)
	usecase := usecase.NewAuthUseCase(repo)

	_, err = usecase.Login("nonexistent", "password")
	assert.Error(t, err)
}

func TestRegisterSuccessfully(t *testing.T) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	assert.NoError(t, err)

	db.AutoMigrate(&domain.User{})

	repo := database.NewUserDB(db)
	usecase := usecase.NewAuthUseCase(repo)
	dtoRegister := dto.RegisterDTO{
		Name:     "name",
		Email:    "email",
		Password: "password",
	}

	dto, err := usecase.Register(&dtoRegister)
	assert.NoError(t, err)
	assert.NotNil(t, dto)
	assert.Equal(t, dto.User.Email, "email")
	assert.Equal(t, dto.User.Name, "name")
	assert.NotEmpty(t, dto.Token)
}

func TestRegisterDuplicateEmail(t *testing.T) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	assert.NoError(t, err)

	db.AutoMigrate(&domain.User{})

	repo := database.NewUserDB(db)
	usecase := usecase.NewAuthUseCase(repo)

	dtoRegister := dto.RegisterDTO{
		Name:     "name",
		Email:    "email",
		Password: "password",
	}

	_, err = usecase.Register(&dtoRegister)
	assert.NoError(t, err)

	_, err = usecase.Register(&dtoRegister)
	assert.Error(t, err)
}
