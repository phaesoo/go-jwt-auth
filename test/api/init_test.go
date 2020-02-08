package api_test

import (
	"log"
	"net/http/httptest"
	"os"
	"testing"

	"go-jwt-auth/internal/app/auth"
	"go-jwt-auth/internal/app/auth/db"

	"github.com/gorilla/handlers"
)

var (
	server  *httptest.Server
	testUrl string
)

func TestMain(m *testing.M) {
	log.Println("Start")
	InitTest()
	exitVal := m.Run()
	log.Println("End")

	os.Exit(exitVal)
}

func InitTest() {
	// prepare for test db
	db.InitTestDB()
	db.MigrateDB()
	db.SetInitialData()
	defer db.RemoveTestDB()

	// prepare for httptest server
	logfile, err := os.OpenFile("/tmp/test.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0755)
	if err != nil {
		panic("InitTest failed")
	}
	app := auth.App{}
	app.Initialize()
	server = httptest.NewServer(handlers.CombinedLoggingHandler(logfile, app.Router))
	testUrl = server.URL
}
