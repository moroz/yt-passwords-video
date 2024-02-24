package service

import (
	"errors"

	"github.com/alexedwards/argon2id"
	"github.com/jmoiron/sqlx"
	"github.com/moroz/yt-passwords-video/code/store"
	"github.com/moroz/yt-passwords-video/code/types"
)

type userService struct {
	userStore store.UserStore
}

func NewUserService(db *sqlx.DB) userService {
	return userService{userStore: store.NewUserStore(db)}
}

var argonProductionParams = &argon2id.Params{Memory: 19, Iterations: 2, Parallelism: 1, KeyLength: 16, SaltLength: 16}

func (us userService) RegisterUser(params types.RegisterUserParams) (*types.User, error, types.ValidationResult) {
	validationResult := params.Validate()
	if !validationResult.Valid() {
		return nil, errors.New("invalid params"), validationResult
	}

	passwordHash, err := argon2id.CreateHash(params.Password, argonProductionParams)
	if err != nil {
		return nil, err, validationResult
	}

	newUser := types.User{
		Email:        params.Email,
		PasswordHash: passwordHash,
	}
	user, err := us.userStore.InsertUser(newUser)
	return user, err, validationResult
}

func noUserVerify() {
	argon2id.CreateHash("dummy", argonProductionParams)
}

func (us userService) AuthenticateUserByEmailPassword(email, password string) (*types.User, error) {
	user, err := us.userStore.GetUserByEmail(email)
	if err != nil || user == nil {
		noUserVerify()
		return nil, err
	}

	if match, err := argon2id.ComparePasswordAndHash(password, user.PasswordHash); !match {
		return nil, err
	}

	return user, nil
}
