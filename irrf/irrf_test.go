package irrf

import (
	"testing"

	"github.com/shopspring/decimal"
)

var testCases = []struct {
	irrfBase float64
	irrfDue  float64
}{
	{0, 0},
	{500, 0},
	{1000, 0},
	{1500, 50},
	{2000, 100},
	{3000, 400},
	{10000, 1800},
}

var irrfByYear = map[string]IRRFRange{
	"2016": {
		{decimal.NewFromFloat(0), decimal.NewFromFloat(1000), decimal.NewFromFloat(0), decimal.NewFromFloat(0)},
		{decimal.NewFromFloat(1000), decimal.NewFromFloat(2000), decimal.NewFromFloat(.1), decimal.NewFromFloat(100)},
		{decimal.NewFromFloat(2000), decimal.NewFromFloat(0), decimal.NewFromFloat(.2), decimal.NewFromFloat(200)},
	},
}

func TestCalculate(t *testing.T) {
	irrf := NewIRRF(irrfByYear["2016"])
	for _, testCase := range testCases {
		irrfDue := irrf.Calculate(decimal.NewFromFloat(testCase.irrfBase))
		if irrfDue.Cmp(decimal.NewFromFloat(testCase.irrfDue)) != 0 {
			t.Errorf("Expected %v, got %v", testCase.irrfDue, irrfDue)
		}
	}
}
