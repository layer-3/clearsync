package precision

import (
	"github.com/shopspring/decimal"
)

func ToSignificant(number decimal.Decimal, digits int32, maxDecimalPoints int) decimal.Decimal {
	if number.IsZero() {
		return number
	}

	bytes := []byte(number.String())
	pointIndex := -1
	zerosAfterPoint := -1

	for i, r := range bytes {
		if r == '0' {
			zerosAfterPoint++
			continue
		}
		if r == '.' {
			pointIndex = i
		}
		if r != '.' && pointIndex != -1 {
			break
		}
	}

	exponent := len(bytes)
	if bytes[0] == '0' {
		if zerosAfterPoint >= maxDecimalPoints {
			return decimal.Zero
		}
		exponent = -zerosAfterPoint
	} else {
		if pointIndex != -1 {
			exponent = pointIndex
		}
	}

	c := number.Coefficient()
	// TODO: choose the rounding rule: round, round floor, bank round etc.
	coefficient := decimal.NewFromBigInt(c, int32(-len(c.String()))).Round(digits).Coefficient()

	return decimal.NewFromBigInt(coefficient, int32(exponent)-digits)
}
