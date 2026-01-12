package main

import (
	"fmt"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request)  {
	w.Write([]byte("Succes request"))
}

func main() {
	http.HandleFunc("/", IndexHandler)
	fmt.Println("Server started...")
	http.ListenAndServe(":8080", nil)
}