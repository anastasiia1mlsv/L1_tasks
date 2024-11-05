package main

import (
	"fmt"
	"reflect"
)

/*
	Разработать программу, которая в рантайме способна
	определить тип переменной: int, string, bool, channel
	из переменной типа interface{}.
*/

func main() {

	var inter interface{}

	inter = 10
	fmt.Println(reflect.TypeOf(inter))

	inter = "string"
	fmt.Println(reflect.TypeOf(inter))
	//Присваиваем inter значение типа chan int (канал int)
	inter = make(chan int)
	fmt.Println(reflect.TypeOf(inter))

	inter = true
	fmt.Println(reflect.TypeOf(inter))
}
