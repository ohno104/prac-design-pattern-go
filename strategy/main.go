package main

import "fmt"

type IStrategy interface {
	do(int, int) int
}

// 策略A
type add struct{}

func (*add) do(a, b int) int {
	return a + b
}

// 策略B
type subtract struct{}

func (*subtract) do(a, b int) int {
	return a - b
}

// 策略執行者
type Operator struct {
	strategy IStrategy
}

func (operator *Operator) setStrategy(strategy IStrategy) {
	operator.strategy = strategy
}

func (operator *Operator) calculate(a, b int) int {
	return operator.strategy.do(a, b)
}

func main() {
	operator := Operator{}

	operator.setStrategy(&add{})
	result := operator.calculate(1, 2)
	fmt.Println("add:", result)

	operator.setStrategy(&subtract{})
	result = operator.calculate(2, 1)
	fmt.Println("reduce:", result)
}
