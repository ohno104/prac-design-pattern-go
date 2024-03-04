package business

import "fmt"

type BusinessSystem struct{}

func NewBusinessSystem() *BusinessSystem {
	return &BusinessSystem{}
}

func (b *BusinessSystem) Submit() {
	fmt.Println("申請貸款")
}
