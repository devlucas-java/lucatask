package module

import (
	"github.com/devlucas-java/lucatask/internal/delivery/handle"
	"github.com/devlucas-java/lucatask/internal/delivery/route"
	"github.com/devlucas-java/lucatask/internal/infra/database"
	"github.com/devlucas-java/lucatask/internal/usecase"
	"gorm.io/gorm"
)

func NewUserModule(db *gorm.DB) *route.UserRoute {

	userRepo := database.NewUserDB(db)
	userUseCase := usecase.NewUserUseCase(userRepo)
	userHandler := handle.NewUserHandle(userUseCase)
	return route.NewUserRoute(userHandler)
}
