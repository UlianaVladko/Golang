package main

import (
	"fmt"
	"slices"
)

/*
Теоретический материал для дз находится на
https://github.com/golang/go/blob/master/src/sort/sort.go
*/

// func SortFloats64Asc(x []float64) {
// 	slices.Sort(x)
// }

func SortFloats64Desc(x []float64) { //x изорилованная переменная - важно помнить!
	slices.Sort(x)
	slices.Reverse(x)
}

func main() {
	fmt.Println("Вход программы 5 натур чисел: ")
	var n1, n2, n3, n4, n5 float64
	fmt.Scan(&n1, &n2, &n3, &n4, &n5)
	n := []float64{n1, n2, n3, n4, n5} // мой динамический массив на всю функцию main

	// Floats64Asc(n)
	// fmt.Println("По возрастанию ", n)

	SortFloats64Desc(n)
	fmt.Println("По убыванию ", n)

	/*
	   https://pkg.go.dev/math#Max
	   https://pkg.go.dev/math#Min
	*/

	maxN := slices.Max(n)
	fmt.Println("Максимальное число: ", maxN)

	minN := slices.Min(n)
	fmt.Println("Минимальное число: ", minN)

	/*
	   Это конструкция range loop (мб удобно, так как у меня инициализирован срез массива)

	   for index, value := range slice {
	   	// тело цикла
	   }

	   нижнее подчеркивание _ это анонимная переменная (индекс нам тут не нужен)
	*/

	var sum float64
	for _, value := range n {
		sum += value
	}

	average := sum / float64(len(n))
	fmt.Printf("Среднее арифметическое: %.4f", average)
}
