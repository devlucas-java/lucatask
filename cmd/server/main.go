// @title LucaTask API
// @version 1.0
// @description Task management API with JWT authentication
// @termsOfService http://swagger.io/terms/

// @contact.name Lucas Dev
// @contact.email lucas@example.com

// @license.name MIT
// @host localhost:8080
// @BasePath /api

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
package main

import (
	"net/http"

	"github.com/devlucas-java/lucatask/config"
	_ "github.com/devlucas-java/lucatask/docs"
	"github.com/devlucas-java/lucatask/internal/infra/jwt"
	"github.com/devlucas-java/lucatask/internal/module"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	httpSwagger "github.com/swaggo/http-swagger"
)

func main() {

	conf := config.InitConfig()
	db := config.InitDatabase()

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Get("/swagger/*", httpSwagger.WrapHandler)

	jwtService := jwt.NewJwtService(conf.JWT_Secret)

	taskRouter := module.NewTaskModule(db)
	taskRouter.Register(r, jwtService)

	userRouter := module.NewUserModule(db)
	userRouter.Route(r, jwtService)

	authRouter := module.NewAuthModule(db, jwtService)
	authRouter.Register(r)

	err := http.ListenAndServe(":"+conf.Port, r)
	if err != nil {
		panic(err)
	}

}
