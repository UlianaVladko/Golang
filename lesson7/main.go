package main

import "fmt"

func main()  {
	
}

// func Cube(x int) (res int) {
// 	res = x * x * x
// 	return
// }

// func Cube(x, y int) int { //или (x int,  y int)
// 	res := x * y
// 	return res
// }

/*
Возвращаемые значения
Ex., если мы где-то принимаем два инта, а потом где-то возвращаем, то
returns2ints()(x int, y int)
input2ints(returns2ints)
*/

// func Sum(x ...int) (res int){ //...int под капотом слайс []int, вспоминай домашку 1.3
// 	for _, value := range x{ //index _ это игнорирование переменной (не нужна тут)
// 		res += value
// 	}
// 	return
// }

// //Функции первого класса func PrintVoice(animal string, how func(string) string)
// func Say(animal string) string {
// 	if animal == "dog" {
// 		return "wuf"
// 	} else if animal == "cat" {
// 		return "myau"
// 	} else {
// 		return "nah"
// 	}
// }

// func PrintVoice(animal string, how func(string) string){
// 	fmt.Println(how(animal))
// }

// // Особенные функции init()
// var name string

// func init() {
// 	name = "Evgeniy"
// }

// func main() {
// 	// Оператор отложенного вызова defer
// 	fmt.Println("1st output")
// 	defer fmt.Println("2nd output")
// 	defer fmt.Println("3nd output")
// 	// defer file.Close()

// 	// fmt.Println(name)

// 	// PrintVoice("cat", Say)

// 	// result := Sum(1, 2, 3, 4, 6)
// 	// fmt.Println(result)

// 	// result := Cube(2, 3)
// 	// fmt.Println(result)
// }

/*
package main
import "fmt"
func foo() {
	defer func() {
		r := recover()
		if r != nil {
			fmt.Println("Я поймал панику, она у меня!")
		}
	}()
	panic("это паника в функции foo()")
}
// func foo() {
//     defer func() {
//         r := recover()
//         fmt.Printf("Recover вернул: %v, тип: %T\n", r, r)
//         if r != nil {
//             fmt.Println("Паника перехвачена!")
//         }
//     }()
//     panic("это паника в функции foo()")
// }
func main() {
	foo()
}

// Пример Дениса
package main

import (
	"log"
)

func catchMe() {
	panic("Attention! Attention! You can't do that.")
}

func main() {
	defer func() {
		if err := recover(); err != nil {
			log.Println("Panic occurred:", err)
		}
	}()

	catchMe()
}

func main() {
    fmt.Println("1. До catchMe")
    
    defer func() {
        fmt.Println("3. Defer выполнился")
        if err := recover(); err != nil {
            fmt.Println("4. Поймал панику:", err)
        }
    }()
    
    catchMe() // 2. Паника здесь!
    
    fmt.Println("Это никогда не выполнится")
}
*/