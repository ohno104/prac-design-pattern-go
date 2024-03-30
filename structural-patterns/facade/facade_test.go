package facade

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoan(t *testing.T) {
	c := NewFacade()
	assert.Equal(t, c.Loan(), true)
}
