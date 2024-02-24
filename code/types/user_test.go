package types_test

import (
	"testing"

	"github.com/moroz/yt-passwords-video/code/types"
	"github.com/stretchr/testify/assert"
)

func TestRegisterUserParamsValidParams(t *testing.T) {
	params := types.RegisterUserParams{
		Email:                "user@example.com",
		Password:             "foobar2000",
		PasswordConfirmation: "foobar2000",
	}

	result := params.Validate()
	assert.True(t, result.Valid())
}

func TestRegisterUserParamsShortPassword(t *testing.T) {
	params := types.RegisterUserParams{
		Email:                "user@example.com",
		Password:             "foobar",
		PasswordConfirmation: "foobar",
	}

	result := params.Validate()
	assert.False(t, result.Valid())
}

func TestRegisterUserParamsEmailWithoutDot(t *testing.T) {
	params := types.RegisterUserParams{
		Email:                "user@invalid",
		Password:             "foobar2000",
		PasswordConfirmation: "foobar2000",
	}

	result := params.Validate()
	assert.False(t, result.Valid())
}

func TestRegisterUserParamsPasswordsDoNotMatch(t *testing.T) {
	params := types.RegisterUserParams{
		Email:                "user@example.com",
		Password:             "foobar2000",
		PasswordConfirmation: "foobar2001",
	}

	result := params.Validate()
	assert.False(t, result.Valid())
}
