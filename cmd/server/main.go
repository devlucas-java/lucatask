package main

import (
	"net/http"

	"github.com/devlucas-java/lucatask/config"
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

	taskRouter := module.NewTaskModule(db)
	taskRouter.Register(r)

	err := http.ListenAndServe(":"+conf.Port, r)
	if err != nil {
		panic(err)
	}

}
