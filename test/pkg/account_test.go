package pkg_test

import (
	"testing"

	"github.com/go-jwt-auth/pkg/account"

	"github.com/stretchr/testify/assert"
)

func TestCheckPassword(t *testing.T) {
	var err error
	err = account.CheckPassword("lowercase")
	assert.NotEqual(t, err, nil, err.Error())

	err = account.CheckPassword("UPPERCASE")
	assert.NotEqual(t, err, nil, err.Error())

	err = account.CheckPassword("Mixed123asd")
	assert.NotEqual(t, err, nil, err.Error())

	// case of success
	err = account.CheckPassword("Password$3")
	assert.Equal(t, err, nil)
}
