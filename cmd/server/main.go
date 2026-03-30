package main

import (
	"net/http"

	"github.com/devlucas-java/lucatask/config"
	"github.com/go-chi/chi"
)

func main() {

	config := config.InitConfig()

	r := chi.NewRouter()

	http.ListenAndServe(config.Port, r)

}
