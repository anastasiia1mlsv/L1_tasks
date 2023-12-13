package main

import "fmt"

/*
	Реализовать бинарный поиск встроенными методами языка.
*/

func main() {
	numbers := []int{1, 4, 10, 11, 18, 24, 28, 50}

	target := 4

	fmt.Println(binarySearch(numbers, target))
}

func binarySearch(sortedNums []int, target int) bool {
	mid := len(sortedNums) / 2
	for {
		// Если средний элемент больше целевого, ищем в левой половине
		if sortedNums[mid] > target {
			sortedNums = sortedNums[:mid]
			mid /= 2
		} else if sortedNums[mid] == target {
			// Если средний элемент является целью, возвращаем true
			return true
		} else {
			// Если средний элемент меньше целевого, ищем в правой половине
			sortedNums = sortedNums[mid:]
			mid /= 2
		}
		// Если Mid становится равным 0, элемент не найден
		if mid == 0 {
			return false
		}
	}
}

/*
	Бинарный поиск — это алгоритм «разделяй и властвуй».
	На каждом шаге он делит массив на половины
	и работает с той половиной, которая может содержать целевой элемент.
*/
