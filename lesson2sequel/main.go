package main

import "fmt"

func main() {
	var a, b int
	a = 1
	b = 2
	
	incB := func() bool{
		b += 1 // b =3
		return true
	}
	if (a == 1) /* true */ || incB() /* даж не выполнялось */ {
		fmt.Println("All good")
	} else {
		fmt.Println("Bad value")
	}

	fmt.Println(a,b) //предполагаем 1, 3; НО выводит 1, 2
}