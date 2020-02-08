package api_test

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"testing"

	handler "go-jwt-auth/internal/app/auth/api"

	"github.com/stretchr/testify/assert"
)

func TestPostLogin(t *testing.T) {
	log.Println("TestPostLogin")

	pbytes, _ := json.Marshal(handler.Login{Username: "admin", Password: "password"})
	buff := bytes.NewBuffer(pbytes)

	log.Println(testUrl + "/api/auth/login")
	resp, err := http.Post(testUrl+"/api/auth/login", "application/json", buff)
	assert.Equal(t, 200, resp.StatusCode, "Call PostLogin")
	log.Println(resp.StatusCode)
	if err != nil {
		log.Println("Post error", err)
	}

	token := handler.Token{}
	json.NewDecoder(resp.Body).Decode(&token)

	log.Println("token", token.AccessToken)
}
