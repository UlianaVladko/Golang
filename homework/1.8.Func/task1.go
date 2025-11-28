package main

import (
	"errors"
	"fmt"
	// "math/rand"
	// "time"
)

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	result := a / b
	return result, nil
}

// Понимаю, что глобальные переменные - плохая практтика в многопоточном окружении
// var count = 0
// var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

// func BForTest() float64 {
// 	count++

// 	if count%3 == 0 {
// 		return 0
// 	}
// 	return rnd.Float64() * 100
// }

func main() {
	result, err := divide(564.6456, 798.7987)
	if err != nil {
		fmt.Println("ERROR:", err)
	} else {
		fmt.Printf("a / b = %.4f\n", result)
	}

	result, err = divide(564.6456, 0)
	if err != nil {
		fmt.Println("ERROR:", err)
	} else {
		fmt.Printf("a / b = %.4f\n", result)
	}

	// for i := 0; i <= 2; i++ {
	// 	a := rnd.Float64() * 100
	// 	b := BForTest()
	// 	fmt.Printf("a = %.4f\nb = %.4f\n", a, b)

	// 	result, err := divide(a, b)
	// 	if err != nil {
	// 		fmt.Println("ERROR:", err)
	// 	} else {
	// 		fmt.Printf("a / b = %.4f\n", result)
	// 	}
	// }

}
