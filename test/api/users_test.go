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

func TestGetAll(t *testing.T) {
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

func TestGet(t *testing.T) {
	// login information
	username := "admin"
	password := "password"

	targetUsername := "test"

	var err error

	// convert login struct to buffer
	pbytes, _ := json.Marshal(handler.Login{Username: username, Password: password})
	buff := bytes.NewBuffer(pbytes)

	// request PostLogin with buffer to get access token
	resp_login, err := http.Post(testUrl+"/auth/login", "application/json", buff)
	assert.Equal(t, 200, resp_login.StatusCode, "Call PostLogin")
	assert.Equal(t, err, nil, "Post Login returns error")

	// decode access token
	token := handler.Token{}
	json.NewDecoder(resp_login.Body).Decode(&token)

	// request GET with access token
	req, err := http.NewRequest("GET", testUrl+"/users/"+targetUsername, nil)
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

	user := model.User{}
	err = json.NewDecoder(resp.Body).Decode(&user)
	assert.Equal(t, user.Username, targetUsername, "Error while decoding UserSummary")
}

func TestPut(t *testing.T) {
	// login information
	username := "admin"
	password := "password"

	targetUsername := "test"
	targetPassword := "password"
	targetNewPassword := "NewPassword!2"

	var err error
	var buff *bytes.Buffer

	// convert login struct to buffer
	if pbytes, err := json.Marshal(handler.Login{Username: username, Password: password}); err != nil {
		t.Error(err.Error())
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
	if pbytes, err := json.Marshal(handler.ParseUpdate{Password: targetPassword, NewPassword: targetNewPassword}); err != nil {
		t.Error(err.Error())
	} else {
		buff = bytes.NewBuffer(pbytes)
	}

	req, err := http.NewRequest("PUT", testUrl+"/users/"+targetUsername, buff)
	if err != nil {
		t.Error(err.Error())
	}
	req.Header.Add("Authorization", token.AccessToken)

	client := &http.Client{}
	if resp, err := client.Do(req); err != nil {
		t.Error(err.Error())
	} else {
		assert.Equal(t, 200, resp.StatusCode, "PutUser")
	}

	// try login with new password
	if pbytes, err := json.Marshal(handler.Login{Username: targetUsername, Password: targetNewPassword}); err != nil {
		t.Error(err.Error())
	} else {
		buff = bytes.NewBuffer(pbytes)
	}
	if resp_login, err := http.Post(testUrl+"/auth/login", "application/json", buff); err != nil {
		t.Error(err.Error())
	} else {
		assert.Equal(t, 200, resp_login.StatusCode, "Call PostLogin")
	}
}

func TestDelete(t *testing.T) {
	// login information
	username := "admin"
	password := "password"

	targetUsername := "tobedeleted"

	var err error
	var buff *bytes.Buffer

	// convert login struct to buffer
	if pbytes, err := json.Marshal(handler.Login{Username: username, Password: password}); err != nil {
		t.Error(err.Error())
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

	req, err := http.NewRequest("DELETE", testUrl+"/users/"+targetUsername, nil)
	if err != nil {
		t.Error(err.Error())
	}
	req.Header.Add("Authorization", token.AccessToken)

	client := &http.Client{}
	if resp, err := client.Do(req); err != nil {
		t.Error(err.Error())
	} else {
		assert.Equal(t, 200, resp.StatusCode, "DeleteUser")
	}

	// check target username has been removed from db
	if req, err := http.NewRequest("GET", testUrl+"/users/"+targetUsername, nil); err != nil {
		t.Error(err.Error())
	} else {
		req.Header.Add("Authorization", token.AccessToken)
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			t.Error(err.Error())
		}

		// validate response
		assert.Equal(t, resp.StatusCode, http.StatusBadRequest, "Error while http Get")
	}
}
