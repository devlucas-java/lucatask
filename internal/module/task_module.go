package module

import (
	"github.com/devlucas-java/lucatask/internal/delivery/handle"
	"github.com/devlucas-java/lucatask/internal/delivery/route"
	"github.com/devlucas-java/lucatask/internal/infra/database"
	"github.com/devlucas-java/lucatask/internal/usecase"
	"gorm.io/gorm"
)

func NewTaskModule(db *gorm.DB) *route.TaskRoute {

	taskDB := database.NewTaskDB(db)
	taskUseCase := usecase.NewTaskUseCase(taskDB)
	taskHandle := handle.NewTaskHandle(*taskUseCase)
	return route.NewTaskRoute(taskHandle)
}
