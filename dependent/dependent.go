package dependent

import "github.com/shopspring/decimal"

type Dependent struct {
	deduction decimal.Decimal
}

func NewDependent(deduction decimal.Decimal) Dependent {
	return Dependent{
		deduction: deduction,
	}
}

func (dep *Dependent) Calculate(quantity decimal.Decimal) decimal.Decimal {
	return quantity.Mul(dep.deduction)
}
