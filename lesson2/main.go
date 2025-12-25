package main

import "fmt"

// func NameSurnameAge() {
// 	// Задача1
// 	var name, surname string
// 	var age int
// 	fmt.Print("Enter ur name: ")
// 	fmt.Scan(&name)
// 	fmt.Print("Enter ur surname: ")
// 	fmt.Scan(&surname)
// 	fmt.Print("Enter ur age: ")
// 	fmt.Scan(&age)
// 	fmt.Printf("Hello, %s %s\nI know that you are %d", name, surname, age)
// }

func main() {
	//Задача2
	var passAge int
	fmt.Scan(&passAge)
	fmt.Printf("Возраст: %d - ", passAge)
	if passAge < 7 || passAge > 64 {
		fmt.Println("Проезд бесплатный")
	} else {
		fmt.Println("Проезд 70 рублей")
	}

	// var age int
	// fmt.Print("Enter ur age: ")
	// fmt.Scan(&age)
	// if age <= 17 {
	// 	fmt.Println("You are not of legal age.")
	// } else {
	// 	fmt.Println("You are of legal age.")
	// }

	// name := "Ulyana" //это строка. Для вывода строки используется %s
	// fmt.Println("My name is", name, "San")
	// age := 22 //Это число. Для него - %d
	// fmt.Printf("My name is %v San. I am %v years old", name, age)
	// fmt.Sprintf("Hello, I am %s", name)

	// var name string
	// fmt.Print("Enter ur name: ")
	// fmt.Scan(&name)
	// fmt.Printf("Hello, %s\n", name)

	// NameSurnameAge()

}

/*
ВАЖНЫЙ МОМЕНТ:
fmt.Scan() - считывает ввод до первого пробела
ДАЛЬНЕЙШЕЕ ИЗУЧЕНИЕ - bufio.NewReader() / os
*/

/*
зАДАЧА1:
Считать от пользователя имя, фамилию, возраст.
Вывести эта данные в красивом форматированном выводе в консоль.
*/

/*
ЗАДАЧА2:
Вы пишите терминал в автобусе.
Программный модуль принимает на вход возраст пассажира
Если ему <7 - проезд бесплатный
Если ему от 7 до 64 лет - проезд 70 рублей
Если ему больше 64 - проезд бесплатный
Вывести в консоль на основании вздных данных: сколько должен
заплатить пассажир.
*/
