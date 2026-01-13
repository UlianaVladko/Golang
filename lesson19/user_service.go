/*
На яп Го требуется реализовать веб-сервис приветствия пользователя в микросервисной архитектуре:
1) Получает на вход GET запрос /user/{id}
2) Получает (от 1 сервиса) ID, возвращает Привет, {name(id)}
*/

package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type User struct {
	Id   string
	Name string
}

var Ivan User = User{Id: "1", Name: "Ivan"}

func userHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/user/"):]
	if id == Ivan.Id {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(Ivan)
	} else {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(User{Id: "0", Name: "Guest"})
	}
}

func main() {

	http.HandleFunc("/user/", userHandler)

	log.Println("User service started...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("8080 - error")
	}

}
