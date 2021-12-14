package keeper

import (
	"github.com/bianjieai/iritamod/modules/blist/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
)

// AddToBList add name to BList
func (k Keeper) AddToBList(ctx sdk.Context, name string) error {
	store := k.GetStore(ctx)
	if get := store.Get(types.BListKey(name)); get != nil {
		return errors.Wrap(types.ErrNameAlreadyExist, "name already in BList")
	}
	store.Set(types.BListKey(name), []byte("true"))
	return nil
}

// RemoveFromBList remove name from BList
func (k Keeper) RemoveFromBList(ctx sdk.Context, name string) error {
	store := k.GetStore(ctx)
	get := store.Get(types.BListKey(name))
	if get != nil {
		store.Delete(types.BListKey(name))
	} else {
		return errors.Wrapf(types.ErrNotFound, "the %s is not in blist", name)
	}
	return nil
}

// GetBListByName Check if the name is in the BList
func (k Keeper) GetBListByName(ctx sdk.Context, name string) (bool, error) {
	store := k.GetStore(ctx)
	get := store.Get(types.BListKey(name))
	if get != nil {
		return true, nil
	} else {
		return false, nil
	}
}

// GetBList get the BList
func (k Keeper) GetBList(ctx sdk.Context) ([]string, error) {
	list, err := k.IteratorBList(ctx)
	if err != nil {
		return nil, err
	}
	return list, nil
}

