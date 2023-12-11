package precision

import (
	"math"

	"github.com/shopspring/decimal"
)

func ToSignificant(input decimal.Decimal, sigDigits int32, maxScale int32) decimal.Decimal {
	if input.Equal(decimal.Zero) {
		return input
	}

	integralDigits := int32(len(input.Coefficient().String()))
	scale := input.Exponent()
	if scale > 0 {
		scale = 0
	} else {
		scale = -scale
	}

	adjustedScale := sigDigits - (integralDigits - scale)
	if adjustedScale < 0 {
		// If the number of integral digits is greater than significant digits,
		// round the number to a scale that maintains the significant digits in the integral part.
		roundedValue := input.RoundBank(0)
		multiplier := decimal.NewFromInt(int64(math.Pow10(int(-adjustedScale))))
		return roundedValue.DivRound(multiplier, 0).Mul(multiplier)
	} else if adjustedScale > maxScale {
		adjustedScale = maxScale
	}

	return input.RoundBank(adjustedScale)
}
