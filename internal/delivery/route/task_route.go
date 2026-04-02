package route

import (
	"github.com/devlucas-java/lucatask/internal/delivery/handle"
	"github.com/devlucas-java/lucatask/internal/delivery/middleware"
	"github.com/devlucas-java/lucatask/internal/infra/jwt"
	"github.com/go-chi/chi"
)

type TaskRoute struct {
	TaskHandle *handle.TaskHandle
}

func NewTaskRoute(taskHandle *handle.TaskHandle) *TaskRoute {
	return &TaskRoute{
		TaskHandle: taskHandle,
	}
}

func (tr *TaskRoute) Register(c chi.Router, jwtService *jwt.JwtService) {

	c.Route("/tasks", func(r chi.Router) {
		r.Use(middleware.AuthMiddleware(jwtService))
		r.Post("/", tr.TaskHandle.CreateTask)
		r.Get("/", tr.TaskHandle.GetAllTasks)

		r.Get("/{id}", tr.TaskHandle.GetTask)
		r.Put("/{id}", tr.TaskHandle.UpdateTask)
		r.Patch("/{id}/complete", tr.TaskHandle.CompletedTask)
		r.Delete("/{id}", tr.TaskHandle.DeleteTask)
	})
}
