package salary

import (
	"github.com/shopspring/decimal"
	"github.com/thiagocaiubi/isiro/calculator"
)

type Salary struct {
	GrossSalary decimal.Decimal
	INSS        decimal.Decimal
	Dependents  decimal.Decimal
	IRRF        decimal.Decimal
	NetSalary   decimal.Decimal
}

type salaryCalculator struct {
	inss      calculator.Calculator
	depedents calculator.Calculator
	irrf      calculator.Calculator
}

func NewSalaryCalculator(inss, depedents, irrf calculator.Calculator) salaryCalculator {
	return salaryCalculator{
		inss:      inss,
		depedents: depedents,
		irrf:      irrf,
	}
}

func (sc *salaryCalculator) Calculate(grossSalary, dependentsQty decimal.Decimal) Salary {
	inss := sc.inss.Calculate(grossSalary)
	dependents := sc.depedents.Calculate(dependentsQty)
	irrfBase := grossSalary.Sub(inss).Sub(dependents)
	irrf := sc.irrf.Calculate(irrfBase)
	netSalary := grossSalary.Sub(inss).Sub(irrf)
	return Salary{
		GrossSalary: grossSalary,
		INSS:        inss,
		Dependents:  dependents,
		IRRF:        irrf,
		NetSalary:   netSalary,
	}
}
