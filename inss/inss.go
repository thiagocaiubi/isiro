package inss

import (
	"github.com/shopspring/decimal"
)

type INSSRange []struct {
	Lower   decimal.Decimal `json:"lower"`
	Upper   decimal.Decimal `json:"upper"`
	Aliquot decimal.Decimal `json:"aliquot"`
}

type INSSByYear map[string]INSSRange

type INSS struct {
	inssRange INSSRange
}

func NewINSS(inssRange INSSRange) INSS {
	return INSS{
		inssRange: inssRange,
	}
}

func (inss *INSS) Calculate(grossSalary decimal.Decimal) decimal.Decimal {
	for _, inssInterval := range inss.inssRange {
		if grossSalary.Cmp(inssInterval.Lower) >= 0 && grossSalary.Cmp(inssInterval.Upper) <= 0 {
			return grossSalary.Mul(inssInterval.Aliquot)
		}
	}

	lastINSSInterval := inss.inssRange[len(inss.inssRange)-1]

	return lastINSSInterval.Upper.Mul(lastINSSInterval.Aliquot)
}
