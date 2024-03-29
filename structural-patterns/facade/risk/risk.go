package risk

import "fmt"

type RiskSystem struct{}

func NewRiskSystem() *RiskSystem {
	return &RiskSystem{}
}

func (r *RiskSystem) Assess() {
	fmt.Println("assess")
}
