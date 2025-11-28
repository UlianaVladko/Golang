package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

/*
Функция заполняет массив случайными числами,
а основнеая логика (создание массива, копирование в слайс, сортировка и вывод)
находится в main

Проблемы были с генератором (блокировка со стороны антивируса попытки системного вызова)
Реализовала через создание локального генератора rand.New() с источником на основе
текущего времени.

slice := myarr[:] создает зависимую копию от исходного, то есть будем сортировать исходный массив тоже
*/

// https://pkg.go.dev/math/rand#New
// https://pkg.go.dev/math/rand#NewSource
func FillArray(arr *[10]int) {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano())) //создаем локальный генератор (это меннее подозрительно для антивируса)
	for index := range arr {                               // value не нужен
		arr[index] = rnd.Intn(100) + 1 //при глобальном генераторе rand.Intn программа падает в ошибке Access is denied
	}
}

func main() {
	var myarr [10]int
	FillArray(&myarr)
	fmt.Println(myarr)

	slice := make([]int, len(myarr))
	copy(slice, myarr[:]) //копируем массив в слайс, чтобы не изменять исходный

	sort.Ints(slice)

	fmt.Println("Исходный массив: ", myarr)
	fmt.Println("Отсортированный массив: ", slice)
}
