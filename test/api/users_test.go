package api_test

import (
	"encoding/json"
	"log"
	"net/http"
	"testing"

	"go-jwt-auth/internal/app/auth/model"

	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	var err error
	resp, err := http.Get(testUrl + "/users")
	assert.Equal(t, resp.StatusCode, 200, "Error while http Get")
	assert.Equal(t, err, nil, "Error while http Get")

	users := []model.User{}
	err = json.NewDecoder(resp.Body).Decode(&users)
	assert.Equal(t, err, nil, "Error while decoding UserSummary")

	log.Println("username", users[0].Username, users[1].Username)
}
