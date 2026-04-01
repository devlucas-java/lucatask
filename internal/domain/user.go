package domain

import (
	"time"

	"github.com/devlucas-java/lucatask/pkg/idgen"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        idgen.ID `gorm:"primary_key"`
	Name      string
	Email     string `gorm:"unique"`
	Role      string
	Password  string
	CreatedAt time.Time
}

func NewUser(name, email, password string) *User {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil
	}
	return &User{
		ID:        idgen.NewID(),
		Name:      name,
		Email:     email,
		Role:      "USER",
		Password:  string(hash),
		CreatedAt: time.Now(),
	}
}

func (u *User) ValidatePassword(password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)) == nil
}
