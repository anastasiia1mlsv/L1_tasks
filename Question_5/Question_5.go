package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
	Разработать программу, которая будет последовательно отправлять
	значения в канал, а с другой стороны канала — читать.
	По истечению N секунд программа должна завершаться.
*/

func main() {
	mainChannel := make(chan int)
	completionChannel := make(chan struct{})
	const N = 5
	go publisher(mainChannel, completionChannel)
	go subscriber(mainChannel, completionChannel)

	<-time.After(N * time.Second)
	close(completionChannel)    // Отправляем сигнал о завершении
	time.Sleep(1 * time.Second) // Даем время на обработку сигнала о завершении
	close(mainChannel)          // Закрываем канал
}

func publisher(mainChan chan int, done <-chan struct{}) {
	for {
		select {
		case <-done:
			return
		default:
			v := rand.Intn(100)
			mainChan <- v
			fmt.Println("Message sent: ", v)
			time.Sleep(1 * time.Second)
		}
	}
}

func subscriber(mainChan chan int, done <-chan struct{}) {
	for {
		select {
		case <-done:
			return
		case v := <-mainChan:
			fmt.Println("Message received: ", v)
		}
	}
}
