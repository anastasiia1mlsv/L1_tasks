package main

import (
	"fmt"
)

/*
	Разработать конвейер чисел.
	Даны два канала:
	в первый пишутся числа (x) из массива,
	во второй — результат операции x*2,
	после чего данные из второго канала должны выводиться в stdout.
*/

func main() {
	// Исходный массив чисел
	numbers := [5]int{0, 1, 2, 3, 4}

	// Канал для отправки исходных чисел
	inputChannel := make(chan int)
	// Канал для отправки удвоенных чисел
	doubledChannel := make(chan int)
	// Канал для управления завершением горутин
	doneChannel := make(chan struct{})

	// Горутина для удвоения чисел
	go func() {
		for {
			select {
			case num := <-inputChannel:
				doubledChannel <- num * 2
			case <-doneChannel:
				close(doubledChannel)
				return
			}
		}
	}()

	// Горутина для вывода удвоенных чисел
	go func() {
		for doubledNum := range doubledChannel {
			fmt.Printf("%d ", doubledNum)
		}
	}()

	// Отправка чисел в inputChannel
	for _, num := range numbers {
		inputChannel <- num
	}
	close(inputChannel) // Закрытие inputChannel после отправки всех чисел

	// Посылаем сигнал в doneChannel, чтобы завершить первую горутину
	doneChannel <- struct{}{}
	close(doneChannel) // Закрытие doneChannel
}
