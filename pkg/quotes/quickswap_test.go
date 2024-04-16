package quotes

import (
	"fmt"
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
		pool *dexPool[iquickswap_v3_pool.IQuickswapV3PoolSwap]
	}

	tests := []struct {
		name    string
		args    args
		want    TradeEvent
		wantErr bool
	}{
		{
			name: "Sell trade",
			args: args{
				swap: &iquickswap_v3_pool.IQuickswapV3PoolSwap{
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
				pool: &dexPool[iquickswap_v3_pool.IQuickswapV3PoolSwap]{
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
				Price:     decimal.RequireFromString("0.0002917105144227"),
				Amount:    decimal.RequireFromString("6035.9860273201849237"),
				Total:     decimal.RequireFromString("1.76076058907780048041581994904799"),
				TakerType: TakerTypeBuy,
			},
			wantErr: false,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			driver := quickswap{}
			got, err := driver.parseSwap(test.args.swap, test.args.pool)

			require.True(t, test.wantErr == (err != nil))

			require.Equal(t, test.want.Source, got.Source, fmt.Sprintf("want: `%s`, got `%s`", test.want.Source, got.Source))
			require.Equal(t, test.want.Market, got.Market, fmt.Sprintf("want: `%s`, got `%s`", test.want.Market, got.Market))
			require.True(t, test.want.Price.Equal(got.Price), fmt.Sprintf("want: `%s`, got `%s`", test.want.Price, got.Price))
			require.True(t, test.want.Amount.Equal(got.Amount), fmt.Sprintf("want: `%s`, got `%s`", test.want.Amount, got.Amount))
			require.True(t, test.want.Total.Equal(got.Total), fmt.Sprintf("want: `%s`, got `%s`", test.want.Total, got.Total))
			require.Equal(t, test.want.TakerType, got.TakerType, fmt.Sprintf("want: `%s`, got `%s`", test.want.TakerType, got.TakerType))
		})
	}
}
