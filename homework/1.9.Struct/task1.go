package main

import "fmt"

// https://sdo.bmstu.ru/mod/page/view.php?id=15988
// структура
type Book struct {
	Title  string
	Author string
	Year   int
}

// метод
func (b Book) GetInfo() string {
	return fmt.Sprintf("\"%s\" by %s, %d.\n", b.Title, b.Author, b.Year)
}

func main() {
	book := Book{
		Title:  "The Great Gatsby",
		Author: "F. Scott Fitzgerald",
		Year:   1925,
	}
	fmt.Println(book.GetInfo())
}
