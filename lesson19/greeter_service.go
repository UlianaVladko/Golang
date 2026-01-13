package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type User struct {
	Id   string
	Name string
}

func greeterHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/greeter/"):]
	resp, err := http.Get("http://127.0.0.1:8080/user/" + id)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	var user User
	err = json.NewDecoder(resp.Body).Decode(&user)
	if err != nil {
		fmt.Println(err)
		return
	}
	w.Write([]byte(fmt.Sprintf("Hello, %s!", user.Name)))
}

func main() {
	http.HandleFunc("/greeter/", greeterHandler)

	log.Println("Greeter service started...")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatal("8081 - error")
	}
}
