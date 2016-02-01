package salary

import (
	"testing"

	"github.com/shopspring/decimal"
)

var tenPercent = decimal.NewFromFloat(.1)

type fakeCalculator struct{}

func (f fakeCalculator) Calculate(d decimal.Decimal) decimal.Decimal {
	return d.Mul(tenPercent)
}

var testCases = []struct {
	grossSalary  decimal.Decimal
	depedentsQty decimal.Decimal
}{
	{
		decimal.NewFromFloat(10000),
		decimal.NewFromFloat(1),
	},
	{
		decimal.NewFromFloat(5000),
		decimal.NewFromFloat(3),
	},
	{
		decimal.NewFromFloat(3000),
		decimal.NewFromFloat(0),
	},
}

func TestCalculate(t *testing.T) {
	for _, testCase := range testCases {
		sc := NewSalaryCalculator(fakeCalculator{}, fakeCalculator{}, fakeCalculator{})
		sal := sc.Calculate(testCase.grossSalary, testCase.depedentsQty)

		if sal.GrossSalary.Cmp(testCase.grossSalary) != 0 {
			t.Errorf("Expected %v, got %v", testCase.grossSalary, sal.GrossSalary)
		}

		inss := testCase.grossSalary.Mul(tenPercent)
		if sal.INSS.Cmp(inss) != 0 {
			t.Errorf("Expected %v, got %v", inss, sal.INSS)
		}

		dependents := testCase.depedentsQty.Mul(tenPercent)
		if sal.Dependents.Cmp(dependents) != 0 {
			t.Errorf("Expected %v, got %v", dependents, sal.Dependents)
		}

		irrf := (testCase.grossSalary.Sub(inss).Sub(dependents)).Mul(tenPercent)
		if sal.IRRF.Cmp(irrf) != 0 {
			t.Errorf("Expected %v, got %v", irrf, sal.IRRF)
		}

		netSalary := testCase.grossSalary.Sub(inss).Sub(irrf)
		if sal.NetSalary.Cmp(netSalary) != 0 {
			t.Errorf("Expected %v, got %v", netSalary, sal.NetSalary)
		}
	}
}
