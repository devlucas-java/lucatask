package route

import (
	"github.com/devlucas-java/lucatask/internal/delivery/handle"
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

func (ur *UserRoute) Route(c chi.Router) {

	c.Route("/user", func(r chi.Router) {
		r.Delete("/me", ur.UserHandler.DeleteMe)
		r.Put("/me", ur.UserHandler.UpdateMe)
	})
}
