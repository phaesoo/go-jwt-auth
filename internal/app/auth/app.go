package auth

import (
	"log"
	"net/http"

	handler "github.com/go-jwt-auth/internal/app/auth/api"

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
	api.HandleFunc("/auth/me", handler.GetMe).Methods(http.MethodGet)
	api.HandleFunc("/users", handler.Get).Methods(http.MethodGet)
	api.HandleFunc("/users/{username}", handler.GetUser).Methods(http.MethodGet)
	api.HandleFunc("/users/{username}", handler.PutUser).Methods(http.MethodPut)
	api.HandleFunc("/users/{username}", handler.DeleteUser).Methods(http.MethodDelete)
}

func (a *App) Run(host string) {
	log.Fatal(http.ListenAndServe(host, a.Router))
}
