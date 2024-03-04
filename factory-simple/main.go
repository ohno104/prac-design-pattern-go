package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) Greet() {
	fmt.Println("Hi! My name is", p.Name)
}

// 避免直接使用p := &Person{}的方式, 確保創建的實例具有需要的參數
func NewPerson(name string, age int) *Person {
	return &Person{
		Name: name,
		Age:  age,
	}
}
