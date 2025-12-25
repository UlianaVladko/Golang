package main

import (
	"errors"
	"fmt"
	"os"
	"time"
)

// error
type TimeError struct {
	Time time.Time
	Err  error
}

// TimeError ~ error ? error - interface {Error() string}

func (te *TimeError) Error() string {
	return fmt.Sprintf("[%v]: %v", te.Time.Format("2006/01/02 15:04:05"), te.Err)
}

func (te *TimeError) Unwrap() error {
	return te.Err
}

func NewTimeError(err error) error {
	return &TimeError{
		Time: time.Now(),
		Err:  err,
	}
}

func ReadFile(fileName string) (string, error) {
	data, err := os.ReadFile(fileName) //err = os.FileNotExist = "Error file not Exist"
	if err != nil {
		return "", NewTimeError(err) // TimeError {Time: 11:26, err: os.FileNotExist}
	}
	return string(data), nil
}

// ReadFile -> no such file or directiory
// os -> open filename: no such file or directiory
// TimeError -> [time]: open filename: no such file or directiory

// func Is(err, target error) bool {}
// io.EOF
// err == io.EOF - fail!!!

func main() {
	_, err := ReadFile("config.txt")
	if err != nil {
		// if errors.Is(err, os.ErrNotExist) {
		// 	// упустим функционал
		// 	fmt.Println("File Created.")
		// } else {
		// 	fmt.Println("такой ошибки не знаю и не орабатываю:", err)
		// 	os.Exit(1)
		// }

		/*
			if myErr, ok = err.(*TimeError); ok {
			// ...
			}
		*/

		var myErr *TimeError
		if errors.As(err, &myErr) {
			fmt.Println(myErr.Error())
			fmt.Println(myErr.Err)
			fmt.Println(myErr.Time)
		}

		// fmt.Println(err) // с временем
		// // скорлупа на скорлупе
		// fmt.Printf("Изначальная ошибка os: %v\n", errors.Unwrap(err))
		// fmt.Printf("Изначальная ошибка ReadFile: %v\n", errors.Unwrap(errors.Unwrap(err)))
		// fmt.Printf("Изначальная ошибка ReadFile: %v\n", errors.Unwrap(errors.Unwrap(errors.Unwrap(err))))

		// os.Exit(1)
	}
}

/*
Домашняя работа
package main
import (
	"fmt"
	"os"
)
// LabelError описывает ошибку с дополнительной меткой.
type LabelError struct {
	Label string // метка должна быть в верхнем регистре
	Err   error
}
// добавьте методы Error() и NewLabelError(label string, err error) error
// ...
func main() {
	_, err := os.ReadFile("mytest.txt") //net.Dial() err-> NewLableEror("net", err)
	if err != nil {
		err = NewLabelError("file", err)
	}
	fmt.Println(err)
	// должна выводить
	// [FILE] open mytest.txt: no such file or directory
}
*/
