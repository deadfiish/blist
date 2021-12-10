package types

import (
	"errors"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	codeErrNotFound = uint32(iota) + 2
	codeErrNameAlreadyExist
)

var ErrPostTxProcessing = errors.New("failed to execute post processing")

var (
	// ErrNotFound returns an error when not found contract from to deny list
	ErrNotFound = sdkerrors.Register(ModuleName, codeErrNotFound, "not found")

	// ErrNameAlreadyExist returns an error that the contract is already in ContractDenyList
	ErrNameAlreadyExist = sdkerrors.Register(ModuleName, codeErrNameAlreadyExist, "name already exist")
)

