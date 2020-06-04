package genaccounts

import (
	"encoding/base64"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/exported"
)

// ExportGenesis exports genesis for all accounts
func ExportGenesis(ctx sdk.Context, accountKeeper AccountKeeper, removeBlacklistAddrs map[string]bool) GenesisState {

	// iterate to get the accounts
	var accounts []GenesisAccount
	accountKeeper.IterateAccounts(ctx,
		func(acc exported.Account) (stop bool) {
			if removeBlacklistAddrs != nil {
				if acc.GetCoins().IsZero() {
					if _, ok := removeBlacklistAddrs[base64.StdEncoding.EncodeToString(acc.GetAddress())]; !ok {
						return false
					}
				}
			}

			account, err := NewGenesisAccountI(acc)
			if err != nil {
				panic(err)
			}
			accounts = append(accounts, account)
			return false
		},
	)

	return accounts
}
