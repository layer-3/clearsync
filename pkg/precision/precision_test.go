package precision

import (
	"fmt"
	"math"
	"strings"
	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/require"
)

const (
	sigDigits    int32 = 8
	maxPrecision int32 = 18
)

func TestConvert(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		input  decimal.Decimal
		expect decimal.Decimal
	}{
		{
			input:  newFromString("0.0"),
			expect: newFromString("0.0"),
		}, {
			input:  decimal.NewFromFloat(math.SmallestNonzeroFloat64),
			expect: newFromString("0.0"),
		}, {
			input:  newFromString("0.0000000000000004"),
			expect: newFromString("0.0000000000000004"),
		}, {
			input:  newFromString("0.0001"),
			expect: newFromString("0.0001"),
		}, {
			input:  newFromString("0.100103456123"),
			expect: newFromString("0.10010346"),
		}, {
			input:  newFromString("0.012345678999"),
			expect: newFromString("0.012345679"),
		}, {
			input:  newFromString("0.001234023499"),
			expect: newFromString("0.0012340235"),
		}, {
			input:  newFromString("0.012345678928"),
			expect: newFromString("0.012345679"),
		}, {
			input:  newFromString("0.100001023499"),
			expect: newFromString("0.10000102"),
		}, {
			input:  newFromString("0.012345"),
			expect: newFromString("0.012345"),
		}, {
			input:  newFromString("1.0000000234"),
			expect: newFromString("1.0"),
		}, {
			input:  newFromString("1.0002"),
			expect: newFromString("1.0002"),
		}, {
			input:  newFromString("2"),
			expect: newFromString("2"),
		}, {
			input:  newFromString("1000000060"),
			expect: newFromString("1000000100"),
		}, {
			input:  newFromString("12345"),
			expect: newFromString("12345"),
		}, {
			input:  newFromString("222.222222"),
			expect: newFromString("222.22222"),
		}, {
			input:  newFromString("2222022.2202"),
			expect: newFromString("2222022.2"),
		}, {
			input:  newFromString("218166.0002"),
			expect: newFromString("218166"),
		},
	}

	for _, test := range tests {
		tt := test
		t.Run(fmt.Sprintf("%s -> %s", tt.input, tt.expect), func(t *testing.T) {
			t.Parallel()

			actual := Convert(tt.input, sigDigits, maxPrecision)
			require.True(t, tt.expect.Equals(actual))
		})
	}
}

func TestValidate(t *testing.T) {
	t.Parallel()

	t.Run("Should return error on negative numbers", func(t *testing.T) {
		t.Parallel()

		input := newFromString("-12345")
		err := Validate(input, sigDigits, maxPrecision)
		require.Error(t, err)
	})

	t.Run("Should return error if number of significant digits is greater than allowed", func(t *testing.T) {
		t.Parallel()

		input := newFromString("1234567890")
		digits := int32(len(input.Coefficient().String()))
		require.Greater(t, digits, sigDigits)

		err := Validate(input, sigDigits, maxPrecision)
		require.Error(t, err)
	})

	t.Run("Should return error if precision is greater than allowed", func(t *testing.T) {
		t.Parallel()

		input := newFromString("0.00000000000000012345678")
		precision := int32(math.Abs(float64(input.Exponent())))
		require.Greater(t, precision, maxPrecision)

		err := Validate(input, sigDigits, maxPrecision)
		require.Error(t, err)
	})

	t.Run("Successful test", func(t *testing.T) {
		t.Parallel()

		input := newFromString("1.2345")
		digits := int32(len(input.Coefficient().String()))
		precision := int32(math.Abs(float64(input.Exponent())))
		require.LessOrEqual(t, digits, sigDigits)
		require.LessOrEqual(t, precision, maxPrecision)

		err := Validate(input, sigDigits, maxPrecision)
		require.NoError(t, err)
	})
}

func BenchmarkConvert_DecimalWithLeadingZeros(b *testing.B) {
	d := newFromString("0.000123456")

	// Reset timer to exclude time taken by setup operations
	// before the actual benchmark begins
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		Convert(d, sigDigits, maxPrecision)
	}
}

func BenchmarkConvert_TruncateDecimals(b *testing.B) {
	d := newFromString("1.00000000234")

	// Reset timer to exclude time taken by setup operations
	// before the actual benchmark begins
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		Convert(d, sigDigits, maxPrecision)
	}
}

func BenchmarkConvert_ExactSignificantDigits(b *testing.B) {
	d := newFromString("12345678")
	if int32(len(d.String())) != sigDigits {
		panic("expected number of digits to be equal to number of significant digits")
	}

	// Reset timer to exclude time taken by setup operations
	// before the actual benchmark begins
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		Convert(d, sigDigits, maxPrecision)
	}
}

func BenchmarkConvert_IntegralPartSizeGreaterThanSignificantDigits(b *testing.B) {
	d := newFromString("1000000060")
	if integralPart := strings.Split(d.String(), ".")[0]; int32(len(integralPart)) <= sigDigits {
		panic("expected number of digits to be greater than number of significant digits")
	}

	// Reset timer to exclude time taken by setup operations
	// before the actual benchmark begins
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		Convert(d, sigDigits, maxPrecision)
	}
}

func newFromString(num string) decimal.Decimal {
	d, err := decimal.NewFromString(num)
	if err != nil {
		panic(err)
	}
	return d
}
