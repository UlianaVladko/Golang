package main

import (
	"fmt"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	if r.Method == http.MethodPost {
		w.Write([]byte("Succes POST request"))
	} else if r.Method == http.MethodGet {
		w.Header().Add("Content-Type", "application/json")
		w.Write([]byte(`{"status": "Succes GET"}`))
	}

}

func main() {
	http.HandleFunc("/", IndexHandler)
	fmt.Println("Server started...")
	http.ListenAndServe(":8080", nil)
}
