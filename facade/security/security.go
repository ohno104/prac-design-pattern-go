package security

import "fmt"

type SecuritySystem struct{}

func NewSecuritySystem() *SecuritySystem {
	return &SecuritySystem{}
}

func (s *SecuritySystem) RiskAssessment() {
	fmt.Println("風險評估")
}
