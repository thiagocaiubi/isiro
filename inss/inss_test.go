package inss

import (
	"testing"

	"github.com/shopspring/decimal"
)

var testCases = []struct {
	grossSalary float64
	inssDue     float64
}{
	{0., 0.},
	{500., 50.},
	{1000., 100.},
	{1500., 300.},
	{2000., 400.},
	{2500., 750.},
	{3000., 900.},
	{3500., 900.},
	{10000., 900.},
}

var inssByYear = map[string]INSSRange{
	"2016": {
		{decimal.NewFromFloat(0.), decimal.NewFromFloat(1000.), decimal.NewFromFloat(.1)},
		{decimal.NewFromFloat(1000.), decimal.NewFromFloat(2000.), decimal.NewFromFloat(.2)},
		{decimal.NewFromFloat(2000.), decimal.NewFromFloat(3000), decimal.NewFromFloat(.3)},
	},
}

func TestCalculate(t *testing.T) {
	inss := NewINSS(inssByYear["2016"])
	for _, testCase := range testCases {
		inssDue := inss.Calculate(decimal.NewFromFloat(testCase.grossSalary))
		if inssDue.Cmp(decimal.NewFromFloat(testCase.inssDue)) != 0 {
			t.Errorf("Expected %v, got %v", testCase.inssDue, inssDue)
		}
	}
}
