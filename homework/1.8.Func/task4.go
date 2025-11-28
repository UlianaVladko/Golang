package main

import (
	"fmt"
)

// дженерики
type Generics interface {
	// ~int | ~float64 // если с производными типы MyInt
	int | int8 | int16 | int32 | int64 | float32 | float64
}

func sumAll[T Generics](setNumbers ...T) T {
	var sum T
	for _, value := range setNumbers {
		sum += value
	}
	return sum
}

// func sumAll(setNumbers ...int) int {
// 	sum := 0
// 	for _, value := range setNumbers {
// 		sum += value
// 	}
// 	return sum
// }

func main() {
	fmt.Println(sumAll(1, 2, 3))       // 6
	fmt.Println(sumAll(10, -2, 4, 7))  // 19
	fmt.Println(sumAll(1.1, 2.2, 3.3)) // 6.6
}
