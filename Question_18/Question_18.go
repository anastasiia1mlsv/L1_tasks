package main

import (
	"fmt"
	"sync"
	"time"
)

/*
	Реализовать структуру-счетчик,
	которая будет инкрементироваться в конкурентной среде.
	По завершению программа должна выводить итоговое значение счетчика.
*/

// Inc — структура для безопасного инкрементирования в конкурентной среде
type Inc struct {
	count int
	mu    sync.Mutex // Мьютекс для синхронизации доступа к count
}

// Increment увеличивает count на 1
func (i *Inc) Increment() {
	i.mu.Lock() // Блокируем доступ к count
	i.count++   // Инкрементируем count
	fmt.Println("new count:", i.count)
	i.mu.Unlock() // Разблокируем доступ
}

func main() {
	var counter Inc // Создаем счетчик

	// Запускаем горутину для инкрементирования счетчика
	go func() {
		for {
			counter.Increment()         // Инкрементируем счетчик
			time.Sleep(2 * time.Second) // Ждем 2 секунды
		}
	}()

	time.Sleep(10 * time.Second)               // Ждем 10 секунд, пока счетчик инкрементируется
	fmt.Println("Final count:", counter.count) // Выводим итоговое значение счетчика
}
