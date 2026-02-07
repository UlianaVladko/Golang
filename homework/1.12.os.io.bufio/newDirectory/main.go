package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	// Получаем список файлов в текущей директории
	files, err := os.ReadDir(".")
	if err != nil {
		fmt.Println("Ошибка чтения директории:", err)
		return
	}

	var textFiles []string
	for _, f := range files {
		if !f.Type().IsRegular() {
			continue // пропускаем папки
		}
		name := f.Name()
		// Проверяем расширение текстового файла
		if isTextFile(name) {
			textFiles = append(textFiles, name)
		}
	}

	if len(textFiles) == 0 {
		fmt.Println("Нет текстовых файлов для объединения")
		return
	}

	// Сортируем имена файлов по алфавиту
	sort.Strings(textFiles)

	// Создаём combined.txt
	outFile, err := os.Create("combined.txt")
	if err != nil {
		fmt.Println("Ошибка создания файла:", err)
		return
	}
	defer outFile.Close()

	writer := bufio.NewWriter(outFile)

	for _, file := range textFiles {
		err := appendFile(file, writer)
		if err != nil {
			fmt.Printf("Ошибка при обработке %s: %v\n", file, err)
			return
		}
	}

	writer.Flush()
	fmt.Println("Файлы успешно объединены в combined.txt")
}

// Проверка, что файл текстовый (по расширению)
func isTextFile(name string) bool {
	lower := strings.ToLower(name)
	// Только .txt
	return strings.HasSuffix(lower, ".txt")
}

// Чтение одного файла и запись в writer
func appendFile(filename string, writer *bufio.Writer) error {
	inFile, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer inFile.Close()

	scanner := bufio.NewScanner(inFile)
	var lastEmpty bool
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			if lastEmpty {
				continue // пропускаем лишние пустые строки
			}
			lastEmpty = true
		} else {
			lastEmpty = false
		}
		writer.WriteString(line + "\n")
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	// Добавляем ровно один перенос строки между файлами
	writer.WriteString("\n")
	return nil
}
