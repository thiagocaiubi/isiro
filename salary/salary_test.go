package salary

import (
	"testing"

	"github.com/shopspring/decimal"
)

// Magic numbers bellow this line are better explained at:
// https://docs.google.com/spreadsheets/d/1K0jD3Fs9-z_yJ_eOloBXVDK85BIY_5qs2veEUs5wZKU/edit#gid=0
// This spreadsheet describes inputs, outputs and the formula to get there.
var rebatePercentage = decimal.NewFromFloat(.1)

type fakeCalculator struct{}

func (f fakeCalculator) Calculate(d decimal.Decimal) decimal.Decimal {
	return d.Mul(rebatePercentage)
}

var testCases = []struct {
	// input
	grossSalary   decimal.Decimal
	dependentsQty decimal.Decimal
	// output
	inss       decimal.Decimal
	dependents decimal.Decimal
	irrf       decimal.Decimal
	netSalary  decimal.Decimal
}{
	{
		decimal.NewFromFloat(10000),
		decimal.NewFromFloat(1),
		decimal.NewFromFloat(1000),
		decimal.NewFromFloat(.1),
		decimal.NewFromFloat(899.99),
		decimal.NewFromFloat(8100.01),
	},
	{
		decimal.NewFromFloat(5000),
		decimal.NewFromFloat(0),
		decimal.NewFromFloat(500),
		decimal.NewFromFloat(0),
		decimal.NewFromFloat(450),
		decimal.NewFromFloat(4050),
	},
	{
		decimal.NewFromFloat(15000),
		decimal.NewFromFloat(3),
		decimal.NewFromFloat(1500),
		decimal.NewFromFloat(.3),
		decimal.NewFromFloat(1349.97),
		decimal.NewFromFloat(12150.03),
	},
}

func TestCalculate(t *testing.T) {
	for _, testCase := range testCases {
		sc := NewSalaryCalculator(fakeCalculator{}, fakeCalculator{}, fakeCalculator{})
		sal := sc.Calculate(testCase.grossSalary, testCase.dependentsQty)

		if sal.GrossSalary.Cmp(testCase.grossSalary) != 0 {
			t.Errorf("Expected %v, got %v", testCase.grossSalary, sal.GrossSalary)
		}

		if sal.INSS.Cmp(testCase.inss) != 0 {
			t.Errorf("Expected %v, got %v", testCase.inss, sal.INSS)
		}

		if sal.Dependents.Cmp(testCase.dependents) != 0 {
			t.Errorf("Expected %v, got %v", testCase.dependents, sal.Dependents)
		}

		if sal.IRRF.Cmp(testCase.irrf) != 0 {
			t.Errorf("Expected %v, got %v", testCase.irrf, sal.IRRF)
		}

		if sal.NetSalary.Cmp(testCase.netSalary) != 0 {
			t.Errorf("Expected %v, got %v", testCase.netSalary, sal.NetSalary)
		}
	}
}
