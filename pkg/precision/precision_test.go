package precision

import (
	"fmt"
	"math"
	"math/big"
	"strings"
	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/require"
)

const (
	sigDigits    int32 = 8
	maxPrecision int32 = 18
)

func TestToSignificant(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		input  decimal.Decimal
		expect decimal.Decimal
	}{
		{
			input:  decimal.RequireFromString("0.0"),
			expect: decimal.RequireFromString("0.0"),
		}, {
			input:  decimal.NewFromFloat(math.SmallestNonzeroFloat64),
			expect: decimal.NewFromFloat(math.SmallestNonzeroFloat64),
		}, {
			input:  decimal.NewFromFloat(math.MaxFloat64),
			expect: decimal.NewFromBigInt(big.NewInt(17976931), 301),
		}, {
			input:  decimal.RequireFromString("0.0000000000000004"),
			expect: decimal.RequireFromString("0.0000000000000004"),
		}, {
			input:  decimal.RequireFromString("0.0001"),
			expect: decimal.RequireFromString("0.0001"),
		}, {
			input:  decimal.RequireFromString("0.100103456123"),
			expect: decimal.RequireFromString("0.10010345"),
		}, {
			input:  decimal.RequireFromString("0.012345678999"),
			expect: decimal.RequireFromString("0.012345678"),
		}, {
			input:  decimal.RequireFromString("0.001234023499"),
			expect: decimal.RequireFromString("0.0012340234"),
		}, {
			input:  decimal.RequireFromString("0.012345678928"),
			expect: decimal.RequireFromString("0.012345678"),
		}, {
			input:  decimal.RequireFromString("0.100001023499"),
			expect: decimal.RequireFromString("0.10000102"),
		}, {
			input:  decimal.RequireFromString("0.012345"),
			expect: decimal.RequireFromString("0.012345"),
		}, {
			input:  decimal.RequireFromString("1.0000000234"),
			expect: decimal.RequireFromString("1.0"),
		}, {
			input:  decimal.RequireFromString("1.0002"),
			expect: decimal.RequireFromString("1.0002"),
		}, {
			input:  decimal.RequireFromString("2"),
			expect: decimal.RequireFromString("2"),
		}, {
			input:  decimal.RequireFromString("1000000060"),
			expect: decimal.RequireFromString("1000000000"),
		}, {
			input:  decimal.RequireFromString("12345"),
			expect: decimal.RequireFromString("12345"),
		}, {
			input:  decimal.RequireFromString("222.222222"),
			expect: decimal.RequireFromString("222.22222"),
		}, {
			input:  decimal.RequireFromString("2222022.2202"),
			expect: decimal.RequireFromString("2222022.2"),
		}, {
			input:  decimal.RequireFromString("218166.0002"),
			expect: decimal.RequireFromString("218166"),
		},
	}

	for _, test := range tests {
		tt := test
		t.Run(fmt.Sprintf("%s -> %s", tt.input, tt.expect), func(t *testing.T) {
			t.Parallel()

			actual := ToSignificant(tt.input, sigDigits)
			require.True(t, tt.expect.Equals(actual), "expected: %s; actual: %s", tt.expect, actual)
		})
	}
}

func BenchmarkToSignificant_DecimalWithLeadingZeros(b *testing.B) {
	d := decimal.RequireFromString("0.000123456")

	// Reset timer to exclude time taken by setup operations
	// before the actual benchmark begins
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		ToSignificant(d, sigDigits)
	}
}

func BenchmarkToSignificant_TruncateDecimals(b *testing.B) {
	d := decimal.RequireFromString("1.00000000234")

	// Reset timer to exclude time taken by setup operations
	// before the actual benchmark begins
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		ToSignificant(d, sigDigits)
	}
}

func BenchmarkToSignificant_ExactSignificantDigits(b *testing.B) {
	d := decimal.RequireFromString("12345678")
	if int32(len(d.String())) != sigDigits {
		panic("expected number of digits to be equal to number of significant digits")
	}

	// Reset timer to exclude time taken by setup operations
	// before the actual benchmark begins
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		ToSignificant(d, sigDigits)
	}
}

func BenchmarkToSignificant_IntegralPartSizeGreaterThanSignificantDigits(b *testing.B) {
	d := decimal.RequireFromString("1000000060")
	if integralPart := strings.Split(d.String(), ".")[0]; int32(len(integralPart)) <= sigDigits {
		panic("expected number of digits to be greater than number of significant digits")
	}

	// Reset timer to exclude time taken by setup operations
	// before the actual benchmark begins
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		ToSignificant(d, sigDigits)
	}
}

func TestValidate(t *testing.T) {
	t.Parallel()

	t.Run("Should return error on negative numbers", func(t *testing.T) {
		t.Parallel()

		input := decimal.RequireFromString("-12345")
		err := Validate(input, maxPrecision)
		require.Error(t, err)
	})

	t.Run("Should return error if precision is greater than allowed", func(t *testing.T) {
		t.Parallel()

		input := decimal.RequireFromString("0.00000000000000012345678")
		precision := int32(math.Abs(float64(input.Exponent())))
		require.Greater(t, precision, maxPrecision)

		err := Validate(input, maxPrecision)
		require.Error(t, err)
	})

	t.Run("Successful test", func(t *testing.T) {
		t.Parallel()

		input := decimal.RequireFromString("1.2345")
		digits := int32(len(input.Coefficient().String()))
		precision := int32(math.Abs(float64(input.Exponent())))
		require.LessOrEqual(t, digits, sigDigits)
		require.LessOrEqual(t, precision, maxPrecision)

		err := Validate(input, maxPrecision)
		require.NoError(t, err)
	})
}

func BenchmarkValidate_SuccessfulCase(b *testing.B) {
	d := decimal.RequireFromString("1.2345")
	digits := int32(len(d.Coefficient().String()))
	precision := int32(math.Abs(float64(d.Exponent())))
	require.LessOrEqual(b, digits, sigDigits)
	require.LessOrEqual(b, precision, maxPrecision)

	// Reset timer to exclude time taken by setup operations
	// before the actual benchmark begins
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = Validate(d, maxPrecision)
	}
}

func BenchmarkValidate_UnsuccessfulCase(b *testing.B) {
	d := decimal.RequireFromString("1.234523452345234523452345")

	// Reset timer to exclude time taken by setup operations
	// before the actual benchmark begins
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = Validate(d, maxPrecision)
	}
}
