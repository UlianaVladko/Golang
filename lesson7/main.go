package main

import "fmt"

// func Cube(x int) (res int) {
// 	res = x * x * x
// 	return
// }

// func Cube(x, y int) int { //или (x int,  y int)
// 	res := x * y
// 	return res
// }

func Sum(x ...int) (res int){ //...int под капотом слайс []int, вспоминай домашку 1.3
	for _, value := range x{ //index анонимная переменная (не нужна тут)
		res += value
	}
	return
}

func main() {
	result := Sum(1, 2, 3, 4, 6)
	fmt.Println(result)

	// result := Cube(2, 3)
	// fmt.Println(result)
}
