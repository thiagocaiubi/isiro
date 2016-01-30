package irrf

import "github.com/shopspring/decimal"

type IRRFRange []struct {
	Lower     decimal.Decimal `json:"lower"`
	Upper     decimal.Decimal `json:"upper"`
	Aliquot   decimal.Decimal `json:"aliquot"`
	Deduction decimal.Decimal `json:"deduction"`
}

type IRRFByYear map[string]IRRFRange

type IRRF struct {
	irrfRange IRRFRange
}

func NewIRRF(irrfRange IRRFRange) IRRF {
	return IRRF{
		irrfRange: irrfRange,
	}
}

func (irrf *IRRF) Calculate(grossSalary decimal.Decimal) decimal.Decimal {
	for _, irrfInterval := range irrf.irrfRange {
		if grossSalary.Cmp(irrfInterval.Lower) >= 0 && grossSalary.Cmp(irrfInterval.Upper) <= 0 {
			return grossSalary.Mul(irrfInterval.Aliquot).Sub(irrfInterval.Deduction)
		}
	}

	lastIRRFInterval := irrf.irrfRange[len(irrf.irrfRange)-1]

	return grossSalary.Mul(lastIRRFInterval.Aliquot).Sub(lastIRRFInterval.Deduction)
}
