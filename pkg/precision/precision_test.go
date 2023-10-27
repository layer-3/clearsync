package precision 

import (
	"fmt"
	"math"
	"testing"

	"github.com/shopspring/decimal"
)

func TestToSignificant(t *testing.T) {
	tests := []struct {
		input    decimal.Decimal
		expected decimal.Decimal
	}{
		{decimal.NewFromFloat(0.0), decimal.NewFromFloat(0.0)},
		{decimal.NewFromFloat(math.SmallestNonzeroFloat64), decimal.NewFromFloat(0.0)},
		{decimal.NewFromFloat(0.0001), decimal.NewFromFloat(0.0001)},
		{decimal.NewFromFloat(0.000123456), decimal.NewFromFloat(0.00012345)},
		{decimal.NewFromFloat(0.012344789), decimal.NewFromFloat(0.012345)},
		{decimal.NewFromFloat(0.012345), decimal.NewFromFloat(0.012345)},
		{decimal.NewFromFloat(1.00000234), decimal.NewFromFloat(1.0)},
		{decimal.NewFromInt(2), decimal.NewFromInt(2)},
		{decimal.NewFromInt(100006), decimal.NewFromInt(100010)},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(fmt.Sprintf("%s -> %s", tt.input, tt.expected), func(t *testing.T) {
			actual := ToSignificant(tt.input, 5, 8)
			if !actual.Equals(tt.expected) {
				t.Errorf("ToSignificant(%s): expected %s, got %s", tt.input, tt.expected, actual)
			}
		})
	}
}

func BenchmarkToSignificant_DecimalWithLeadingZeros(b *testing.B) {
	// Reset timer to exclude time taken by setup operations
	// before the actual benchmark begins
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		ToSignificant(decimal.NewFromFloat(0.000123456), 5, 8)
	}
}

func BenchmarkToSignificant_TruncateDecimals(b *testing.B) {
	// Reset timer to exclude time taken by setup operations
	// before the actual benchmark begins
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		ToSignificant(decimal.NewFromFloat(1.00000234), 5, 8)
	}
}

func BenchmarkToSignificant_IntegralPartSizeGreaterThanSignificantDigits(b *testing.B) {
	// Reset timer to exclude time taken by setup operations
	// before the actual benchmark begins
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		ToSignificant(decimal.NewFromInt(100006), 5, 8)
	}
}
