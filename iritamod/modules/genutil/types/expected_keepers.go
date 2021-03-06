package types

import (
	abci "github.com/tendermint/tendermint/abci/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
)

// NodeKeeper defines the expected node keeper (noalias)
type NodeKeeper interface {
	ApplyAndReturnValidatorSetUpdates(sdk.Context) (updates []abci.ValidatorUpdate, err error)
}

// AccountKeeper defines the expected account keeper (noalias)
type AccountKeeper interface {
	NewAccount(sdk.Context, authtypes.AccountI) authtypes.AccountI
	SetAccount(sdk.Context, authtypes.AccountI)
	IterateAccounts(ctx sdk.Context, process func(i authtypes.AccountI) (stop bool))
}
