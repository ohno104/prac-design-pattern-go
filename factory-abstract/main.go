package main

import "fmt"

type Person interface {
	Greet()
}

type person struct {
	name string
	age  int
}

func (p person) Greet() {
	fmt.Println("Hi! My name is", p.name)
}

// 返回interface, 在不公開內部實現的情況下讓調用者使用提供的func
func NewPerson(name string, age int) Person {
	return person{
		name: name,
		age:  age,
	}
}
