package api_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	handler "go-jwt-auth/internal/app/auth/api"

	"github.com/stretchr/testify/assert"
)

func TestPostLogin(t *testing.T) {
	pbytes, _ := json.Marshal(handler.Login{Username: "admin", Password: "password"})
	buff := bytes.NewBuffer(pbytes)

	resp, err := http.Post(testUrl+"/auth/login", "application/json", buff)
	assert.Equal(t, 200, resp.StatusCode, "Call PostLogin")
	assert.Equal(t, err, nil, "Post Login returns error")

	token := handler.Token{}
	json.NewDecoder(resp.Body).Decode(&token)

	assert.NotEqual(t, token.AccessToken, "", "Empty access token")
}
