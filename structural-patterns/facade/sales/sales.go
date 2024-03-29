package sales

import "fmt"

type SalesSystem struct{}

func NewSalesSystem() *SalesSystem {
	return &SalesSystem{}
}

func (b *SalesSystem) Apply() {
	fmt.Println("apply")
}
