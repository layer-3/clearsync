package driver

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/layer-3/clearsync/pkg/artifacts/quickswap_v3_pool"
	quotes_common "github.com/layer-3/clearsync/pkg/quotes/common"
	"github.com/layer-3/clearsync/pkg/quotes/driver/base"
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
		swap *quickswap_v3_pool.IQuickswapV3PoolSwap
		pool *base.DexPool[quickswap_v3_pool.IQuickswapV3PoolSwap, *quickswap_v3_pool.IQuickswapV3PoolSwapIterator]
	}

	tests := []struct {
		name    string
		args    args
		want    quotes_common.TradeEvent
		wantErr bool
	}{
		{
			name: "0xa77f02fe9abda2ab43d77bc3ef4cf19bc75f60085b0437b2321e9a89248c6dc6",
			args: args{
				swap: &quickswap_v3_pool.IQuickswapV3PoolSwap{
					// This is a REAL swap event from Polygon chain.
					// See at https://polygonscan.com/tx/0xa77f02fe9abda2ab43d77bc3ef4cf19bc75f60085b0437b2321e9a89248c6dc6
					Sender:    common.HexToAddress("0xf5b509bB0909a69B1c207E495f687a596C168E12"),
					Recipient: common.HexToAddress("0x403022AF121cDCA9C4AcF4B94B2934429594eA29"),
					Amount0:   newBigInt("+6035986027320184923721"),
					Amount1:   newBigInt("-1760760589077853834"),
					Price:     newBigInt("+1354461025717973021575669155"),
					Liquidity: newBigInt("+710909202585697062870544"),
					Tick:      newBigInt("-81383"),
					Raw: types.Log{
						Address: common.HexToAddress("0x479e1B71A702a595e19b6d5932CD5c863ab57ee0"),
						Topics: []common.Hash{
							common.HexToHash("0xc42079f94a6350d7e6235f29174924f928cc2ac818eb64fed8004e115fbcca67"),
							common.HexToHash("0x000000000000000000000000f5b509bb0909a69b1c207e495f687a596c168e12"),
							common.HexToHash("0x000000000000000000000000403022af121cdca9c4acf4b94b2934429594ea29"),
						},
						BlockNumber: 54877299,
						TxHash:      common.BytesToHash([]byte("0xa77f02fe9abda2ab43d77bc3ef4cf19bc75f60085b0437b2321e9a89248c6dc6")),
						TxIndex:     2,
						BlockHash:   common.BytesToHash([]byte("0x3c083a9522b1cc23cf0139eabf382387846722bdc8a4946129ffc38b727ae9bf")),
						Index:       21,
						Removed:     false,
					},
				},
				pool: &base.DexPool[quickswap_v3_pool.IQuickswapV3PoolSwap, *quickswap_v3_pool.IQuickswapV3PoolSwapIterator]{
					BaseToken: base.DexPoolToken{
						Address:  common.HexToAddress("0x0d500B1d8E8eF31E21C99d1Db9A6444d3ADf1270"),
						Symbol:   "wmatic",
						Decimals: decimal.NewFromInt(18),
					},
					QuoteToken: base.DexPoolToken{
						Address:  common.HexToAddress("0x7ceB23fD6bC0adD59E62ac25578270cFf1b9f619"),
						Symbol:   "weth",
						Decimals: decimal.NewFromInt(18),
					},
					Reversed: true,
					Market:   quotes_common.NewMarket("wmatic", "weth"),
				},
			},
			want: quotes_common.TradeEvent{
				Source:    quotes_common.DriverQuickswap,
				Market:    quotes_common.NewMarket("wmatic", "weth"),
				Price:     decimal.RequireFromString("3421.5756126711160865"),
				Amount:    decimal.RequireFromString("1.7607605890778538"),
				Total:     decimal.RequireFromString("6035.9860273201849237"),
				TakerType: quotes_common.TakerTypeBuy,
			},
			wantErr: false,
		},
		{
			name: "0x43671f14131020bad4265ce9d9110589d2554875b2d554bed7cdb8766ac3be12",
			args: args{
				swap: &quickswap_v3_pool.IQuickswapV3PoolSwap{
					// This is a REAL swap event from Polygon chain.
					// See at https://polygonscan.com/tx/0x43671f14131020bad4265ce9d9110589d2554875b2d554bed7cdb8766ac3be12
					Sender:    common.HexToAddress("0x802b65b5d9016621e66003aed0b16615093f328b"),
					Recipient: common.HexToAddress("0x802b65b5d9016621e66003aed0b16615093f328b"),
					Amount0:   newBigInt("+366413350971486913"),
					Amount1:   newBigInt("-1059059758"),
					Price:     newBigInt("4262211506663986413431754"),
					Liquidity: newBigInt("189778357758686423"),
					Tick:      newBigInt("-196616"),
					Raw: types.Log{
						Address: common.HexToAddress("0x9ceff2f5138fc59eb925d270b8a7a9c02a1810f2"),
						Topics: []common.Hash{
							common.HexToHash("0xc42079f94a6350d7e6235f29174924f928cc2ac818eb64fed8004e115fbcca67"),
							common.HexToHash("0x000000000000000000000000802b65b5d9016621e66003aed0b16615093f328b"),
							common.HexToHash("0x000000000000000000000000802b65b5d9016621e66003aed0b16615093f328b"),
						},
						Data:        []byte("0x0000000000000000000000000000000000000000000000000515c300599db2c1ffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0e007d200000000000000000000000000000000000000000003868ef2e1aa8165d3ebca00000000000000000000000000000000000000000000000002a23a69304374d7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffcfff8"),
						BlockNumber: 0x35d81b2,
						TxHash:      common.BytesToHash([]byte("0x43671f14131020bad4265ce9d9110589d2554875b2d554bed7cdb8766ac3be12")),
						TxIndex:     0x27,
						BlockHash:   common.BytesToHash([]byte("0xd3c0b0229699bd9a5eeaf92e7e3162b7c36610324174d634d1d64c7221daa86d")),
						Index:       0x81,
						Removed:     false,
					},
				},
				pool: &base.DexPool[quickswap_v3_pool.IQuickswapV3PoolSwap, *quickswap_v3_pool.IQuickswapV3PoolSwapIterator]{
					Address: common.HexToAddress("0x9ceff2f5138fc59eb925d270b8a7a9c02a1810f2"),
					BaseToken: base.DexPoolToken{
						Address:  common.HexToAddress("0x7ceb23fd6bc0add59e62ac25578270cff1b9f619"),
						Symbol:   "weth",
						Decimals: decimal.NewFromInt(18),
					},
					QuoteToken: base.DexPoolToken{
						Address:  common.HexToAddress("0xc2132d05d31c914a87c6611c10748aeb04b58e8f"),
						Symbol:   "usdt",
						Decimals: decimal.NewFromInt(6),
					},
					Reversed: false,
					Market:   quotes_common.NewMarket("weth", "usdt"),
				},
			},
			want: quotes_common.TradeEvent{
				Source:    quotes_common.DriverQuickswap,
				Market:    quotes_common.NewMarket("weth", "usdt"),
				Price:     decimal.RequireFromString("2894.0819654229875328"),
				Amount:    decimal.RequireFromString("0.3664133509714869"),
				Total:     decimal.RequireFromString("1059.059758"),
				TakerType: quotes_common.TakerTypeBuy,
			},
			wantErr: false,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			driver := quickswap{}
			parser := driver.buildParser(test.args.swap, test.args.pool)

			logger := loggerQuickswap.With("swap", test.args.swap)
			got, err := parser.ParseSwap(quotes_common.DriverQuickswap, logger)

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
