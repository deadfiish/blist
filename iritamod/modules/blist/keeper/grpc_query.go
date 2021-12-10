package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/bianjieai/iritamod/modules/blist/types"
)

// QueryBListByName Check if the name is in the BList
func (k Keeper) QueryBListByName(goCtx context.Context, in *types.QueryBListByNameRequest) (*types.QueryBListByNameResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	state, err := k.GetBListByName(ctx, in.Name)
	if err != nil {
		return &types.QueryBListByNameResponse{}, err
	}
	return &types.QueryBListByNameResponse{Exist: state}, nil

}

// QueryBList get the BList
func (k Keeper) QueryBList(goCtx context.Context, in *types.QueryBListRequest) (*types.QueryBListResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	list, err := k.GetBList(ctx)
	if err != nil {
		return nil, err
	}
	return &types.QueryBListResponse{Blist: list}, nil
}