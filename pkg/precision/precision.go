// Package precision provides functionality for handling the precision of asset prices
// and quantities in accordance with the YIP-0001 - Asset Price Precision specification.
//
// This package includes functions to ensure that asset prices and quantities are represented
// with the appropriate level of significant digits and precision, as required in the trading
// industry. It is designed to align with the global standards for price-point (PIP) representation
// in FOREX and other financial markets, facilitating configless deployments and network-wide
// harmony of configuration.
//
// Key Features:
//
//   - ToSignificant: Truncates a decimal number to a specified number of significant digits, adhering
//     to the YIP-0001 specification for price representation. This function is essential for maintaining
//     the standard price precision in trading scenarios.
//
//   - Validate: Ensures that the given decimal number does not exceed the maximum allowed precision for
//     prices (18 significant digits) and quantities (8 significant digits). It plays a crucial role in
//     validating input data for trading operations.
//
// This package is an integral part of the Yellow Network's trading infrastructure, ensuring accuracy,
// consistency, and compliance with established financial market standards.
package precision

import (
	"errors"
	"fmt"
	"math/big"

	"github.com/shopspring/decimal"
)

// ToSignificant truncates coefficient of any decimal.Decimal
// to a significant number of digits while preserving exponent.
// The input number is assumed to be non-negative.
// If the input number has fewer significant digits than requested, it is returned unchanged.
//
// Parameters:
//   - input: The number to be truncated. The number must be non-negative.
//   - sigDigits: The number of significant digits to retain in the number.
//     Significant digits are counted from the most significant digit (first non-zero digit) to the
//     specified count. Trailing zeros after the decimal point are not considered significant.
//
// Returns:
//   - A decimal.Decimal representing the input number truncated to the specified number of
//     significant digits. If the input is zero, it returns zero without modification.
//
// Examples:
//
//   - Truncating a number to a specific number of significant digits:
//     result := ToSignificant(decimal.NewFromFloat(123.456), 4) // Returns 123.4
//
//   - Processing a number with fewer significant digits than requested:
//     result := ToSignificant(decimal.NewFromFloat(0.0123), 5) // Returns 0.0123 (unchanged)
//
// Note:
// - This function does not round the last significant digit; it merely truncates the number.
// - Negative inputs are not handled by this function and may lead to unintended results.
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

// Validate checks if the given number input adheres to
// precision rules specified in the YIP-0001.
// The function ensures that the input is non-negative and that
// its precision does not exceed a limit.
// Precision in this context is defined as the number of digits
// after the decimal point for non-integer numbers.
//
// Parameters:
//   - input: The number to be validated.
//   - maxPrecision: The maximum allowed precision (number of digits
//     after the decimal point). For example, if maxPrecision is 2, then 1.234 would be
//     considered invalid, since allowed exponent is 10^-2
//     and the number is 10^-3, while 1.23 would be valid.
//
// Returns:
//   - An error if the input is negative or if its precision exceeds maxPrecision. If the input
//     is valid according to the specified constraints, the function returns nil.
//
// Example:
//
//   - Validating a number with a precision that does not exceed the limit:
//     err := Validate(decimal.NewFromFloat(1.234), 3) // Returns nil, as the precision is within the limit.
//
//   - Validating a negative number or a number with excessive precision:
//     err := Validate(decimal.NewFromFloat(-1.234), 3) // Returns an error, as the input is negative.
//     err := Validate(decimal.NewFromFloat(1.2345), 3) // Returns an error, as the precision exceeds the limit.
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
