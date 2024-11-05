package main

import (
	"fmt"
	"sync"
	"time"
)

/*
	Реализовать конкурентную запись данных в map
*/

func main() {
	mappa := map[int]int{}

	wg := sync.WaitGroup{}
	mutex := sync.Mutex{} // Мьютекс — это блокировка взаимного исключения.
	wg.Add(1)

	go func() {
		for i := 0; i < 5; i++ {
			mutex.Lock()     // Блокируем доступ к map
			mappa[i] = i * 2 // Записываем данные
			mutex.Unlock()   // Разблокируем доступ

			time.Sleep(1 * time.Second)
		}
		defer wg.Done()
	}()

	wg.Wait()
	fmt.Println(mappa)
}
