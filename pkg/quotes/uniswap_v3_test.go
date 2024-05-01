package quotes

import (
	"fmt"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/require"

	"github.com/layer-3/clearsync/pkg/abi/iuniswap_v3_pool"
)

func Test_uniswapV3_parseSwap(t *testing.T) {
	t.Parallel()

	type args struct {
		swap *iuniswap_v3_pool.IUniswapV3PoolSwap
		pool *dexPool[iuniswap_v3_pool.IUniswapV3PoolSwap]
	}

	tests := []struct {
		name    string
		args    args
		want    TradeEvent
		wantErr bool
	}{
		{
			name: "0x261125387e73be9e527ac6cbc57c1741b293d8ffea44c47e4330b6948b7e50d2",
			args: args{
				swap: &iuniswap_v3_pool.IUniswapV3PoolSwap{
					// This is a REAL swap event from Ethereum chain.
					// See at https://etherscan.io/tx/0x261125387e73be9e527ac6cbc57c1741b293d8ffea44c47e4330b6948b7e50d2
					Sender:       common.HexToAddress("0x000000000c923384110e9dca557279491e00f521"),
					Recipient:    common.HexToAddress("0x000000000c923384110e9dca557279491e00f521"),
					Amount0:      newBigInt("-3"),
					Amount1:      newBigInt("+963265801"),
					SqrtPriceX96: newBigInt("1419330143817945704579005219627373"),
					Liquidity:    newBigInt("+9843516503795384297"),
					Tick:         newBigInt("195877"),
					Raw: types.Log{
						Address: common.HexToAddress("0x88e6a0c2ddd26feeb64f039a2c41296fcb3f5640"),
						Topics: []common.Hash{
							common.HexToHash("0xc42079f94a6350d7e6235f29174924f928cc2ac818eb64fed8004e115fbcca67"),
							common.HexToHash("0x000000000000000000000000000000000c923384110e9dca557279491e00f521"),
							common.HexToHash("0x000000000000000000000000000000000c923384110e9dca557279491e00f521"),
						},
						Data:        []byte("0xfffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffd00000000000000000000000000000000000000000000000000000000396a450900000000000000000000000000000000000045fa7709c73f0259b5386e14616d000000000000000000000000000000000000000000000000889b321b656b3fe9000000000000000000000000000000000000000000000000000000000002fd25"),
						BlockNumber: 0x12d9ec9,
						TxHash:      common.BytesToHash([]byte("0x261125387e73be9e527ac6cbc57c1741b293d8ffea44c47e4330b6948b7e50d2")),
						TxIndex:     0x5,
						BlockHash:   common.BytesToHash([]byte("0x14f8f16a6229afe8da139fef4d7f83ab590c5e9df9314a516d313f7aad5ef769")),
						Index:       0x24,
						Removed:     false,
					},
				},
				pool: &dexPool[iuniswap_v3_pool.IUniswapV3PoolSwap]{
					BaseToken: poolToken{
						Address:  common.HexToAddress("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"),
						Symbol:   "weth",
						Decimals: decimal.NewFromInt(18),
					},
					QuoteToken: poolToken{
						Address:  common.HexToAddress("0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48"),
						Symbol:   "usdc",
						Decimals: decimal.NewFromInt(6),
					},
					Reversed: true,
					Market:   NewMarket("weth", "usdc"),
				},
			},
			want: TradeEvent{
				Source:    DriverUniswapV3,
				Market:    NewMarket("weth", "usdc"),
				Price:     decimal.RequireFromString("3115.9631616951436041"),
				Amount:    decimal.RequireFromString("0.0000000009632658"),
				Total:     decimal.RequireFromString("0.000003"),
				TakerType: TakerTypeSell,
			},
			wantErr: false,
		},
		{
			name: "0x2f5e87b7f645599a422a86eaca8fc11e5368e1242d5f06f2687bf5c6736dbbc7",
			args: args{
				swap: &iuniswap_v3_pool.IUniswapV3PoolSwap{
					// This is a REAL swap event from Ethereum chain.
					// See at https://etherscan.io/tx/0x2f5e87b7f645599a422a86eaca8fc11e5368e1242d5f06f2687bf5c6736dbbc7
					Sender:       common.HexToAddress("0x000000000c923384110e9dca557279491e00f521"),
					Recipient:    common.HexToAddress("0x000000000c923384110e9dca557279491e00f521"),
					Amount0:      newBigInt("+2"),
					Amount1:      newBigInt("-1000000"),
					SqrtPriceX96: newBigInt("1419330143817937955354376479905162"),
					Liquidity:    newBigInt("9843516503795384297"),
					Tick:         newBigInt("195877"),
					Raw: types.Log{
						Address: common.HexToAddress("0x88e6a0c2ddd26feeb64f039a2c41296fcb3f5640"),
						Topics: []common.Hash{
							common.HexToHash("0xc42079f94a6350d7e6235f29174924f928cc2ac818eb64fed8004e115fbcca67"),
							common.HexToHash("0x000000000000000000000000000000000c923384110e9dca557279491e00f521"),
							common.HexToHash("0x000000000000000000000000000000000c923384110e9dca557279491e00f521"),
						},
						Data:        []byte("0x0000000000000000000000000000000000000000000000000000000000000002fffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0bdc000000000000000000000000000000000000045fa7709c73e96ceee7a4d57158a000000000000000000000000000000000000000000000000889b321b656b3fe9000000000000000000000000000000000000000000000000000000000002fd25"),
						BlockNumber: 0x12d9ec9,
						TxHash:      common.BytesToHash([]byte("0x2f5e87b7f645599a422a86eaca8fc11e5368e1242d5f06f2687bf5c6736dbbc7")),
						TxIndex:     0x3,
						BlockHash:   common.BytesToHash([]byte("0x14f8f16a6229afe8da139fef4d7f83ab590c5e9df9314a516d313f7aad5ef769")),
						Index:       0x13,
						Removed:     false,
					},
				},
				pool: &dexPool[iuniswap_v3_pool.IUniswapV3PoolSwap]{
					BaseToken: poolToken{
						Address:  common.HexToAddress("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"),
						Symbol:   "weth",
						Decimals: decimal.NewFromInt(18),
					},
					QuoteToken: poolToken{
						Address:  common.HexToAddress("0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48"),
						Symbol:   "usdc",
						Decimals: decimal.NewFromInt(6),
					},
					Reversed: true,
					Market:   NewMarket("weth", "usdc"),
				},
			},
			want: TradeEvent{
				Source:    DriverUniswapV3,
				Market:    NewMarket("weth", "usdc"),
				Price:     decimal.RequireFromString("0.0003209280559838"),
				Amount:    decimal.RequireFromString("0.000002"),
				Total:     decimal.RequireFromString("0.000000000001"),
				TakerType: TakerTypeBuy,
			},
			wantErr: false,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			driver := uniswapV3{}
			got, err := driver.parseSwap(test.args.swap, test.args.pool)

			if test.wantErr {
				require.True(t, err != nil)
				return
			}
			require.Equal(t, test.want.Source, got.Source, fmt.Sprintf("want Source: `%s`, got `%s`", test.want.Source, got.Source))
			require.Equal(t, test.want.Market, got.Market, fmt.Sprintf("want Market: `%s`, got `%s`", test.want.Market, got.Market))
			require.True(t, test.want.Price.Equal(got.Price), fmt.Sprintf("want Price: `%s`, got `%s`", test.want.Price, got.Price))
			require.True(t, test.want.Amount.Equal(got.Amount), fmt.Sprintf("want Amount: `%s`, got `%s`", test.want.Amount, got.Amount))
			require.True(t, test.want.Total.Equal(got.Total), fmt.Sprintf("want Total: `%s`, got `%s`", test.want.Total, got.Total))
			require.Equal(t, test.want.TakerType, got.TakerType, fmt.Sprintf("want TakerType: `%s`, got `%s`", test.want.TakerType, got.TakerType))
		})
	}
}
