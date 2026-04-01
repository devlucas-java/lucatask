package module

import (
	"github.com/devlucas-java/lucatask/config"
	"github.com/devlucas-java/lucatask/internal/delivery/handle"
	"github.com/devlucas-java/lucatask/internal/delivery/route"
	"github.com/devlucas-java/lucatask/internal/infra/database"
	"github.com/devlucas-java/lucatask/internal/infra/jwt"
	"github.com/devlucas-java/lucatask/internal/usecase"
	"gorm.io/gorm"
)

func NewAuthModule(db *gorm.DB) *route.AuthRoute {

	cfg := config.GetConfig()

	userRepo := database.NewUserDB(db)
	jwtService := jwt.NewJwtService(cfg.JWT_Secret)
	userUseCase := usecase.NewAuthUseCase(userRepo, jwtService)
	userHandler := handle.NewAuthHandle(userUseCase)

	return route.NewAuthRoute(userHandler)
}
