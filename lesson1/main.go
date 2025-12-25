package main

import "fmt"

// "fmt"
// "reflect"

// "fmt"
// printingfunctions "lesson1/PrintingFunctions"

func Cube(x int) int {
	result := x*x*x
	return result
}

func Square(x int) int {
	return x*x
}

func main() {
	// //print("Hello! Students")
	// fmt.Println("Hello it is Println")
	// fmt.Print("Hello from fmt!")
	// printingfunctions.PrintMyAge()
	// var name, surname string
	// name, surname = "Ulyana", "Vladko"
	// // surname = "Vladko"

	// var age = 23
	// height := 1.62
	// height = 3
	// const city = "Saint-Petersburg"

	// // fmt.Println(name, surname, age, reflect.TypeOf(height))
	// fmt.Println(name, surname, age, height, city)

	value := 2
	fmt.Println(Cube(value))
	fmt.Println(Square(value))
}


