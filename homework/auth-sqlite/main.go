package main

import (
	"fmt"
	"log"

	"auth-sqlite/internal/auth"
	"auth-sqlite/internal/db"
	"auth-sqlite/internal/user"
)

func main() {
	database, err := db.NewSQLite("users.db")
	if err != nil {
		log.Fatal(err)
	}
	defer database.Close()

	if err := user.InitTable(database); err != nil {
		log.Fatal(err)
	}

	if err := user.Seed(database, 5); err != nil {
		log.Fatal(err)
	}

	err = auth.Authenticate(database, "user2@example.com", "password123")
	fmt.Println(authResult(err))
}

func authResult(err error) string {
	switch err {
	case nil:
		return "пароль верный"
	case auth.ErrUserNotFound:
		return "пользователь не найден"
	case auth.ErrWrongPassword:
		return "пароль неверный"
	default:
		return "внутренняя ошибка"
	}
}
