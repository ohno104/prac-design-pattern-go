package main

import (
	"design-pattern/facade/business"
	"design-pattern/facade/financial"
	"design-pattern/facade/security"
)

type Facade struct {
	b business.BusinessSystem
	s security.SecuritySystem
	f financial.FinancialSystem
}

func NewFacaade() *Facade {
	return &Facade{
		b: *business.NewBusinessSystem(),
		s: *security.NewSecuritySystem(),
		f: *financial.NewFinancialSystem(),
	}
}

func (c *Facade) Loan() {
	c.b.Submit()
	c.s.RiskAssessment()
	c.f.Appropriation()
}

func main() {
	c := NewFacaade()
	c.Loan()
}
