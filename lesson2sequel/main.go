package main

import "fmt"

func main() {

/*
ЗАДАЧА
Написать прграмму на ЯП Go с следующим функционалом:
В бесконечном цикле считывать от пользователя числа.
И складывать их друг с другом.
Если пользователь ввел 0 (как число) - вывести на
экран полученную сумму и выйти из программы.
*/
var sum int
fmt.Println("Начинайте вводить числа")
	for {
		var value int
		fmt.Scan(&value)
		sum +=value
		if value == 0{
			break
		}
	}
fmt.Println("Сумма: ", sum)

	// v := 0
	// for i := 0; i < 10; i++ {
	// 	// break - закончить цикл
	// 	// continue - перейти к следующей иттерации
	// 	if i%2 != 0 || i == 0 {
	// 		continue
	// 	}
	// 	fmt.Println(i)

		// v++
		// fmt.Println(v)
		// if v >= 5 {
		// 	break
		// }
	// }

	// v := 0
	// i := 0
	// for i < 10 {
	// 	v++
	// 	fmt.Println(v)
	// 	i++
	// }

	// бесконечный цикл
	// v := 0
	// for {
	// 	v++
	// 	fmt.Println(v)
	// }

	// v := 0
	// for i := 0; i < 10; i++ {
	// 	v++
	// 	fmt.Println(v)
	// }

	// var a int
	// a = 3
	// b = 2

	// if c := 4; c > b && c > a {
	// 	fmt.Println("c > b and c > a")
	// }

	// switch {
	// case a >= 3:
	// 	fmt.Println("more than 3")
	// 	fallthrough
	// case a >= 2:
	// 	fmt.Println("more than 2")
	// 	fallthrough
	// case a >= 1:
	// 	fmt.Println("more than 1")
	// default:
	// 	fmt.Println("Default case")
	// }

	// incB := func() bool{
	// 	b += 1 // b =3
	// 	return true
	// }
	// if (a == 1) /* true */ || incB() /* даж не выполнялось */ {
	// 	fmt.Println("All good")
	// } else {
	// 	fmt.Println("Bad value")
	// }

	// fmt.Println(a,b) //предполагаем 1, 3; НО выводит 1, 2
}
