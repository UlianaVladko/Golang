package main

import "fmt"

func main() {
	a := 1.0 / 3
	fmt.Println(a)
	fmt.Printf("%v\n", a)
	fmt.Printf("%f\n", a)
	fmt.Printf("%.2f\n", a)   //после запятой выводим 2 цифры 0.33
	fmt.Printf("%08.2f\n", a) //00000.33 всего 8 знаков выводим, учитывая точку. при этом первые "пустые" места заполняем нулями
	fmt.Printf("%14.2f\n", a) //10 пробелов в начале

	// const pi float64 = 3.14
	// const (
	// 	long  int = 12
	// 	width int = 12
	// )
	// const long2, width2 = 13, 14

	// /* так объявлять константы нельзя!!!
	// const a int
	// const b
	// */

	// const (
	// 	a = 1
	// 	b //b = ?
	// 	c //c = ?
	// 	d = 2
	// 	e // e = ?
	// )
	// fmt.Println(a, b, c, d, e) //вывод 1, 1, 1, 2, 2
}
