package main

import (
	"fmt"
	"slices"
	"strings"
)

/*
func SearCity() не изменяет слайс, так что без указателей,
то есть принимаем слайс по значению
*/

func AppendCity(slice *[]string, city string) {
	*slice = append(*slice, city)
}

// https://pkg.go.dev/slices#Delete
// https://pkg.go.dev/slices#IndexFunc
// https://pkg.go.dev/strings#EqualFold
func DeleteCity(slice *[]string, city string) {
	index := slices.IndexFunc(*slice, func(s string) bool {
		return strings.EqualFold(s, city)
	})
	if index != -1 {
		*slice = slices.Delete(*slice, index, index+1)
	}
}

func SearchCity(slice []string, city string) bool {
	return slices.IndexFunc(slice, func(s string) bool {
		return strings.EqualFold(s, city)
	}) != -1
}

func main() {
	testslice := []string{
		"Москва",
		"Санкт-Петербург",
		"Норильск",
		"Красноярск",
		"Сочи",
		"нью-Йорк", //краевой случай, должно быть true
	}

	fmt.Println("Исходный слайс:", testslice)

	AppendCity(&testslice, "Лондон")
	fmt.Println("Добавление нового города:", testslice)

	DeleteCity(&testslice, "Красноярск")
	fmt.Println("Удаление города по имени (ex. Красноярск):", testslice)

	exists := SearchCity(testslice, "Нью-Йорк")
	fmt.Println("Поиска города в списке:", exists)
}
