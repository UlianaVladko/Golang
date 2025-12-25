package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

type Circle struct {
	Radius float64
}

type Rectangle struct {
	Length float64
	Width  float64
}

func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.Radius, 2)
}

func (r Rectangle) Area() float64 {
	return r.Length * r.Width
}

// Напишите функцию,
// которая принимает слайс фигур ([]Shape) и выводит площадь каждой фигуры
func printAreas(shapes []Shape) {
	for index, value := range shapes {
		area := value.Area()

		switch s := value.(type) {
		case Circle:
			fmt.Printf("Figure #%d. The area of a circle with radius %.1f is equal to %.2f\n",
				index+1, s.Radius, area)
		case Rectangle:
			fmt.Printf("Figure #%d. The area of a rectangle with length %.1f and width %.1f is equal to %.2f\n",
				index+1, s.Length, s.Width, area)
		default:
			fmt.Println("default")
		}
	}
}

func main() {
	shape := []Shape{
		Circle{Radius: 13.2},
		Rectangle{Length: 4.5, Width: 1.2},
		Circle{Radius: 5.0},
		Rectangle{Length: 10.0, Width: 3.0},
	}
	printAreas(shape)
}
