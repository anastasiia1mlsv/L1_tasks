package main

import "fmt"

/*
	Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.
*/

func main() {
	// Исходный массив строк
	strings := []string{"cat", "cat", "dog", "cat", "tree"}
	// Map для хранения уникальных строк
	uniqueStrings := make(map[string]bool)

	// Добавление каждой строки в map для создания множества
	for _, str := range strings {
		uniqueStrings[str] = true
	}

	// Создание среза для хранения уникальных строк из map
	uniqueArray := make([]string, 0, len(uniqueStrings))
	for str := range uniqueStrings {
		uniqueArray = append(uniqueArray, str)
	}

	fmt.Println(uniqueArray)
}
