package api_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	handler "github.com/go-jwt-auth/internal/app/auth/api"

	"github.com/stretchr/testify/assert"
)

func TestPostLogin(t *testing.T) {
	// prepare for buffer with login info.
	pbytes, _ := json.Marshal(handler.Login{Username: "admin", Password: "password"})
	buff := bytes.NewBuffer(pbytes)

	// reqeust PostLogin
	resp, err := http.Post(testUrl+"/auth/login", "application/json", buff)
	assert.Equal(t, 200, resp.StatusCode, "Call PostLogin")
	assert.Equal(t, err, nil, "Post Login returns error")

	token := handler.Token{}
	json.NewDecoder(resp.Body).Decode(&token)

	// validate access token
	assert.NotEqual(t, token.AccessToken, "", "Empty access token")
}
