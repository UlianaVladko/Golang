package main

import "fmt"

func main() {
	const pi float64 = 3.14
	const (
		long  int = 12
		width int = 12
	)
	const long2, width2 = 13, 14
	
	/* так объявлять константы нельзя!!!
	const a int
	const b
	*/
	
	const (
		a = 1
		b //b = ?
		c //c = ?
		d = 2
		e // e = ?
	)
	fmt.Println(a, b, c, d, e) //вывод 1, 1, 1, 2, 2
}
