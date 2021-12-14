package keeper

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/bianjieai/iritamod/modules/blist/types"
)

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the bank MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}

// MsgAddToBList add name to BList
func (k msgServer) MsgAddToBList(goCtx context.Context, list *types.MsgAddToBListRequest) (*types.MsgAddToBListResponse, error) {
	//if !common.IsHexAddress(list.Name) {
		//return &types.MsgAddToBListResponse{},
			//errors.Wrapf(types.ErrNameAlreadyExist, "name already exist")
	//}
	ctx := sdk.UnwrapSDKContext(goCtx)
	err := k.Keeper.AddToBList(ctx, list.Name)
	if err != nil {
		return &types.MsgAddToBListResponse{}, err
	}
	return &types.MsgAddToBListResponse{}, nil
}

// MsgRemoveFromBList remove name from BList
func (k msgServer) MsgRemoveFromBList(goCtx context.Context, list *types.MsgRemoveFromBListRequest) (*types.MsgRemoveFromBListResponse, error) {
	//if !common.IsHexAddress(list.Name) {
		//return &types.MsgRemoveFromBListResponse{},
			//errors.Wrapf(types.ErrNotFound, "name no found")
	//}
	ctx := sdk.UnwrapSDKContext(goCtx)
	err := k.Keeper.RemoveFromBList(ctx, list.Name)
	if err != nil {
		return &types.MsgRemoveFromBListResponse{}, err
	}
	return &types.MsgRemoveFromBListResponse{}, nil
}
