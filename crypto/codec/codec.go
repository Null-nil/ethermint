package codec

import (
	codectypes "github.com/Null-nil/cosmos-sdk/codec/types"
	cryptotypes "github.com/Null-nil/cosmos-sdk/crypto/types"

	"github.com/Null-nil/ethermint/crypto/ethsecp256k1"
)

// RegisterInterfaces register the Ethermint key concrete types.
func RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	registry.RegisterImplementations((*cryptotypes.PubKey)(nil), &ethsecp256k1.PubKey{})
	registry.RegisterImplementations((*cryptotypes.PrivKey)(nil), &ethsecp256k1.PrivKey{})
}
