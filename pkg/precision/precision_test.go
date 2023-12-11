package precision

import (
	"fmt"
	"math"
	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/require"
)

func TestToSignificant(t *testing.T) {
	t.Parallel()

	tests := []struct {
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
			input:  newFromString("0.000103456"),
			expect: newFromString("0.00010346"),
		}, {
			input:  newFromString("0.012344789"),
			expect: newFromString("0.012345"),
		}, {
			input:  newFromString("0.001023499"),
			expect: newFromString("0.0010235"),
		}, {
			input:  newFromString("0.012345928"),
			expect: newFromString("0.012346"),
		}, {
			input:  newFromString("0.001023499"),
			expect: newFromString("0.0010235"),
		}, {
			input:  newFromString("0.012345"),
			expect: newFromString("0.012345"),
		}, {
			input:  newFromString("1.00000234"),
			expect: newFromString("1.0"),
		}, {
			input:  newFromString("1.0002"),
			expect: newFromString("1.0002"),
		}, {
			input:  newFromString("2"),
			expect: newFromString("2"),
		}, {
			input:  newFromString("100006"),
			expect: newFromString("100010"),
		}, {
			input:  newFromString("12345"),
			expect: newFromString("12345"),
		}, {
			input:  newFromString("222.222222"),
			expect: newFromString("222.22"),
		}, {
			input:  newFromString("2222022.2202"),
			expect: newFromString("2222000"),
		}, {
			input:  newFromString("218166.0002"),
			expect: newFromString("218170"),
		},
	}

	for _, test := range tests {
		tt := test
		t.Run(fmt.Sprintf("%s -> %s", tt.input, tt.expect), func(t *testing.T) {
			t.Parallel()

			actual := ToSignificant(tt.input, 5, 18)
			if !actual.Equals(tt.expect) {
				t.Errorf("ToSignificant(%s): expected %s, got %s", tt.input, tt.expect, actual)
			}
		})
	}
}

func BenchmarkToSignificant_DecimalWithLeadingZeros(b *testing.B) {
	d := newFromString("0.000123456")

	// Reset timer to exclude time taken by setup operations
	// before the actual benchmark begins
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		ToSignificant(d, 5, 18)
	}
}

func BenchmarkToSignificant_TruncateDecimals(b *testing.B) {
	d := newFromString("1.00000234")

	// Reset timer to exclude time taken by setup operations
	// before the actual benchmark begins
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		ToSignificant(d, 5, 18)
	}
}

func BenchmarkToSignificant_ExactSignificantDigits(b *testing.B) {
	num := "12345"
	d := newFromString(num)
	sigDigits := int32(len(num))

	// Reset timer to exclude time taken by setup operations
	// before the actual benchmark begins
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		ToSignificant(d, sigDigits, 18)
	}
}

func BenchmarkToSignificant_IntegralPartSizeGreaterThanSignificantDigits(b *testing.B) {
	d := newFromString("100006")

	// Reset timer to exclude time taken by setup operations
	// before the actual benchmark begins
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		ToSignificant(d, 5, 18)
	}
}

func TestIsValid(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input  decimal.Decimal
		expect bool
	}{
		{
			input:  newFromString("0.0"),
			expect: true,
		}, {
			input:  decimal.NewFromFloat(math.SmallestNonzeroFloat64),
			expect: false,
		}, {
			input:  newFromString("0.0000000000000004"),
			expect: true,
		}, {
			input:  newFromString("0.0001"),
			expect: true,
		}, {
			input:  newFromString("0.000103456"),
			expect: false,
		}, {
			input:  newFromString("0.012344789"),
			expect: false,
		}, {
			input:  newFromString("0.001023499"),
			expect: false,
		}, {
			input:  newFromString("0.012345928"),
			expect: false,
		}, {
			input:  newFromString("0.001023499"),
			expect: false,
		}, {
			input:  newFromString("0.012345"),
			expect: true,
		}, {
			input:  newFromString("1.00000234"),
			expect: false,
		}, {
			input:  newFromString("1.0002"),
			expect: true,
		}, {
			input:  newFromString("2"),
			expect: true,
		}, {
			input:  newFromString("100006"),
			expect: false,
		}, {
			input:  newFromString("12345"),
			expect: true,
		}, {
			input:  newFromString("222.222222"),
			expect: false,
		}, {
			input:  newFromString("2222022.2202"),
			expect: false,
		}, {
			input:  newFromString("218166.0002"),
			expect: false,
		},
	}

	for _, test := range tests {
		tt := test
		t.Run(fmt.Sprintf("%s -> %t", tt.input, tt.expect), func(t *testing.T) {
			t.Parallel()

			isValid := IsValid(tt.input, 5, 18)
			require.Equal(t, tt.expect, isValid)
		})
	}
}

func newFromString(num string) decimal.Decimal {
	d, err := decimal.NewFromString(num)
	if err != nil {
		panic(err)
	}
	return d
}
