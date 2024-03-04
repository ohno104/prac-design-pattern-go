package main

import "fmt"

type Vehicle interface {
	Drive()
}

type Car struct{}

func (c *Car) Drive() {
	fmt.Println("car is being driven")
}

type Driver struct {
	Age int
}

type CarProxy struct {
	vehicle Vehicle
	driver  *Driver
}

func NewCarProxy(driver *Driver) *CarProxy {
	return &CarProxy{&Car{}, driver}
}

func (c *CarProxy) Drive() {
	if c.driver.Age >= 18 {
		c.vehicle.Drive()
	} else {
		fmt.Println("driver too young")
	}
}
