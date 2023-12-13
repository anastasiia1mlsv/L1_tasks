package main

import "fmt"

/*
Удалить i-ый элемент из слайса.
*/

func main() {
	//задаем и выводим для проверки исходный слайс
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(slice)
	//индекс для удаления (начинаем с 0)
	deletable := 3
	sliceTemp := slice[0:deletable]
	//заполняем новый слайс со всеми элементами из первого слайса, кроме одного
	for i := deletable + 1; i < len(slice); i += 1 {
		sliceTemp = append(sliceTemp, slice[i])
	}
	//обновляем первый слайс с уже отсутствующим элементом
	slice = sliceTemp
	fmt.Println(slice)
}
