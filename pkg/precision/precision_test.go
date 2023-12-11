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
	sigDigits = 8
	maxScale  = 18
)

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

func TestToSignificant(t *testing.T) {
	t.Parallel()

	for _, test := range tests {
		tt := test
		t.Run(fmt.Sprintf("%s -> %s", tt.input, tt.expect), func(t *testing.T) {
			t.Parallel()

			actual := ToSignificant(tt.input, sigDigits, maxScale)
			if !actual.Equals(tt.expect) {
				t.Errorf("ToSignificant(%s): expected %s, got %s", tt.input, tt.expect, actual)
			}
		})
	}
}

func TestIsValid(t *testing.T) {
	t.Parallel()

	for _, test := range tests {
		tt := test
		t.Run(fmt.Sprintf("%s -> %s", tt.input, tt.expect), func(t *testing.T) {
			t.Parallel()

			isValid := IsValid(tt.input, sigDigits, maxScale)
			expect := tt.input.Equal(tt.expect)
			require.Equal(t, expect, isValid)
		})
	}

	t.Run("Should return false on negative numbers", func(t *testing.T) {
		d := newFromString("-12345")
		isValid := IsValid(d, sigDigits, maxScale)
		require.False(t, isValid)
	})
}

func BenchmarkToSignificant_DecimalWithLeadingZeros(b *testing.B) {
	d := newFromString("0.000123456")

	// Reset timer to exclude time taken by setup operations
	// before the actual benchmark begins
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		ToSignificant(d, sigDigits, maxScale)
	}
}

func BenchmarkToSignificant_TruncateDecimals(b *testing.B) {
	d := newFromString("1.00000000234")

	// Reset timer to exclude time taken by setup operations
	// before the actual benchmark begins
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		ToSignificant(d, sigDigits, maxScale)
	}
}

func BenchmarkToSignificant_ExactSignificantDigits(b *testing.B) {
	d := newFromString("12345678")
	if len(d.String()) != sigDigits {
		panic("expected number of digits to be equal to number of significant digits")
	}

	// Reset timer to exclude time taken by setup operations
	// before the actual benchmark begins
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		ToSignificant(d, sigDigits, maxScale)
	}
}

func BenchmarkToSignificant_IntegralPartSizeGreaterThanSignificantDigits(b *testing.B) {
	d := newFromString("1000000060")
	if integralPart := strings.Split(d.String(), ".")[0]; len(integralPart) <= sigDigits {
		panic("expected number of digits to be greater than number of significant digits")
	}

	// Reset timer to exclude time taken by setup operations
	// before the actual benchmark begins
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		ToSignificant(d, sigDigits, maxScale)
	}
}

func newFromString(num string) decimal.Decimal {
	d, err := decimal.NewFromString(num)
	if err != nil {
		panic(err)
	}
	return d
}
