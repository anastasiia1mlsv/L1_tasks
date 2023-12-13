package main

import "fmt"

/*
	Реализовать пересечение двух неупорядоченных множеств U
*/

func main() {
	arr := []int{1, 2, 3, 4, 5}
	arrToCompare := []int{1, 2, 3}
	arrMap := make(map[int]bool)

	// Заполнение map элементами из первого массива
	for _, num := range arr {
		arrMap[num] = true
	}

	// U юнион операция
	for _, num := range arrToCompare {
		if arrMap[num] {
			fmt.Printf("Общий элемент: %d\n", num)
		}
	}
}
