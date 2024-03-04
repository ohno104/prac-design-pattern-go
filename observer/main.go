package main

import (
	"fmt"
)

type Item struct {
	//customerList []Customer 各類型用戶
	observerList []Observer
	name         string
	inStock      bool
}

type Observer interface {
	getId() string
	update(string)
}

func NewItem(name string) *Item {
	return &Item{
		name: name,
	}
}

func (i *Item) updateAvailability() {
	fmt.Println("Item", i.name, "is now in stock")
	i.inStock = true
	//
	i.notifyAll()
}

func (i *Item) register(o Observer) {
	i.observerList = append(i.observerList, o)
}

func (i *Item) notifyAll() {
	for _, o := range i.observerList {
		o.update(i.name)
	}
}

type Customer struct {
	id string
}

func (c *Customer) getId() string {
	return c.id
}

func (c *Customer) update(itemName string) {
	fmt.Printf("sending email to customer %s for item %s \n", c.id, itemName)
}

func main() {
	shirtItem := NewItem("Nike Shirt")

	observerFirst := &Customer{id: "abc@gmail.com"}
	observerSecond := &Customer{id: "xyz@gmail.com"}

	shirtItem.register(observerFirst)
	shirtItem.register(observerSecond)

	shirtItem.updateAvailability()
}
