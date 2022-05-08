package cic_net

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/lmittmann/w3"
	"math/big"
	"testing"
)

func TestCicNet_ERC20Token_ERC20TokenInfo(t *testing.T) {
	type args struct {
		contractAddress common.Address
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
		symbol  string
	}{
		{
			name: "Token at kitabu sarafu",
			args: args{
				contractAddress: w3.A("0xaB89822F31c2092861F713F6F34bd6877a8C1878"),
			},
			wantErr: false,
			symbol:  "SRF",
		},
		{
			name: "Token at zero address",
			args: args{
				contractAddress: w3.A("0x0000000000000000000000000000000000000000"),
			},
			wantErr: true,
			symbol:  "",
		},
	}

	for _, testcase := range tests {
		tt := testcase

		t.Run(tt.name, func(t *testing.T) {
			cicnet, err := NewCicNet(conf.rpcProvider, w3.A(conf.tokenIndex))

			if err != nil {
				t.Fatalf("NewCicNet error = %v", err)
			}

			got, err := cicnet.ERC20TokenInfo(context.Background(), tt.args.contractAddress)

			if (err != nil) != tt.wantErr {
				t.Errorf("ERC20TokenInfo() error = %v, wantErr %v", err, tt.wantErr)
			}

			if got.Symbol != tt.symbol {
				t.Fatalf("Token = %v, want %v", got, tt.symbol)
			}
		})
	}
}

func TestCicNet_ERC20Token_BalanceOf(t *testing.T) {
	type args struct {
		contractAddress common.Address
		accountAddress  common.Address
	}

	tests := []struct {
		name       string
		args       args
		wantErr    bool
		balanceGte big.Int
	}{
		{
			name: "Sarafu sink balance",
			args: args{
				contractAddress: w3.A("0xaB89822F31c2092861F713F6F34bd6877a8C1878"),
				accountAddress:  w3.A("0xBBb4a93c8dCd82465B73A143f00FeD4AF7492a27"),
			},
			wantErr:    false,
			balanceGte: *big.NewInt(1),
		},
		{
			name: "Dead address balance",
			args: args{
				contractAddress: w3.A("0xaB89822F31c2092861F713F6F34bd6877a8C1878"),
				accountAddress:  w3.A("0x000000000000000000000000000000000000dEaD"),
			},
			wantErr:    false,
			balanceGte: *big.NewInt(0),
		},
	}

	for _, testcase := range tests {
		tt := testcase

		t.Run(tt.name, func(t *testing.T) {
			cicnet, err := NewCicNet(conf.rpcProvider, w3.A(conf.tokenIndex))

			if err != nil {
				t.Fatalf("NewCicNet error = %v", err)
			}

			got, err := cicnet.BalanceOf(context.Background(), tt.args.contractAddress, tt.args.accountAddress)

			if (err != nil) != tt.wantErr {
				t.Errorf("BalanceOf() error = %v, wantErr %v", err, tt.wantErr)
			}

			if got.Cmp(&tt.balanceGte) < 0 {
				t.Fatalf("Token = %v, want %d", got, tt.balanceGte.Int64())
			}
		})
	}
}
