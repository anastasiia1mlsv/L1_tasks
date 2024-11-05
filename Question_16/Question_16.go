package main

import "fmt"

/*
	Реализовать быструю сортировку массива (quicksort) встроенными методами языка.
*/

func main() {
	unsorted := []int{63, 5, 1, 7, 10, 15, 6}
	fmt.Println("Unsorted:", unsorted)
	quickSort(unsorted, 0, len(unsorted)-1) // передаем на работу первый и последний индекс
	fmt.Println("Sorted:", unsorted)
}

func quickSort(array []int, low, high int) {
	if low < high {
		// Сравним каждый элемент начиная с 0 с последним
		// Если i элемент меньше мы двигаем его в левую часть
		// Если больше мы ничего с ним не делаем таким образом меньшие числа останутся слева
		// Справа останутся большие, между ними будет Pivot
		// Далее считаем стороны по отдельности
		// И так далее в след подмножесте опять будем искать опорную точку
		pivot := partition(array, low, high) // Разбиение
		quickSort(array, low, pivot-1)       // Рекурсивно сортируем левую часть
		quickSort(array, pivot+1, high)      // Рекурсивно сортируем правую часть
	}
}

func partition(array []int, low, high int) int {
	pivot := array[high] // Выбор опорного элемента Pivot
	i := low
	for j := low; j < high; j++ {
		if array[j] < pivot {
			array[i], array[j] = array[j], array[i]
			i++
		}
	}
	array[i], array[high] = array[high], array[i]
	return i
}
