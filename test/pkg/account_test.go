package pkg_test

import (
	"log"
	"testing"

	"github.com/go-jwt-auth/pkg/account"

	"github.com/stretchr/testify/assert"
)

func TestIsValidUsername(t *testing.T) {
	var ok bool
	var err error
	ok, err = account.IsValidUsername("short")
	assert.Equal(t, ok, false, err.Error())

	ok, err = account.IsValidUsername("toooooooooooooooooooooolong")
	assert.Equal(t, ok, false, err.Error())

	ok, err = account.IsValidUsername("1startswithnum")
	assert.Equal(t, ok, false, err.Error())

	ok, err = account.IsValidUsername("Uppercasename1")
	assert.Equal(t, ok, false, err.Error())

	ok, err = account.IsValidUsername("includespecial$^$1")
	assert.Equal(t, ok, false, err.Error())

	// case of success
	ok, err = account.IsValidUsername("shouldhavesucceeded1")
	if err != nil {
		log.Println(err.Error())
	}
	assert.Equal(t, ok, true)
}

func TestIsValidPassword(t *testing.T) {
	var ok bool
	var err error
	ok, err = account.IsValidPassword("lowercase")
	assert.Equal(t, ok, false, err.Error())

	ok, err = account.IsValidPassword("UPPERCASE")
	assert.Equal(t, ok, false, err.Error())

	ok, err = account.IsValidPassword("Mixed123asd")
	assert.Equal(t, ok, false, err.Error())

	// case of success
	ok, err = account.IsValidPassword("Password$3")
	assert.Equal(t, ok, true)
}

func TestIsValidEmail(t *testing.T) {
	var ok bool
	var err error
	ok, err = account.IsValidEmail("test")
	assert.Equal(t, ok, false)
	if err != nil {
		log.Println(err.Error())
	}

	ok, err = account.IsValidEmail("test@test")
	assert.Equal(t, ok, false)

	ok, err = account.IsValidEmail("@test.com")
	assert.Equal(t, ok, false)

	// case of success
	ok, err = account.IsValidEmail("test@test.com")
	assert.Equal(t, ok, true)
}
