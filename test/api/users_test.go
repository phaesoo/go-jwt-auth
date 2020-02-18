package api_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	handler "github.com/go-jwt-auth/internal/app/auth/api"
	"github.com/go-jwt-auth/internal/app/auth/model"

	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	var err error

	// convert login struct to buffer
	pbytes, _ := json.Marshal(handler.Login{Username: "admin", Password: "password"})
	buff := bytes.NewBuffer(pbytes)

	// request PostLogin with buffer to get access token
	resp_login, err := http.Post(testUrl+"/auth/login", "application/json", buff)
	assert.Equal(t, 200, resp_login.StatusCode, "Call PostLogin")
	assert.Equal(t, err, nil, "Post Login returns error")

	// decode access token
	token := handler.Token{}
	json.NewDecoder(resp_login.Body).Decode(&token)

	// request GET with access token
	req, err := http.NewRequest("GET", testUrl+"/users", nil)
	if err != nil {
		panic(err)
	}
	req.Header.Add("Authorization", token.AccessToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	// validate response
	assert.Equal(t, resp.StatusCode, 200, "Error while http Get")
	assert.Equal(t, err, nil, "Error while http Get")

	users := []model.User{}
	err = json.NewDecoder(resp.Body).Decode(&users)
	assert.Equal(t, err, nil, "Error while decoding UserSummary")
}

func TestPut(t *testing.T) {
	var err error
	var buff *bytes.Buffer

	// convert login struct to buffer
	if pbytes, err := json.Marshal(handler.Login{Username: "admin", Password: "password"}); err != nil {

	} else {
		buff = bytes.NewBuffer(pbytes)
	}

	// request PostLogin with buffer to get access token
	resp_login, err := http.Post(testUrl+"/auth/login", "application/json", buff)
	assert.Equal(t, 200, resp_login.StatusCode, "Call PostLogin")
	assert.Equal(t, err, nil, "Post Login returns error")

	// decode access token
	token := handler.Token{}
	json.NewDecoder(resp_login.Body).Decode(&token)

	// request PUT with access token
	pbytes, _ := json.Marshal(handler.ParseUpdate{Password: "password", NewPassword: "NewPassword!2"})
	buff = bytes.NewBuffer(pbytes)

	// request PostLogin with buffer to get access token
	req, err := http.NewRequest("PUT", testUrl+"/users/admin", buff)
	if err != nil {
		panic(err)
	}
	req.Header.Add("Authorization", token.AccessToken)

	client := &http.Client{}
	_, err = client.Do(req)
	if err != nil {
		panic(err)
	}
}