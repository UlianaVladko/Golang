package main

import (
	"fmt"
	"sync"
	"time"
)

// структура
type Executor struct {
	wg sync.WaitGroup
}

// конструктор
func NewExecutor() *Executor {
	return &Executor{}
}

/*
метод

	func (receiver Type) MethodName(parameters) returnType {
	    // тело метода
	}
*/

// https://pkg.go.dev/sync#WaitGroup.Add
// https://pkg.go.dev/sync#WaitGroup.Done

func (e *Executor) RunGoroutine(index int, duration time.Duration) {
	e.wg.Add(1)

	go func(runIndex int, d time.Duration) {
		defer e.wg.Done()

		fmt.Printf("Горутина %d начала работу (продолжительность работы %v)\n", runIndex, d)
		time.Sleep(d)
		fmt.Printf("Горутина %d завершила работу\n", runIndex)
	}(index, duration)
}

// https://pkg.go.dev/sync#WaitGroup.Wait
func (e *Executor) WaitAll() {
	e.wg.Wait()
}

func main() {
	executor := NewExecutor()

	executor.RunGoroutine(1, 3*time.Millisecond)
	executor.RunGoroutine(2, 1*time.Millisecond)
	executor.RunGoroutine(3, 2*time.Millisecond)

	fmt.Println("Основной поток дожидается завершения всех горутин")
	executor.WaitAll()
	fmt.Println("В основном потоке завершены все горутины")
}
