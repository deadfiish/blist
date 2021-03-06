package types

import "github.com/ethereum/go-ethereum/common"

// NewGenesisState constructs a new GenesisState instance
func NewGenesisState(contractAddress []string) *GenesisState {
	return &GenesisState{ContractAddress: contractAddress}
}

// DefaultGenesisState gets the raw genesis raw message for testing
func DefaultGenesisState() *GenesisState {
	return &GenesisState{}
}

// ValidateGenesis validates the provided identity genesis state to ensure the
// expected invariants holds.
func ValidateGenesis(data GenesisState) error {
	for _, contractAddress := range data.ContractAddress {
		if !common.IsHexAddress(contractAddress) {
			return ErrInvalidContractAddress
		}
	}
	return nil
}
