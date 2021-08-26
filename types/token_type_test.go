package types

import (
	"testing"

	"github.com/trustwallet/go-primitives/coin"
)

func TestGetEthereumTokenTypeByIndex(t *testing.T) {
	type args struct {
		coinIndex uint
	}
	tests := []struct {
		name string
		args args
		want TokenType
	}{
		{
			"Ethereum ERC20",
			args{coinIndex: coin.ETHEREUM},
			ERC20,
		},
		{
			"Ethereum Classic ETC20",
			args{coinIndex: coin.CLASSIC},
			ETC20,
		},
		{
			"TomoChain TRC21",
			args{coinIndex: coin.TOMOCHAIN},
			TRC21,
		},
		{
			"WanChain WAN20",
			args{coinIndex: coin.WANCHAIN},
			WAN20,
		},
		{
			"Optimism ERC20",
			args{coinIndex: coin.OPTIMISM},
			OPTIMISM,
		},
		{
			"xDAI ERC20",
			args{coinIndex: coin.XDAI},
			XDAI,
		},
		{
			"Avalanche ERC20",
			args{coinIndex: coin.AVALANCHEC},
			AVALANCHE,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetEthereumTokenTypeByIndex(tt.args.coinIndex); got != tt.want {
				t.Errorf("GetEthereumTokenTypeByIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}
