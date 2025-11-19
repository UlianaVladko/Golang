package main

import (
	"fmt"
	"strings"
	"unicode"
)

func CapitalizeWords(s string) string {
	var sep strings.Builder
	flagToUpper := true

	for _, value := range s {
		if unicode.IsSpace(value) {
			sep.WriteRune(value)
			flagToUpper = true
		} else if value == '-' {
			sep.WriteRune(value)
			flagToUpper = true
		} else {
			if flagToUpper {
				sep.WriteRune(unicode.ToUpper(value))
				flagToUpper = false
			} else {
				sep.WriteRune(unicode.ToLower(value))
			}
		}

	}
	return sep.String()
}

func main() {
	testvalue := []string{
		"привет, мир!",
		"  в начале натыкала проблелов      и здесь тоже",
		"санкт-петербург",
		"заБЫЛА отКЛЮчИть cAps LocK",
	}

	for _,value := range testvalue{
		result:=CapitalizeWords(value)
		fmt.Printf("%s -> %s\n", value, result)
	}

}
