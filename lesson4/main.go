package main

import (


)
/*
Задача:
		Написать программу на ЯП Go cо следующим функционалом:
		внутри функции main создать пеменную phoneNumber := "89995431232"; "+7(999)1232321" "8-999-321-32-43"
		потом в main выозвите функцию isPhoneNumber(phoneNumber)
		isPhoneNumber - функци, примнимает на вход строку, возвращает строку - только цифры без других символов и проверяет равно ли длинна 11?
		подсказка: пройти по всему номеру с помощью range;
		unicode.IsDigit(rune) - проверка нав цифру
		Пусть выводится на экран: только номер без символов и bool номер/не номер телефона
		package main
import (
  "fmt"
  "unicode"
)
func isPhone(number string) string {
  result := ""
  for _, v := range number {
    fmt.Printf("%c,\n", v)
    if unicode.IsDigit(v) {
      result += fmt.Sprintf("%c", v)
    } else {
      continue
    }
  }
  return result
}
func main() {
  var num string = "8(765)4233141"
  fmt.Println(isPhone(num))
}
*/
func main()  {
	
}



// //шифр Цезаря
// func cypto(letter rune) string {
// 	shifered := letter + 3
// 	if shifered > 'z' {
// 		shifered = shifered - 26
// 	}
// 	return (fmt.Sprintf("%c", shifered))
// }


	// fmt.Println("Введите строку:")
	// reader := bufio.NewReader(os.Stdin)
	// input, err := reader.ReadString('\n')
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(input)

	// //strings pack
	// string1:="My name is Ulya"
	// result:=strings.Replace(string1, "Ulya","Ulyana", -1) //-1 значит заменить все
	// fmt.Println(result)

	// AsciiQuestion := "Hello?"
	// fmt.Println(len(AsciiQuestion))

	// question := "¿Cómo estás?"
	// fmt.Println(len(question)) //неверно
	// fmt.Println(utf8.RuneCountInString(question)) //верно
	// letter, size := utf8.DecodeRuneInString(question)
	// fmt.Printf("if letter %c %v bytes", letter, size)
	// for index, value:=range question{ //байтойное представление - индекс руны в строке, учитывая ее байтовое представление (т е занимает два байиа)
	// 	fmt.Printf("%v - %c\n", index, value)
	// }

	// 	fmt.Println("Text above \n Text \t down")
	// 	fmt.Println(`Text above \n Text \t down`)
	// 	fmt.Println(`Text above
	//  Text 	 down`)
	// fmt.Printf("%v - %T", `text in backs`, `text in backs`) // дублируем первый аргумент для %T (нулевой аргумент это само "%v - %T")
	// fmt.Printf("%v - %[1]T", `text in backs`)               //спецификатор %[1]T говорит, что для вывода типа функции передаем первый аргумент

	// var emoji rune = 128515 //руны
	// fmt.Printf("%c", emoji)

	// var symbol rune = 'A'
	// fmt.Printf("%v", symbol) // обратное преобразование

	// message := "Hello!!"
	// fmt.Println(message[5])
	// fmt.Printf("%c", message[5])
	// // message[5] := `?` - так делать нельзя !!!

	// // вариант вывода символом построчно
	// message := "Hello!!"
	// for i := 0; i < 7; i++ {
	// 	fmt.Printf("%c\n", message[i])
	// }

	//Шифр Цезаря. каждый символ кодируем +3, например a будет d
	// a:='a'
	// fmt.Println(a, a+3)
	// d:='d'
	// fmt.Println(d)
	// 	var value rune = 'z'
	// 	fmt.Println(cypto(value))

