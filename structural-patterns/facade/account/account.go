package account

import "fmt"

type AccountSystem struct{}

func NewAccountingSystem() *AccountSystem {
	return &AccountSystem{}
}

func (a *AccountSystem) Trans() {
	fmt.Println("trans")
}
