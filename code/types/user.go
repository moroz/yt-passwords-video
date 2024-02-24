package types

import "time"

type User struct {
	ID           int       `db:"id" json:"id"`
	Email        string    `db:"email" json:"email"`
	PasswordHash string    `db:"password_hash" json:"passwordHash"`
	InsertedAt   time.Time `db:"inserted_at" json:"insertedAt"`
	UpdatedAt    time.Time `db:"updated_at" json:"updatedAt"`
}

type RegisterUserParams struct {
	Email                string `json:"email"`
	Password             string `json:"password"`
	PasswordConfirmation string `json:"passwordConfirmation"`
}

func (p RegisterUserParams) Validate() ValidationResult {
	var result ValidationResult
	if p.Email == "" {
		result.Add("email", "can't be blank")
	}
	if p.Password == "" {
		result.Add("password", "can't be blank")
	}
	if len(p.Password) < 8 {
		result.Add("password", "must be at least 8 characters long")
	}
	if p.PasswordConfirmation != p.Password {
		result.Add("password", "passwords do not match")
	}
	return result
}
