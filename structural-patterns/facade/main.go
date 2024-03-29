package main

// 外觀模式
// 提供一個簡單的接口, 把子系統的複雜性隱藏起來, 容易使用

import (
	"fmt"

	"github.com/hono104/prac-design-pattern-go/structural-patterns/facade/account"
	"github.com/hono104/prac-design-pattern-go/structural-patterns/facade/risk"
	"github.com/hono104/prac-design-pattern-go/structural-patterns/facade/sales"
)

type Facade struct {
	sales   sales.SalesSystem
	risk    risk.RiskSystem
	account account.AccountSystem
}

func NewFacade() *Facade {
	return &Facade{
		*sales.NewSalesSystem(),
		*risk.NewRiskSystem(),
		*account.NewAccountingSystem(),
	}
}

func (f *Facade) Loan() {
	f.sales.Apply()
	f.risk.Assess()
	f.account.Trans()
	fmt.Println("loan success")
}

func main() {
	c := NewFacade()
	c.Loan()
}
