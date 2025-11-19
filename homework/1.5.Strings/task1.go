package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

func main() {
	//хотим читать данные по одной строке за раз, тогда
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введите строку: ")
	input, _ := reader.ReadString('\n')
	stringLine := strings.TrimSpace(input)
	fmt.Printf("Ваша строка: %s\nВ ней %d символов", stringLine, utf8.RuneCountInString(stringLine))
}