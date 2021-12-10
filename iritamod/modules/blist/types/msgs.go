package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var (
	_ sdk.Msg = &MsgAddToBListRequest{}
	_ sdk.Msg = &MsgRemoveFromBListRequest{}
)

func NewMsgAddToBList(name string) *MsgAddToBListRequest {
	return &MsgAddToBListRequest{
		name,
	}
}

func (m MsgAddToBListRequest) ValidateBasic() error {
	return nil
}

func (m *MsgAddToBListRequest) GetSigners() []sdk.AccAddress {
	return nil
}

func NewMsgRemoveFromBList(name string) *MsgRemoveFromBListRequest {
	return &MsgRemoveFromBListRequest{
		name,
	}
}

func (m MsgRemoveFromBListRequest) ValidateBasic() error {
	return nil
}

func (m *MsgRemoveFromBListRequest) GetSigners() []sdk.AccAddress {
	return nil
}
