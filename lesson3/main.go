package main

// //Фишка, вывести ядерку своего ноутбука
// import (
// 	"fmt"
// 	"strconv"
// )

// func main() {
// 	fmt.Println(strconv.IntSize)
// }

// //Еще фишка, вывести максимальное значения типа переменной
// import (
// 	"fmt"
// 	"math"
// )

// func main() {
// 	fmt.Println(math.MaxFloat64)
// 	fmt.Println(math.MaxInt8)
// }

import (
	"fmt"
	"time"
)

/*
	Задача1
	Мы пишем программу для пункта обмена валют.
	В главной функции main определить 2 переменные:
	course = 84.8 (84.8 денег - 1 тугрик), value = 30_000 денег
	Напишите и вызовите функцию в main, которая переводит value в какую-то валюту по курсу.
	Вывести это значение (сколько денег надо отдать клиенту)
*/
// //Задача1
// func exchange(value, course float64) float64 {
// 	return value / course
// }

// func main() {
// 	//Задача1
// 	fmt.Printf("%.4f тугриков\n", exchange(30_000, 84.8))
// }

/*
	Задача2
	Вывести время time.Now() в формате:
	06 Nov 2025, 21:22
	06.11.25 9:22
	2025-06-11 21-23-50
	Nov, 06.
*/

func main() {
	//Задача2
	myTime :=time.Now()
	fmt.Println(myTime.Format("02 Jan 2006, 15:04"))
	fmt.Println(myTime.Format("02.01.06 03:04"))
	fmt.Println(myTime.Format("2006-02-01 15-04-05"))
	fmt.Println(myTime.Format("Jan, 02."))

	/*Дата и время в Golang
	01 - месяц (январь)
	02 - число (2 января)
	03 - час (2 января 3 часа)
	04 - минута (2 января 3 часа 4 минуты)
	05 - секунда (2 января 3 часа 4 минуты 5 секунд)
	06 - год (2 января 3 часа 4 минуты 5 секунд 2006 года)
	*/

	// template := "2006-01-02"
	// myDate := "2025-11-06"

	// myTime, err := time.Parse(template, myDate)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(myTime)
	// fmt.Printf("%T", myTime)

	// //Дата и время с собственном формате
	// myTime := time.Now()
	// fmt.Println(myTime)
	// fmt.Println(myTime.Format("02.01.06 15:04"))
	// fmt.Println(myTime.Format("2006 Jan 02 03:04:05"))

	// var a int
	// fmt.Printf("%T", a) //посмотреть тип переменной %T

	// //Шестнадцатиричные числа
	// var _, _, _ = 0, 120, 55 //нижнее подчеркивание _ это анонимная переменная, она игнорируется golang-ом
	// var red16, _, _ = 0x00, 0x7b, 0xff
	// fmt.Printf("%x", red16) //%x для вывода шестнадцатирички

	// a := 1.0 / 3
	// fmt.Println(a + a + a)

	// b := 0.1
	// b += 0.2
	// // fmt.Println(b)
	// // fmt.Printf("%.60f\n", 0.1)
	// // fmt.Printf("%.60f\n", 0.2)
	// fmt.Println((b - 0.3) < 0.00001)

	// a := 1.0 / 3
	// fmt.Println(a)
	// fmt.Printf("%v\n", a)
	// fmt.Printf("%f\n", a)
	// fmt.Printf("%.2f\n", a)   //после запятой выводим 2 цифры 0.33
	// fmt.Printf("%08.2f\n", a) //00000.33 всего 8 знаков выводим, учитывая точку. при этом первые "пустые" места заполняем нулями
	// fmt.Printf("%14.2f\n", a) //10 пробелов в начале

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
