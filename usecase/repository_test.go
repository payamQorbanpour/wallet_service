package usecase

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testsCreateWallet = []struct {
	name      string
	entry     repo
	want      Wallet
	wantError bool
}{
	{
		name: "create wallet success",
		entry: repo{
			db: map[string]int{
				"09123456789": 1300,
			},
		},
		want: Wallet{
			ID:      "09123456789",
			Balance: 1300,
		},
		wantError: false,
	},
}

func Test_repository_CreateWallet(t *testing.T) {
	for _, tt := range testsCreateWallet {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			wallet := Wallet{ID: tt.want.ID, Balance: tt.want.Balance}

			err := tt.entry.CreateWallet(context.Background(), wallet)

			if tt.wantError {
				assert.Error(t, err)
				return
			}
		})
	}
}

var testsGetWallet = []struct {
	name      string
	entry     repo
	want      Wallet
	wantError bool
}{
	{
		name: "get wallet success",
		entry: repo{
			db: map[string]int{
				"09123456789": 1300,
				"09123456780": 1400,
			},
		},
		want: Wallet{
			ID:      "09123456789",
			Balance: 1300,
		},
		wantError: false,
	},
	{
		name: "get wallet failed",
		entry: repo{
			db: map[string]int{},
		},
		wantError: true,
	},
}

func Test_repository_GetWallet(t *testing.T) {
	for _, tt := range testsGetWallet {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.entry.GetWallet(context.Background(), "09123456789")
			if tt.wantError {
				assert.Error(t, err)
				return
			}

			assert.Equal(t, tt.want, got)
		})
	}
}

var testsChargeWallet = []struct {
	name      string
	entry     repo
	want      Wallet
	wantError bool
}{
	{
		name: "charge wallet success",
		entry: repo{
			db: map[string]int{
				"09123456789": 1300,
				"09123456780": 1400,
			},
		},
		want: Wallet{
			ID:      "09123456789",
			Balance: 1301,
		},
		wantError: false,
	},
}

func Test_repository_ChargeWallet(t *testing.T) {
	for _, tt := range testsChargeWallet {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.entry.ChargeWallet(context.Background(), tt.want.ID, 1)
			if tt.wantError {
				assert.Error(t, err)
				return
			}

			assert.Equal(t, tt.want.Balance, got.Balance)
		})
	}
}

var testsCheckWalletExistance = []struct {
	name  string
	entry repo
	want  bool
}{
	{
		name: "check wallet existance true",
		entry: repo{
			db: map[string]int{
				"09123456789": 1300,
				"09123456780": 1400,
			},
		},
		want: true,
	},
	{
		name: "check wallet existance false",
		entry: repo{
			db: map[string]int{},
		},
		want: false,
	},
}

func Test_repository_CheckWalletExistance(t *testing.T) {
	for _, tt := range testsCheckWalletExistance {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got := tt.entry.checkWalletExistance(context.Background(), "09123456789")

			assert.Equal(t, tt.want, got)
		})
	}
}
