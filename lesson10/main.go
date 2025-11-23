package main

import (
	"fmt"
)

func main() {
	c := make(chan int, 2) //c - буферизированный канал
	c <- 2
	fmt.Println(<-c)
	c <- 3
	c <- 4
	fmt.Println(<-c)
	// fmt.Println(<-c)
}

// func PrintInGoroutine(i int, c chan int) {
// 	time.Sleep(time.Duration(rand.Intn(4000)) * time.Millisecond)
// 	c <- i
// }
// func main() {
// 	c := make(chan int)
// 	for i := 0; i < 5; i++ {
// 		go PrintInGoroutine(i, c)
// 	}
// 	timeout := time.After(2 * time.Second)
// 	for i := 0; i < 5; i++ {
// 		select {
// 		case goId := <-c:
// 			fmt.Printf("Routine #%d finished\n", goId)
// 		case <-timeout:
// 			fmt.Println("timeout")
// 			return
// 		}
// 	}
// }

// func PrintInGoroutine(index int, c chan int) {
// 	time.Sleep(1 * time.Second)
// 	c <- index
// 	// fmt.Printf("Routine #%d\n", i)
// 	// fmt.Println("Hi from Parallel")
// }

// func main() {
// 	c := make(chan int)
// 	// go PrintInGoroutine()
// 	for i := 0; i < 5; i++ {
// 		go PrintInGoroutine(i, c)
// 	}
// 	for i := 0; i < 5; i++ {
// 		goID := <-c
// 		fmt.Printf("Routine #%d finished\n", goID)
// 	}

// time.Sleep(2 * time.Second)

/*
	c := make(chan int)
	c <- 1
	<- c //якобы освободили канал
	v := <- c //v переменная, которая примет из канала значение, и оно запишется в v
*/
// }

/*
Пример Денис

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func PrintInGoroutine(i int, c chan int) {
	time.Sleep(time.Duration(rand.Intn(4000)) * time.Millisecond)
	c <- i
}

func main() {
	c := make(chan int)
	for i := 0; i < 5; i++ {
		go PrintInGoroutine(i, c)
	}
	time.Sleep(3 * time.Second)
	for i := 0; i < 6; i++ {
		select {
		case goId := <-c:
			fmt.Printf("Routine #%d finished\n", goId)
		default:
			fmt.Println("No data available on the channel")
		}
	}
}
*/
