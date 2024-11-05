package main

import (
	"fmt"
	"strings"
)

/*
	Разработать программу, которая проверяет,
	что все символы в строке уникальные
	(true — если уникальные, false etc).
	Функция проверки должна быть регистронезависимой.

	Например:
	abcd — true
	abCdefAaf — false
	aabcd — false
*/

func main() {
	str := "abCdefAf" // Исходная строка для проверки

	// Приводим строку к нижнему регистру для регистронезависимости
	lowerStr := strings.ToLower(str)

	// Используем map для отслеживания уникальности символов
	charMap := make(map[rune]bool)

	// Итерируем по строке и проверяем каждый символ
	for _, char := range lowerStr {
		if charMap[char] {
			fmt.Println("false") // Найден повторяющийся символ
			return
		}
		charMap[char] = true
	}

	fmt.Println("true")
}
