package module

import (
	"github.com/devlucas-java/lucatask/internal/delivery/handle"
	"github.com/devlucas-java/lucatask/internal/delivery/route"
	"github.com/devlucas-java/lucatask/internal/infra/database"
	"github.com/devlucas-java/lucatask/internal/infra/jwt"
	"github.com/devlucas-java/lucatask/internal/usecase"
	"gorm.io/gorm"
)

func NewAuthModule(db *gorm.DB, jwtService *jwt.JwtService) *route.AuthRoute {

	userRepo := database.NewUserDB(db)
	userUseCase := usecase.NewAuthUseCase(userRepo, jwtService)
	userHandler := handle.NewAuthHandle(userUseCase)

	return route.NewAuthRoute(userHandler)
}
