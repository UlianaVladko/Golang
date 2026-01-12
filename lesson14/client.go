package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"time"
)

/*
// функция для обработки ошибки (но не надо всегда так делать!!!)
// каждая ошибка должна обрабатываться ПО-СВОЕМУ!!!!
func getErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}*/

func main() {
	// response, err := http.Get("http://127.0.0.1:8080/")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	client := http.Client{
		Timeout: 1 * time.Second,
	}
	data := []byte(`{"key": "value"}`)

	request, err := http.NewRequest(http.MethodPost, "http://127.0.0.1:8080/", bytes.NewBuffer(data))

	// getErr(err)
	if err != nil {
		fmt.Println(err)
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Cookie", "123456789")

	response, err := client.Do(request)
	// getErr(err)
	if err != nil {
		fmt.Println(err)
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	// getErr(err)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(body))
}
