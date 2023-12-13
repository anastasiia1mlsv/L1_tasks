package main

import "fmt"

/*
	Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0.
*/

// Устанавливает i-й бит в 1
func setBit(x, i int64) int64 {
	return x | (1 << (i - 1))
}

// Сбрасывает i-й бит в 0
func clearBit(x, i int64) int64 {
	return x &^ (1 << (i - 1)) // И-НЕ
}

func main() {
	x := int64(777)
	fmt.Println("Исходное число:", x)

	x = setBit(x, 2) // Устанавием второй бит в 1
	fmt.Println("Число после установки 2-го бита в 1:", x)

	x = clearBit(x, 2) // Сбросим второй бит в 0
	fmt.Println("Число после сброса 2-го бита в 0:", x)
}
