package usecase

import (
	"context"
	"os"
	"testing"

	"github.com/go-kit/kit/log"
	"github.com/stretchr/testify/assert"
)

var testsLogicCreateWallet = []struct {
	name      string
	entry     service
	want      Wallet
	wantError bool
}{
	{
		name: "create wallet success",
		entry: service{
			repository: &repo{
				db: map[string]int{},
			},
			logger: log.NewLogfmtLogger(os.Stderr),
		},
		want: Wallet{
			ID:      "09123456789",
			Balance: 1300,
		},
		wantError: false,
	},
	{
		name: "create wallet failure",
		entry: service{
			repository: &repo{
				db: map[string]int{
					"09123456789": 1300,
				},
			},
			logger: log.NewLogfmtLogger(os.Stderr),
		},
		wantError: true,
	},
}

func Test_logic_CreateWallet(t *testing.T) {
	for _, tt := range testsLogicCreateWallet {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.entry.CreateWallet(context.Background(), 1300, "09123456789")
			if tt.wantError {
				assert.Error(t, err)
				return
			}

			assert.Equal(t, tt.want.ID, got)
		})
	}
}

var testsLogicGetBalance = []struct {
	name      string
	entry     service
	want      Wallet
	wantError bool
}{
	{
		name: "get balance success",
		entry: service{
			repository: &repo{
				db: map[string]int{
					"09123456789": 1300,
				},
			},
			logger: log.NewLogfmtLogger(os.Stderr),
		},
		want: Wallet{
			ID:      "09123456789",
			Balance: 1300,
		},
		wantError: false,
	},
	{
		name: "get balance failure",
		entry: service{
			repository: &repo{
				db: map[string]int{},
			},
			logger: log.NewLogfmtLogger(os.Stderr),
		},
		want: Wallet{
			ID:      "09123456789",
			Balance: 1300,
		},
		wantError: true,
	},
}

func Test_logic_GetBalance(t *testing.T) {
	for _, tt := range testsLogicGetBalance {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.entry.GetBalance(context.Background(), "09123456789")
			if tt.wantError {
				assert.Error(t, err)
				return
			}

			assert.Equal(t, tt.want.Balance, got)
		})
	}
}

var testsLogicChargeWallet = []struct {
	name      string
	entry     service
	want      Wallet
	wantError bool
}{
	{
		name: "charge wallet success",
		entry: service{
			repository: &repo{
				db: map[string]int{
					"09123456789": 1300,
				},
			},
			logger: log.NewLogfmtLogger(os.Stderr),
		},
		want: Wallet{
			ID:      "09123456789",
			Balance: 1301,
		},
		wantError: false,
	},
	{
		name: "charge wallet failure",
		entry: service{
			repository: &repo{
				db: map[string]int{},
			},
			logger: log.NewLogfmtLogger(os.Stderr),
		},
		want: Wallet{
			ID:      "09123456789",
			Balance: 1301,
		},
		wantError: true,
	},
}

func Test_logic_ChargeWallet(t *testing.T) {
	for _, tt := range testsLogicChargeWallet {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.entry.ChargeWallet(context.Background(), "09123456789", 1)
			if tt.wantError {
				assert.Error(t, err)
				return
			}

			assert.Equal(t, tt.want.Balance, got)
		})
	}
}
