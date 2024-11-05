package main

import (
	"fmt"
	"strings"
)

/*
	К каким негативным последствиям может привести
	данный фрагмент кода, и как это исправить?
	Приведите корректный пример реализации.

	var justString string
	func someFunc() {
	  v := createHugeString(1 << 10) //  не нужно делать подстроку во избежание утечки памяти
	}

	func main() {
	  someFunc()
	}
*/

func createHugeString(s int) string {
	//создаем большую строку
	return strings.Repeat("s", s)
}

func someFunc() string {
	v := createHugeString(1 << 10) //1024
	// Создаем новую строку с копией нужных данных
	justString := make([]byte, 100)
	copy(justString, v[:100])
	return string(justString)
}

func main() {
	// justString в main теперь ссылается на отдельную строку,
	// а не на подстроку исходной большой строки,
	// что позволяет избежать удержания исходной большой строки в памяти.
	justString := someFunc()
	fmt.Println(justString)
}
