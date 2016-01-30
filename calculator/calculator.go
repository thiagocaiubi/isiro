package calculator

import "github.com/shopspring/decimal"

type Calculator interface {
	Calculate(decimal.Decimal) decimal.Decimal
}
