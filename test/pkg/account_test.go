package pkg_test

import (
	"log"
	"testing"

	"github.com/go-jwt-auth/pkg/account"

	"github.com/stretchr/testify/assert"
)

func TestCheckUsername(t *testing.T) {
	var err error
	err = account.CheckUsername("short")
	assert.NotEqual(t, err, nil, err.Error())

	err = account.CheckUsername("toooooooooooooooooooooolong")
	assert.NotEqual(t, err, nil, err.Error())

	err = account.CheckUsername("1startswithnum")
	assert.NotEqual(t, err, nil, err.Error())

	err = account.CheckUsername("Uppercasename1")
	assert.NotEqual(t, err, nil, err.Error())

	err = account.CheckUsername("includespecial$^$1")
	assert.NotEqual(t, err, nil, err.Error())

	// case of success
	err = account.CheckUsername("shouldhavesucceeded1")
	if err != nil {
		log.Println(err.Error())
	}
	assert.Equal(t, err, nil)
}


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
