package irpf

import "github.com/shopspring/decimal"

type IRPFRange []struct {
	Lower     decimal.Decimal `json:"lower"`
	Upper     decimal.Decimal `json:"upper"`
	Aliquot   decimal.Decimal `json:"aliquot"`
	Deduction decimal.Decimal `json:"deduction"`
}

type IRPFByYear map[string]IRPFRange

type IRPF struct {
	irpfRange IRPFRange
}

func NewIRPF(irpfRange IRPFRange) IRPF {
	return IRPF{
		irpfRange: irpfRange,
	}
}

func (irpf *IRPF) Calculate(grossSalary decimal.Decimal) decimal.Decimal {
	for _, irpfInterval := range irpf.irpfRange {
		if grossSalary.Cmp(irpfInterval.Lower) >= 0 && grossSalary.Cmp(irpfInterval.Upper) <= 0 {
			return grossSalary.Mul(irpfInterval.Aliquot).Sub(irpfInterval.Deduction)
		}
	}

	lastIRPFInterval := irpf.irpfRange[len(irpf.irpfRange)-1]

	return grossSalary.Mul(lastIRPFInterval.Aliquot).Sub(lastIRPFInterval.Deduction)
}
