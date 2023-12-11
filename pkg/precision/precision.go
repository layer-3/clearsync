package precision

import (
	"errors"
	"fmt"

	"github.com/shopspring/decimal"
)

func Convert(input decimal.Decimal, sigDigits, maxPrecision int32) decimal.Decimal {
	if input.Equal(decimal.Zero) {
		return input
	}

	integralDigits := int32(len(input.Coefficient().String()))
	precision := input.Exponent()
	if precision <= 0 {
		precision = -precision
	}

	adjustedPrecision := sigDigits - (integralDigits - precision)
	if adjustedPrecision < 0 {
		// If the number of integral digits is greater than significant digits,
		// round the number to a scale that maintains the significant digits in the integral part.
		multiplier := decimal.NewFromInt(10).Pow(decimal.NewFromInt32(-adjustedPrecision))
		return input.DivRound(multiplier, 0).Mul(multiplier)
	} else if adjustedPrecision > maxPrecision {
		adjustedPrecision = maxPrecision
	}

	return input.RoundBank(adjustedPrecision)
}

func Validate(input decimal.Decimal, sigDigits, maxPrecision int32) error {
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

	integralDigits := int32(len(input.Coefficient().String()))
	if integralDigits > sigDigits {
		return fmt.Errorf("input must not exceed max number of significant digits (allowed: %d; actual: %d)", sigDigits, int32(integralDigits))
	}

	return nil
}
