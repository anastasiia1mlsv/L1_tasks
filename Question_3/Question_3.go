package main

import (
	"fmt"
	"sync"
)

/*
	Дана последовательность чисел: 2,4,6,8,10.
	Найти сумму их квадратов(22+32+42….) с использованием конкурентных вычислений.
*/

func main() {
	array := [5]int{2, 4, 6, 8, 10}
	wg := sync.WaitGroup{}
	sum := 0
	for _, v := range array {
		wg.Add(1)
		go sumOfSquares(v, &sum, &wg) // For each element v, a new goroutine is started by the statement go
	}
	wg.Wait()
	fmt.Print(sum)
}

func sumOfSquares(num int, sum *int, wg *sync.WaitGroup) {
	defer wg.Done()
	*sum += num * num
}

/*
	В Go, когда вы передаете переменную функции,
	вы передаете копию этой переменной, а не исходную.
	Это называется передачей по значению.
	Однако иногда вам нужно передать ссылку на переменную,
	чтобы функция могла изменить исходную переменную.
	Это называется передачей по ссылке и достигается с помощью указателей.
*/
