package api_test

import (
	"net/http/httptest"
	"os"
	"testing"

	"github.com/go-jwt-auth/internal/app/auth"
	"github.com/go-jwt-auth/internal/app/auth/db"

	"github.com/gorilla/handlers"
)

var (
	server  *httptest.Server
	testUrl string
)

func TestMain(m *testing.M) {
	Setup()
	exitVal := m.Run()
	Teardown()
	os.Exit(exitVal)
}

func Setup() {
	// prepare for test db
	db.InitTestDB()
	db.MigrateDB()
	db.SetInitialData()

	// prepare for httptest server
	logfile, err := os.OpenFile("/tmp/test.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0755)
	if err != nil {
		panic("InitTest failed")
	}
	app := auth.App{}
	app.Initialize()
	server = httptest.NewServer(handlers.CombinedLoggingHandler(logfile, app.Router))
	testUrl = server.URL + "/api"
}

func Teardown() {
	db.RemoveTestDB()
}
