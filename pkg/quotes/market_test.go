package quotes

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func Test_sortTradeEventsInPlace(t *testing.T) {
	t.Parallel()

	now := time.Now()
	tests := []struct {
		name string
		args []TradeEvent
		want []TradeEvent
	}{
		{
			name: "empty",
			args: nil,
			want: nil,
		},
		{
			name: "single",
			args: []TradeEvent{{CreatedAt: now}},
			want: []TradeEvent{{CreatedAt: now}},
		},
		{
			name: "multiple",
			args: []TradeEvent{
				{CreatedAt: now.Add(-2 * time.Hour)},
				{CreatedAt: now.Add(-1 * time.Hour)},
				{CreatedAt: now.Add(-3 * time.Hour)},
			},
			want: []TradeEvent{
				{CreatedAt: now.Add(-1 * time.Hour)},
				{CreatedAt: now.Add(-2 * time.Hour)},
				{CreatedAt: now.Add(-3 * time.Hour)},
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			sortTradeEventsInPlace(tt.args)
			require.Equal(t, tt.want, tt.args)
		})
	}
}
