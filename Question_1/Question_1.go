package main

import "fmt"

type Human struct {
	name string
	age  int
}

type Action struct {
	*Human
}

func main() {
	nprof := &Action{
		&Human{name: "John Doe", age: 20},
	}
	fmt.Printf("Result: age %d, name %s\n", nprof.age, nprof.name)
	nprof.changeName("John Newman")
	nprof.changeAge(30)
	fmt.Printf("New result: age %d, name %s", nprof.age, nprof.name)
}

func (a *Action) changeAge(change int) {
	a.age = change
}

func (a *Action) changeName(change string) {
	a.name = change
}
