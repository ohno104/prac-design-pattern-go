package main

import "fmt"

type Workflow interface {
	start()
	execute()
	end()
}

type BaseWorkflow struct{}

func (b *BaseWorkflow) start() {
	fmt.Println("工作流程開始")
}

func (b *BaseWorkflow) end() {
	fmt.Println("工作流程結束")
}

func (b *BaseWorkflow) execute() {
}

func RunWorkflow(w Workflow) {
	w.start()
	w.execute()
	w.end()
}

type A struct {
	BaseWorkflow
}

func (*A) execute() {
	fmt.Println("工作流程A執行中")
}

type B struct {
	BaseWorkflow
}

func (*B) execute() {
	fmt.Println("工作流程B執行中")
}

func main() {
	RunWorkflow(&A{})
	RunWorkflow(&B{})
}
