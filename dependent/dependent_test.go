package dependent

import (
	"testing"

	"github.com/shopspring/decimal"
)

var testCases = []struct {
	quantity decimal.Decimal
	expected decimal.Decimal
}{
	{decimal.NewFromFloat(0), decimal.NewFromFloat(0)},
	{decimal.NewFromFloat(1), decimal.NewFromFloat(50)},
	{decimal.NewFromFloat(3), decimal.NewFromFloat(150)},
	{decimal.NewFromFloat(10), decimal.NewFromFloat(500)},
}

func TestCalculate(t *testing.T) {
	dep := NewDependent(decimal.NewFromFloat(50))
	for _, testCase := range testCases {
		depDeduction := dep.Calculate(testCase.quantity)
		if depDeduction.Cmp(testCase.expected) != 0 {
			t.Errorf("Expected %v, but got %v", testCase.expected, depDeduction)
		}
	}
}
