package main

import (
	"net/http"

	"github.com/devlucas-java/lucatask/config"
	"github.com/devlucas-java/lucatask/internal/infra/jwt"
	"github.com/devlucas-java/lucatask/internal/module"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {

	conf := config.InitConfig()
	db := config.InitDatabase()

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	jwtService := jwt.NewJwtService(conf.JWT_Secret)

	taskRouter := module.NewTaskModule(db)
	taskRouter.Register(r, jwtService)

	userRouter := module.NewUserModule(db)
	userRouter.Route(r)

	authRouter := module.NewAuthModule(db)
	authRouter.Register(r)

	err := http.ListenAndServe(":"+conf.Port, r)
	if err != nil {
		panic(err)
	}

}
