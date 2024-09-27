package codec

import (
	"github.com/Null-nil/cosmos-sdk/codec"
	codectypes "github.com/Null-nil/cosmos-sdk/codec/types"
	"github.com/Null-nil/cosmos-sdk/std"
	sdk "github.com/Null-nil/cosmos-sdk/types"

	cryptocodec "github.com/Null-nil/ethermint/crypto/codec"
	ethermint "github.com/Null-nil/ethermint/types"
)

// RegisterLegacyAminoCodec registers Interfaces from types, crypto, and SDK std.
func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	sdk.RegisterLegacyAminoCodec(cdc)
	cryptocodec.RegisterCrypto(cdc)
	codec.RegisterEvidences(cdc)
}

// RegisterInterfaces registers Interfaces from types, crypto, and SDK std.
func RegisterInterfaces(interfaceRegistry codectypes.InterfaceRegistry) {
	std.RegisterInterfaces(interfaceRegistry)
	cryptocodec.RegisterInterfaces(interfaceRegistry)
	ethermint.RegisterInterfaces(interfaceRegistry)
}
