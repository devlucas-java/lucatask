package route

import (
	"github.com/devlucas-java/lucatask/internal/delivery/handle"
	"github.com/go-chi/chi"
)

type AuthRoute struct {
	AuthHandler handle.AuthHandle
}

func NewAuthRoute(authHandler handle.AuthHandle) *AuthRoute {
	return &AuthRoute{
		AuthHandler: authHandler,
	}
}
func (ar *AuthRoute) Register(c chi.Router) {
	c.Route("/auth", func(r chi.Router) {
		r.Post("/login", ar.AuthHandler.Login)
		r.Post("/register", ar.AuthHandler.Register)
	})
}
