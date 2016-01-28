package irpf

import (
	"testing"

	"github.com/shopspring/decimal"
)

var irpfTests = []struct {
	irpfBase float64
	irpfDue  float64
}{
	{0, 0},
	{500, 0},
	{1000, 0},
	{1500, 50},
	{2000, 100},
	{3000, 400},
	{10000, 1800},
}

var irpfByYear = map[string]IRPFRange{
	"2015": {
		{decimal.NewFromFloat(0), decimal.NewFromFloat(1000), decimal.NewFromFloat(0), decimal.NewFromFloat(0)},
		{decimal.NewFromFloat(1000), decimal.NewFromFloat(2000), decimal.NewFromFloat(.1), decimal.NewFromFloat(100)},
		{decimal.NewFromFloat(2000), decimal.NewFromFloat(0), decimal.NewFromFloat(.2), decimal.NewFromFloat(200)},
	},
}

func TestCalculate(t *testing.T) {
	irpf := NewIRPF(irpfByYear["2015"])
	for _, tt := range irpfTests {
		irpfDue := irpf.Calculate(decimal.NewFromFloat(tt.irpfBase))
		if irpfDue.Cmp(decimal.NewFromFloat(tt.irpfDue)) != 0 {
			t.Errorf("Expected IRPF contribution %v, but got %v", tt.irpfDue, irpfDue)
		}
	}
}
