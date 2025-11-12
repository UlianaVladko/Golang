package main

import "fmt"

func main() {
	var a int
	a = 3
	// b = 2
	
	// if c := 4; c > b && c > a {
	// 	fmt.Println("c > b and c > a")
	// }

	switch a {
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	case 3, 4:
		fmt.Println("3 or 4")
	default:
		fmt.Println("Default case")
	}

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