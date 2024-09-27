package encoding

import (
	amino "github.com/Null-nil/cosmos-sdk/codec"
	"github.com/Null-nil/cosmos-sdk/codec/types"
	"github.com/Null-nil/cosmos-sdk/simapp/params"
	"github.com/Null-nil/cosmos-sdk/types/module"
	"github.com/Null-nil/cosmos-sdk/x/auth/tx"

	enccodec "github.com/Null-nil/ethermint/encoding/codec"
)

// MakeConfig creates an EncodingConfig for testing
func MakeConfig(mb module.BasicManager) params.EncodingConfig {
	cdc := amino.NewLegacyAmino()
	interfaceRegistry := types.NewInterfaceRegistry()
	codec := amino.NewProtoCodec(interfaceRegistry)

	encodingConfig := params.EncodingConfig{
		InterfaceRegistry: interfaceRegistry,
		Codec:             codec,
		TxConfig:          tx.NewTxConfig(codec, tx.DefaultSignModes),
		Amino:             cdc,
	}

	enccodec.RegisterLegacyAminoCodec(encodingConfig.Amino)
	mb.RegisterLegacyAminoCodec(encodingConfig.Amino)
	enccodec.RegisterInterfaces(encodingConfig.InterfaceRegistry)
	mb.RegisterInterfaces(encodingConfig.InterfaceRegistry)
	return encodingConfig
}
