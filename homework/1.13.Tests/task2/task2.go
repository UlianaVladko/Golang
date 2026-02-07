package main

func sumAll(setNumbers ...int) int {
	sum := 0
	for _, value := range setNumbers {
		sum += value
	}
	return sum
}