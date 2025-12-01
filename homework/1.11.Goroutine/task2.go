package main

import "fmt"

/*
https://habr.com/ru/articles/904892/

pipeline — это последовательность этапов обработки данных,
где выход одного этапа становится входом для другого.
*/

func generate() <-chan int {
	out := make(chan int)
	go func() {
		for i := 1; i <= 10; i++ {
			out <- i
		}
		close(out)
	}()
	return out
}

func consumer(in <-chan int) <-chan struct{} {
	done := make(chan struct{})
	go func() {
		for index := range in {
			fmt.Println(index)
		}
		close(done)
	}()
	return done
}

func main() {
	done := consumer(generate())
	<-done
}
