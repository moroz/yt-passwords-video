package store

import (
	"github.com/jmoiron/sqlx"
	"github.com/moroz/yt-passwords-video/code/types"
)

type userStore struct {
	db *sqlx.DB
}

type UserStore interface {
	GetUserByEmail(email string) (*types.User, error)
	InsertUser(types.User) (*types.User, error)
}

func NewUserStore(db *sqlx.DB) UserStore {
	return userStore{db}
}

const userColumns = `id, email, password_hash, inserted_at, updated_at`
const getUserByEmailQuery = `select ` + userColumns + ` from users where email = $1`

func (us userStore) GetUserByEmail(email string) (*types.User, error) {
	var user types.User
	err := us.db.Get(&user, getUserByEmailQuery, email)
	return &user, err
}

const insertUserQuery = `insert into users (email, password_hash) values ($1, $2) returning ` + userColumns

func (us userStore) InsertUser(user types.User) (*types.User, error) {
	var newUser types.User
	err := us.db.Get(&newUser, user.Email, user.PasswordHash)
	return &newUser, err
}
