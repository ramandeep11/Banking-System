package util

import (
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func TestPassword(t *testing.T) {
	password := RandomString(6);

	hased, err := HashPassword(password)

	require.NoError(t, err)
	require.NotEmpty(t, hased)

	randomStr := RandomString(10)
	err = CheckPassword(randomStr,hased)
	require.EqualError(t, err , bcrypt.ErrMismatchedHashAndPassword.Error())

	hashed2 , err := HashPassword(password)

	require.NotEqual(t,hased,hashed2)
	
}
