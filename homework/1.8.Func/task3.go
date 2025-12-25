package main

import "fmt"

func main() {
	// переменная, котор хранит функцию
	counter := func() func() int { // внешняя функция, котор возвращает другую функцию
		count := 0          //замыкание
		return func() int { // внутренняя функция
			count++
			return count
		}
	}() // немедленный вызов

	fmt.Println(counter()) // 1
	fmt.Println(counter()) // 2
	fmt.Println(counter()) // 3
}
