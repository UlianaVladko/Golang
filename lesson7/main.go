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

// Особенные функции init()
var name string

func init() {
	name = "Evgeniy"
}

func main() {
	fmt.Println(name)

	// PrintVoice("cat", Say)

	// result := Sum(1, 2, 3, 4, 6)
	// fmt.Println(result)

	// result := Cube(2, 3)
	// fmt.Println(result)
}
