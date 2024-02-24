package token_test

import (
	"testing"

	"github.com/moroz/yt-passwords-video/code/internal/token"
	"github.com/moroz/yt-passwords-video/code/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestIssueTokenForUser(t *testing.T) {
	user := types.User{
		ID: 42,
	}

	tokenString, err := token.IssueTokenForUser(&user)
	require.NoError(t, err)

	assert.NotEqual(t, "", tokenString)
}

func TestVerifyTokenForUser(t *testing.T) {
	user := types.User{
		ID: 42,
	}

	tokenString, err := token.IssueTokenForUser(&user)
	require.NoError(t, err)

	claims, err := token.DecodeToken(tokenString)
	require.NoError(t, err)
	assert.Equal(t, "42", claims.Subject)
}
