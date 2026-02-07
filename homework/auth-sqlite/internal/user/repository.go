package user

import "database/sql"

func InitTable(db *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		email TEXT UNIQUE NOT NULL,
		password TEXT NOT NULL
	);`
	_, err := db.Exec(query)
	return err
}

func GetPasswordHashByEmail(db *sql.DB, email string) (string, error) {
	var hash string
	err := db.QueryRow(
		"SELECT password FROM users WHERE email = ?",
		email,
	).Scan(&hash)

	return hash, err
}
