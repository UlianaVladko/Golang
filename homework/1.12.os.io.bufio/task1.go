package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("logfile.log")
	if err != nil {
		log.Fatal("Error opening file:", err)
		os.Exit(1)
	}
	defer file.Close()

	countErr, countLin := 0, 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		countLin++
		line := scanner.Text()

		if strings.Contains(line, "error") {
			countErr++
			// fmt.Printf("Found error at line %d: %s\n", countLin, line)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal("Error reading file:", err)
	}

	fmt.Printf("Total lines with 'error': %d", countErr)

	// data := make([]byte, 64)

	// for {
	// 	n, err := file.Read(data)
	// 	if err == io.EOF {
	// 		break
	// 	}
	// 	fmt.Print(string(data[:n]))
	// }

}
