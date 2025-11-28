package main

import "fmt"

type Student struct {
	Name   string
	Grades []float64
}

func (s Student) AverageGrade() float64 {
	// конструкция из домашки 1.3. Основные алгоритмические конструкции языка Go
	var sum float64
	for _, value := range s.Grades {
		sum += value
	}

	return sum / float64(len(s.Grades))
}

func main() {
	student := Student{
		Name:   "Владко Ульяна",
		Grades: []float64{4, 5, 4, 4, 3, 4, 3, 3}, // просто оценик обычно int, а вывод будет float
	}
	fmt.Printf("Средний бал: %.1f", student.AverageGrade())
}
