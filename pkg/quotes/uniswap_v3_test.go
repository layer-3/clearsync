package quotes

import (
	"fmt"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/layer-3/clearsync/pkg/abi/iuniswap_v3_pool"
)

func Test_uniswapV3_parseSwap(t *testing.T) {
	t.Parallel()

	type args struct {
		swap *iuniswap_v3_pool.IUniswapV3PoolSwap
		pool *dexPool[iuniswap_v3_pool.IUniswapV3PoolSwap, *iuniswap_v3_pool.IUniswapV3PoolSwapIterator]
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
				pool: &dexPool[iuniswap_v3_pool.IUniswapV3PoolSwap, *iuniswap_v3_pool.IUniswapV3PoolSwapIterator]{
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
				pool: &dexPool[iuniswap_v3_pool.IUniswapV3PoolSwap, *iuniswap_v3_pool.IUniswapV3PoolSwapIterator]{
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
				Price:     decimal.RequireFromString("3115.963161695177629"),
				Amount:    decimal.RequireFromString("0.000000000001"),
				Total:     decimal.RequireFromString("0.000002"),
				TakerType: TakerTypeBuy,
			},
			wantErr: false,
		},
		{
			name: "0xe1f72d25ee98ecbaa94e8b943b1eb833342f814e3f32cad24b771e9dd60e14fb",
			args: args{
				swap: &iuniswap_v3_pool.IUniswapV3PoolSwap{
					// This is a REAL swap event from Ethereum chain.
					// See at https://etherscan.io/tx/0xe1f72d25ee98ecbaa94e8b943b1eb833342f814e3f32cad24b771e9dd60e14fb
					Sender:       common.HexToAddress("0x3b3ae790df4f312e745d270119c6052904fb6790"),
					Recipient:    common.HexToAddress("0x3b3ae790df4f312e745d270119c6052904fb6790"),
					Amount0:      newBigInt("-998588202869525530"),
					Amount1:      newBigInt("3095357352"),
					SqrtPriceX96: newBigInt("4411194026449351583332407"),
					Liquidity:    newBigInt("334304175015581424"),
					Tick:         newBigInt("-195929"),
					Raw: types.Log{
						Address: common.HexToAddress("0xc7bbec68d12a0d1830360f8ec58fa599ba1b0e9b"),
						Topics: []common.Hash{
							common.HexToHash("0xc42079f94a6350d7e6235f29174924f928cc2ac818eb64fed8004e115fbcca67"),
							common.HexToHash("0x0000000000000000000000003b3ae790df4f312e745d270119c6052904fb6790"),
							common.HexToHash("0x0000000000000000000000003b3ae790df4f312e745d270119c6052904fb6790"),
						},
						Data:        []byte("0xfffffffffffffffffffffffffffffffffffffffffffffffff2244d51fb10f3e600000000000000000000000000000000000000000000000000000000b87f67a800000000000000000000000000000000000000000003a61b4ead9c135651f83700000000000000000000000000000000000000000000000004a3afe03ebbd6f0fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffd02a7"),
						BlockNumber: 0x12e64cf,
						TxHash:      common.BytesToHash([]byte("0xe1f72d25ee98ecbaa94e8b943b1eb833342f814e3f32cad24b771e9dd60e14fb")),
						TxIndex:     0x2a,
						BlockHash:   common.BytesToHash([]byte("0xcd14ca89d2e53afa2571f4ceb19773ae837c6cd7bd2fe8718ccab1e1a9b8c5ca")),
						Index:       0x50,
						Removed:     false,
					},
				},
				pool: &dexPool[iuniswap_v3_pool.IUniswapV3PoolSwap, *iuniswap_v3_pool.IUniswapV3PoolSwapIterator]{
					BaseToken: poolToken{
						Address:  common.HexToAddress("0x0"), // native token
						Symbol:   "eth",
						Decimals: decimal.NewFromInt(18),
					},
					QuoteToken: poolToken{
						Address:  common.HexToAddress("0xdAC17F958D2ee523a2206206994597C13D831ec7"),
						Symbol:   "usdt",
						Decimals: decimal.NewFromInt(6),
					},
					Reversed: false,
					Market:   NewMarket("eth", "usdt"),
				},
			},
			want: TradeEvent{
				Source:    DriverUniswapV3,
				Market:    NewMarket("eth", "usdt"),
				Price:     decimal.RequireFromString("3099.939041820825402"),
				Amount:    decimal.RequireFromString("0.9985882028695255"),
				Total:     decimal.RequireFromString("3095.357352"),
				TakerType: TakerTypeSell,
			},
			wantErr: false,
		},
		{
			name: "0x02a18c555fd367fcdeee6047a236129d472a2f70b5ffba280a0ae0fec9f43a13",
			args: args{
				swap: &iuniswap_v3_pool.IUniswapV3PoolSwap{
					// This is a REAL swap event from Ethereum chain.
					// See at https://etherscan.io/tx/0x02a18c555fd367fcdeee6047a236129d472a2f70b5ffba280a0ae0fec9f43a13
					Sender:       common.HexToAddress("0x3b3ae790df4f312e745d270119c6052904fb6790"),
					Recipient:    common.HexToAddress("0x3b3ae790df4f312e745d270119c6052904fb6790"),
					Amount0:      newBigInt("9886811256"),
					Amount1:      newBigInt("-3189117537947014095"),
					SqrtPriceX96: newBigInt("1423284675873585018512167229078758"),
					Liquidity:    newBigInt("10851620357380928399"),
					Tick:         newBigInt("195932"),
					Raw: types.Log{
						Address: common.HexToAddress("0x88e6a0c2ddd26feeb64f039a2c41296fcb3f5640"),
						Topics: []common.Hash{
							common.HexToHash("0xc42079f94a6350d7e6235f29174924f928cc2ac818eb64fed8004e115fbcca67"),
							common.HexToHash("0x0000000000000000000000003b3ae790df4f312e745d270119c6052904fb6790"),
							common.HexToHash("0x0000000000000000000000003b3ae790df4f312e745d270119c6052904fb6790"),
						},
						Data:        []byte("0x000000000000000000000000000000000000000000000000000000024d4cc478ffffffffffffffffffffffffffffffffffffffffffffffffd3bdfa7ef3b32c31000000000000000000000000000000000000462c60d206c5650ff04025da3ce60000000000000000000000000000000000000000000000009698b3387f16738f000000000000000000000000000000000000000000000000000000000002fd5c"),
						BlockNumber: 0x12e6535,
						TxHash:      common.BytesToHash([]byte("0x02a18c555fd367fcdeee6047a236129d472a2f70b5ffba280a0ae0fec9f43a13")),
						TxIndex:     0x28,
						BlockHash:   common.BytesToHash([]byte("0x82dff64bc5a0936a44f9287fb1e8bd48d0527a44585838f9e22a575078ed617d")),
						Index:       0x55,
						Removed:     false,
					},
				},
				pool: &dexPool[iuniswap_v3_pool.IUniswapV3PoolSwap, *iuniswap_v3_pool.IUniswapV3PoolSwapIterator]{
					Address: common.HexToAddress("0x88e6a0c2ddd26feeb64f039a2c41296fcb3f5640"),
					BaseToken: poolToken{
						Address:  common.HexToAddress("0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2"),
						Symbol:   "eth",
						Decimals: decimal.NewFromInt(18),
					},
					QuoteToken: poolToken{
						Address:  common.HexToAddress("0xdAC17F958D2ee523a2206206994597C13D831ec7"),
						Symbol:   "usdt",
						Decimals: decimal.NewFromInt(6),
					},
					Reversed: true,
					Market:   NewMarket("eth", "usdt"),
				},
			},
			want: TradeEvent{
				Source:    DriverUniswapV3,
				Market:    NewMarket("eth", "usdt"),
				Price:     decimal.RequireFromString("3098.6720911015389941"),
				Amount:    decimal.RequireFromString("3.1891175379470141"),
				Total:     decimal.RequireFromString("9886.811256"),
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
			assert.Equal(t, test.want.Source, got.Source, fmt.Sprintf("want Source: `%s`, got `%s`", test.want.Source, got.Source))
			assert.Equal(t, test.want.Market, got.Market, fmt.Sprintf("want Market: `%s`, got `%s`", test.want.Market, got.Market))
			assert.True(t, test.want.Price.Equal(got.Price), fmt.Sprintf("want Price: `%s`, got `%s`", test.want.Price, got.Price))
			assert.True(t, test.want.Amount.Equal(got.Amount), fmt.Sprintf("want Amount: `%s`, got `%s`", test.want.Amount, got.Amount))
			assert.True(t, test.want.Total.Equal(got.Total), fmt.Sprintf("want Total: `%s`, got `%s`", test.want.Total, got.Total))
			assert.Equal(t, test.want.TakerType, got.TakerType, fmt.Sprintf("want TakerType: `%s`, got `%s`", test.want.TakerType, got.TakerType))
		})
	}
}
