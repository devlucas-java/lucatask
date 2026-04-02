package route

import (
	"github.com/devlucas-java/lucatask/internal/delivery/handle"
	"github.com/devlucas-java/lucatask/internal/delivery/middleware"
	"github.com/devlucas-java/lucatask/internal/infra/jwt"
	"github.com/go-chi/chi"
)

type UserRoute struct {
	UserHandler handle.UserHandle
}

func NewUserRoute(uh handle.UserHandle) *UserRoute {
	return &UserRoute{
		UserHandler: uh,
	}
}

func (ur *UserRoute) Route(c chi.Router, jwtService *jwt.JwtService) {

	c.Route("/user", func(r chi.Router) {
		r.Use(middleware.AuthMiddleware(jwtService))
		r.Delete("/me", ur.UserHandler.DeleteMe)
		r.Put("/me", ur.UserHandler.UpdateMe)
		r.Get("/me", ur.UserHandler.GetMe)
	})
}
