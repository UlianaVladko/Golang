package auth

import (
	"database/sql"
	"errors"

	"auth-sqlite/internal/user"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrUserNotFound  = errors.New("user not found")
	ErrWrongPassword = errors.New("wrong password")
)

func Authenticate(db *sql.DB, email, password string) error {
	hash, err := user.GetPasswordHashByEmail(db, email)

	if err == sql.ErrNoRows {
		return ErrUserNotFound
	}
	if err != nil {
		return err
	}

	if bcrypt.CompareHashAndPassword(
		[]byte(hash),
		[]byte(password),
	) != nil {
		return ErrWrongPassword
	}

	return nil
}
