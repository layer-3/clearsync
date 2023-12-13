package precision

import (
	"errors"
	"fmt"
	"math/big"

	"github.com/shopspring/decimal"
)

var ten = decimal.NewFromInt(10)

// ToSignificant truncates coefficient of any decimal.Decimal
// to a significant number of digits while preserving exponent.
// The input number is assumed to be non-negative.
func ToSignificant(input decimal.Decimal, sigDigits int32) decimal.Decimal {
	if input.Equal(decimal.Zero) {
		return input
	}

	coef := input.Coefficient()
	adjustedDigits := sigDigits - int32(len(coef.String()))
	if adjustedDigits >= 0 {
		return input
	}

	divisor := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(-adjustedDigits)), nil)
	coef.QuoRem(coef, divisor, new(big.Int))
	return decimal.NewFromBigInt(coef, input.Exponent()-adjustedDigits)
}

func Validate(input decimal.Decimal, maxPrecision int32) error {
	if input.IsNegative() {
		return errors.New("input must not be negative")
	}

	precision := int32(input.Exponent())
	if precision < 0 {
		precision = -precision
	}

	if precision > maxPrecision {
		return fmt.Errorf("input must not exceed max precision (allowed: %d; actual: %d)", maxPrecision, precision)
	}

	return nil
}
