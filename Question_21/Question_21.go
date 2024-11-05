package main

import "fmt"

/*
 Реализовать паттерн «адаптер» на любом примере.
*/

type newPrintS struct{}

type basePrintS struct{}

type adapter struct {
	adapted basePrintS
}

func (a *adapter) AdaptedPrint() string {
	return a.adapted.BasePrint("base hello")
}

func (b *basePrintS) BasePrint(str string) string {
	return "base print " + str
}

func (s *newPrintS) NewPrint() string {
	return "new print"
}

func main() {
	str := &basePrintS{}
	newStr := &newPrintS{}
	adapt := &adapter{*str}
	fmt.Println(newStr.NewPrint())
	fmt.Println(adapt.AdaptedPrint())
}

/*
	Паттерн "Адаптер" в программировании используется
	для преобразования интерфейса одного класса в интерфейс другого.
	Это особенно полезно, когда невозможно изменить исходный класс,
	но необходимо интегрировать его функциональность с другими классами.
*/
