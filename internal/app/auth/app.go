package auth

import (
	"log"
	"net/http"

	handler "go-jwt-auth/internal/app/auth/api"

	"github.com/gorilla/mux"
)

type App struct {
	Router *mux.Router
}

func (a *App) Initialize() {
	a.Router = mux.NewRouter()
	a.setRouters()
}

func (a *App) setRouters() {
	api := a.Router.PathPrefix("/api").Subrouter()
	api.HandleFunc("/auth/login", handler.PostLogin).Methods(http.MethodPost)
}

func (a *App) Run(host string) {
	log.Fatal(http.ListenAndServe(host, a.Router))
}
