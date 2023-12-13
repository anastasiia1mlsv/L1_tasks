package main

import (
	"fmt"
	"math"
)

/*	Разработать программу нахождения расстояния
	между двумя точками, которые представлены
	в виде структуры Point с инкапсулированными
	параметрами x,y и конструктором.
*/

type Point struct {
	x, y float64
}

// NewPoint - конструктор для Point
func NewPoint(x, y float64) *Point {
	return &Point{x: x, y: y}
}

// calcDist вычисляет евклидово расстояние между двумя точками
func calcDist(p1, p2 *Point) float64 {
	return math.Sqrt(math.Pow(p2.x-p1.x, 2) + math.Pow(p2.y-p1.y, 2))
}

func main() {
	dot1 := NewPoint(1, 1)
	dot2 := NewPoint(2, 2)

	fmt.Printf("Расстояние: %.2f\n", calcDist(dot1, dot2))
}
