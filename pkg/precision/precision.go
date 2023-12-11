package precision

import (
	"github.com/shopspring/decimal"
)

func ToSignificant(input decimal.Decimal, sigDigits int32, maxScale int32) decimal.Decimal {
	if input.Equal(decimal.Zero) {
		return input
	}

	integralDigits := int32(len(input.Coefficient().String()))
	scale := input.Exponent()
	if scale <= 0 {
		scale = -scale
	}

	adjustedScale := sigDigits - (integralDigits - scale)
	if adjustedScale < 0 {
		// If the number of integral digits is greater than significant digits,
		// round the number to a scale that maintains the significant digits in the integral part.
		multiplier := decimal.NewFromInt(10).Pow(decimal.NewFromInt32(-adjustedScale))
		return input.DivRound(multiplier, 0).Mul(multiplier)
	} else if adjustedScale > maxScale {
		adjustedScale = maxScale
	}

	return input.RoundBank(adjustedScale)
}

func IsValid(input decimal.Decimal, sigDigits int32, maxScale int32) bool {
	if input.IsNegative() {
		return false
	}
	result := ToSignificant(input, sigDigits, maxScale)
	return result.Equal(input)
}
