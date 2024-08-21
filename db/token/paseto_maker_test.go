package token

import (
	"simplebank/db/util"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestPasetoMaker(t *testing.T) {
	maker, err := NewpasetoMaker(util.RandomString(32))
	require.NoError(t, err)

	username := util.RandomOwner()

	duration := time.Minute

	// issuedAt := time.Now()
	// expiredAt := issuedAt.Add(duration)

	token, err := maker.CreateToken(username, duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	payload, err := maker.VerifyToken(token)
	require.NoError(t, err)
	require.NotEmpty(t, payload)
	// exp, _ := payload.parseNumericDate()
	// require.WithinDuration(t, expiredAt)

}



func TestExpiredPasetoMaker(t *testing.T) {
	maker, err := NewpasetoMaker(util.RandomString(32))
	require.NoError(t, err)

	username := util.RandomOwner()

	duration := -time.Minute

	// issuedAt := time.Now()
	// expiredAt := issuedAt.Add(duration)

	token, err := maker.CreateToken(username, duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	payload, err := maker.VerifyToken(token)
	require.NoError(t, err)
	require.NotEmpty(t, payload)
	// exp, _ := payload.parseNumericDate()
	// require.WithinDuration(t, expiredAt)

}