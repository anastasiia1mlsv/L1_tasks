package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

/*
	Реализовать постоянную запись данных в канал (главный поток).
	Реализовать набор из N воркеров, которые читают
	произвольные данные из канала и выводят в stdout.
	Необходима возможность выбора количества воркеров при старте.
*/

func main() {

	mainChannel := make(chan int) // проинициализировали интовый канал
	const workerCount = 5         // для контроля количества рабочих горутин
	wg := sync.WaitGroup{}

	for i := 0; i < workerCount; i++ {
		// Before starting each goroutine, wg.Add(1) is called to increment
		// the WaitGroup's counter.
		// This indicates to the WaitGroup that there is one more goroutine to wait for.
		wg.Add(1)                      // Increments the WaitGroup counter by 1
		go worker(mainChannel, &wg, i) // для каждого i запускаем горутину
	}

	// Возможность остановки по сигналу Control C
	stopChan := make(chan os.Signal, 1)
	// Уведоми stopChan когда SIGINT т.е. Control C получен
	signal.Notify(stopChan, os.Interrupt, syscall.SIGINT)

	go func() {
		for {
			select {
			case <-stopChan: // Stop signal received
				close(mainChannel) // Close the mainChannel
				return
			default:
				mainChannel <- rand.Intn(100) // Send random data
				time.Sleep(1 * time.Second)   // sleep for a second before sending next data
			}
		}
	}()

	wg.Wait()
}

func worker(ch chan int, wg *sync.WaitGroup, workerId int) {
	defer wg.Done()
	for msg := range ch {
		fmt.Printf("message: %d; workerId: %d\n", msg, workerId)
	}
}

/*
	Select
	Горутины и каналы являются фундаментальными строительными блоками конкурентного дизайна в Go. Последним элементом является оператор select. Он похож на простой переключатель, но решение принимается не на основе равенства, а на основе способности к взаимодействию.
	Следующий пример выдаёт один из трёх вариантов:
	Если канал ch1 готов (имеет значение), то выполняется первый случай.
	Если канал ch2 готов (имеет значение), то выполняется второй случай.
	Если ни один из каналов не готов, то выполняется случай по умолчанию.
*/

/*
	A Signal represents an operating system signal.
	The usual underlying implementation is operating system-dependent:
	on Unix it is syscall.Signal.
	type Signal interface {
		String() string
		Signal() // to distinguish from other Stringers
	}
*/
