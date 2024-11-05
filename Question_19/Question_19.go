package main

import "fmt"

/*
	Разработать программу, которая переворачивает
	подаваемую на ход строку (например: «главрыба — абырвалг»).
	Символы могут быть unicode.
*/

func main() {
	// Исходная строка, преобразованная в срез rune для поддержки Unicode
	str := "главрыба"
	runes := []rune(str)

	// Создаем новый срез rune той же длины для хранения перевернутой строки
	reversedRunes := make([]rune, len(runes))

	// Переворачиваем строку
	for i, r := range runes {
		reversedRunes[len(runes)-1-i] = r
	}

	// Преобразуем срез rune обратно в строку и выводим
	fmt.Println(string(reversedRunes))
}
