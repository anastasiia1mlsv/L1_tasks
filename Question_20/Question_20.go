package main

import (
	"fmt"
	"strings"
)

/*
	Разработать программу, которая переворачивает слова в строке.
	Пример: «snow dog sun — sun dog snow».
*/

func main() {

	// создаем новую строку и логаем ее
	stringSentence := "snow dog sun"
	fmt.Println(stringSentence)
	// чистим пробелы
	trimmed := strings.Split(stringSentence, " ")
	newString := ""

	//проходимся по нашему списку слов в обратном порядке и кидаем каждое новое слово в новое слово, добавляя пробел
	for i := len(trimmed) - 1; i >= 0; i-- {
		newString += trimmed[i] + " "
	}

	//делаем трим чтобы убрать крайние пробелы, присваиваем исходной строке новое значние
	stringSentence = strings.TrimRight(newString, " ")
	fmt.Println(stringSentence)
}
