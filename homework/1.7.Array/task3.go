package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
	"unicode"
)

// названия переменных повторяются, чтобы легче прочитать логику. они итак изолированы

// приводим к одному регистру слова в строчке str (связано с unicode)
func SmallString(str string) string {
	return strings.ToLower(strings.TrimSpace(str))
}

// строчку с малым регистром разделяем на слова, игнорируя знаки препинания
func SplitWords(smallStr string) []string {
	separateStr := strings.FieldsFunc(smallStr, func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsNumber(r) && r != '\''
	})
	return separateStr
}

// считаем сепарированные слова
func CountWords(separateStr []string) map[string]int {
	count := make(map[string]int) //создаем мапу для хранения результатов
	for _, value := range separateStr {
		count[value]++
	}
	return count
}

func main() {
	// тестировала I i Don't know Know kNow
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введите строчку: ")
	text, _ := reader.ReadString('\n')
	str := strings.TrimSpace(text)

	smallStr := SmallString(str)
	fmt.Println(smallStr)
	separateStr := SplitWords(smallStr)
	fmt.Println(separateStr)
	count := CountWords(separateStr)
	fmt.Println(count)

	keys := make([]string, 0, len(count))

	for word := range count {
		keys = append(keys, word)
	}
	sort.Slice(keys, func(i, j int) bool {
		return count[keys[i]] > count[keys[j]]
	})

	for _, word := range keys {
		fmt.Printf("'%s': %d\n", word, count[word])
	}
}
