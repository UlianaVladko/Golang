package main

import "fmt"

//Анонимная функция
func applyOperation(a, b int, op func(int, int) int) int {
	return op(a, b)
}

//Именованная функция (Создайте несколько функций-операций)
func Add(a, b int) int {
	return a + b
}

func Diff(a, b int) int {
	return a - b
}

func Mult(a, b int) int {
	return a * b
}

// в main подставляю свои функции-операции в op
func main() {
	// res := applyOperation(10, 5, func(a, b int) int { // просто попробовать
	// 	return a + b
	// })
	// fmt.Printf("a + b = %d\n", res)

	res := applyOperation(10, 5, Add)
	fmt.Printf("a + b = %d\n", res)

	res = applyOperation(10, 5, Diff)
	fmt.Printf("a - b = %d\n", res)

	res = applyOperation(10, 5, Mult)
	fmt.Printf("a * b = %d", res)

}
