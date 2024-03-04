package financial

import "fmt"

type FinancialSystem struct{}

func NewFinancialSystem() *FinancialSystem {
	return &FinancialSystem{}
}

func (f *FinancialSystem) Appropriation() {
	fmt.Println("撥款")
}
