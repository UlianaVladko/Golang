package user

import (
	"database/sql"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func Seed(db *sql.DB, count int) error {
	for i := 1; i <= count; i++ {
		email := fmt.Sprintf("user%d@example.com", i)
		password := "password123"

		hash, err := bcrypt.GenerateFromPassword(
			[]byte(password),
			bcrypt.DefaultCost,
		)
		if err != nil {
			return err
		}

		_, err = db.Exec(
			"INSERT OR IGNORE INTO users(email, password) VALUES (?, ?)",
			email,
			string(hash),
		)
		if err != nil {
			return err
		}
	}
	return nil
}
