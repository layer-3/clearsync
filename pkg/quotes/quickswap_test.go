package quotes

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/layer-3/clearsync/pkg/abi/iquickswap_v3_pool"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/require"
)

func newBigInt(s string) *big.Int {
	x, ok := new(big.Int).SetString(s, 10)
	if !ok {
		panic(x)
	}
	return x
}

func Test_quickswap_parseSwap(t *testing.T) {
	t.Parallel()

	type args struct {
		swap *iquickswap_v3_pool.IQuickswapV3PoolSwap
		pool *quickswapPoolWrapper
	}

	tests := []struct {
		name    string
		args    args
		want    TradeEvent
		wantErr bool
	}{
		{
			name: "Buy trade",
			args: args{
				swap: &iquickswap_v3_pool.IQuickswapV3PoolSwap{
					// This is a REAL swap event from Polygon chain.
					// See at https://polygonscan.com/tx/0xe1051bb29489bcf4af4622346325d3179a58f6893bdaed48bb4190ababb73578
					Sender:    common.HexToAddress("0xf5b509bB0909a69B1c207E495f687a596C168E12"),
					Recipient: common.HexToAddress("0xf5b509bB0909a69B1c207E495f687a596C168E12"),
					Amount0:   newBigInt("+960868500000000000000"),
					Amount1:   newBigInt("-283933565629654405"),
					Price:     newBigInt("+1362218571192993247714913208"),
					Liquidity: newBigInt("+756062884725008673952739"),
					Tick:      newBigInt("-81269"),
					Raw: types.Log{
						Address:     common.HexToAddress("0x479e1B71A702a595e19b6d5932CD5c863ab57ee0"),
						BlockNumber: 54805723,
						TxHash:      common.BytesToHash([]byte("0xe1051bb29489bcf4af4622346325d3179a58f6893bdaed48bb4190ababb73578")),
						TxIndex:     66,
						BlockHash:   common.BytesToHash([]byte("0xc3bda0dbdc54a62253e4b4fb9c06819203d91b98600b901ce4bf7b482e1c35f1")),
						Index:       313,
						Removed:     false,
					},
				},
				pool: &quickswapPoolWrapper{
					baseToken: poolToken{
						Address:  "0x0d500B1d8E8eF31E21C99d1Db9A6444d3ADf1270",
						Symbol:   "matic",
						Decimals: decimal.NewFromInt(18),
					},
					quoteToken: poolToken{
						Address:  "0x7ceB23fD6bC0adD59E62ac25578270cFf1b9f619",
						Symbol:   "weth",
						Decimals: decimal.NewFromInt(18),
					},
					reverted: false,
				},
			},
			want: TradeEvent{
				Source:    DriverQuickswap,
				Market:    Market{baseUnit: "matic", quoteUnit: "weth"},
				Price:     decimal.RequireFromString("0.0002954967986042"),
				Amount:    decimal.RequireFromString("960.8685"),
				Total:     decimal.RequireFromString("0.2839335656296544"),
				TakerType: TakerTypeBuy,
			},
			wantErr: false,
		},
		// {
		// 	name: "Sell trade",
		// 	args: args{
		// 		swap: &iquickswap_v3_pool.IQuickswapV3PoolSwap{},
		// 		market: Market{
		// 			baseUnit:  "usdt",
		// 			quoteUnit: "usdc",
		// 		},
		// 		pool: &quickswapPoolWrapper{},
		// 	},
		// 	want:    TradeEvent{},
		// 	wantErr: false,
		// },
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			driver := quickswap{}
			got, err := driver.parseSwap(test.args.swap, test.args.pool)

			require.True(t, test.wantErr == (err != nil))

			require.Equal(t, test.want.Source, got.Source)
			require.Equal(t, test.want.Market, got.Market)
			require.True(t, test.want.Price.Equal(got.Price))
			require.True(t, test.want.Amount.Equal(got.Amount))
			require.True(t, test.want.Total.Equal(got.Total))
			require.Equal(t, test.want.TakerType, got.TakerType)
		})
	}
}
