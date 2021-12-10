package types

import (
	"fmt"
)

const (
	// ModuleName string name of module
	ModuleName = "BList"

	StoreKey = ModuleName

	// RouterKey uses module name for routing
	RouterKey = ModuleName

	BListName = "BList"
)

const (
	KeyPrefixBList = "BList"
)

func BListPath(name string) string {
	return fmt.Sprintf("%s/%s", KeyPrefixBList, name)
}

// BListKey defines the full key under which a blist is stored.
func BListKey(name string) []byte {
	return []byte(BListPath(name))
}
