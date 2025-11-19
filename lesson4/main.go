package main

import "fmt"

func main() {
	// 	fmt.Println("Text above \n Text \t down")
	// 	fmt.Println(`Text above \n Text \t down`)
	// 	fmt.Println(`Text above
	//  Text 	 down`)
	// fmt.Printf("%v - %T", `text in backs`, `text in backs`) // дублируем первый аргумент для %T (нулевой аргумент это само "%v - %T")
	// fmt.Printf("%v - %[1]T", `text in backs`)               //спецификатор %[1]T говорит, что для вывода типа функции передаем первый аргумент

	var emoji rune = 128515 //руны
	fmt.Printf("%c", emoji)

	var symbol rune = 'A'
	fmt.Printf("%v", symbol) // обратное преобразование
}
