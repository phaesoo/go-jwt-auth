package auth

import (
	"log"
	"net/http"
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
}

func (a *App) Run(host string) {
	log.Fatal(http.ListenAndServe(host, a.Router))
}