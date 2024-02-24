package types

import "time"

type User struct {
	ID           int       `db:"id" json:"id"`
	Email        string    `db:"email" json:"email"`
	PasswordHash string    `db:"password_hash" json:"passwordHash"`
	InsertedAt   time.Time `db:"inserted_at" json:"insertedAt"`
	UpdatedAt    time.Time `db:"updated_at" json:"updatedAt"`
}
