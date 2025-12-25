package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// func CountBrackets(s string) [2]int {
// 	bracket := [2]int{}

// 	for _, value := range s {
// 		if value == '(' {
// 			bracket[0]++
// 		} else if value == ')' {
// 			bracket[1]++
// 		}
// 	}
// 	return bracket
// }

func BalanceBrackets(s string) (bool, [2]int) {
	bracket := [2]int{}

	for _, value := range s {
		if value == '(' {
			bracket[0]++
		} else if value == ')' {
			bracket[1]++
		}
	}

	balance := 0
	for _, value := range s {
		if value == '(' {
			balance++
		} else if value == ')' {
			balance--
			if balance < 0 {
				return false, bracket
			}
		}
	}

	return balance == 0, bracket
}

func main() {

	fmt.Println("строка на вход: ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	s := strings.TrimSpace(input)
	// bracket := CountBrackets(s)
	isValid, bracket := BalanceBrackets(s)

	if isValid {
		fmt.Printf("вывод: Скобки расставлены верно, %d открывающиеся, %d закрывающиеся", bracket[0], bracket[1])
	} else {
		fmt.Printf("вывод: Скобки расставлены неверно, %d открывающиеся, %d закрывающиеся", bracket[0], bracket[1])
	}
}
