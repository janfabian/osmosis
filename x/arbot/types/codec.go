package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
)

var (
	// ModuleCdc references the global x/arbot module codec
	ModuleCdc = codec.NewLegacyAmino()
	// Amino is a codec that uses the Amino library
	Amino = codec.NewLegacyAmino()
)

// RegisterLegacyAminoCodec registers the necessary x/arbot interfaces and concrete types
// on the provided LegacyAmino codec.
func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	// Register any concrete types here.
	// If your module defines messages or other types, register them like:
	// cdc.RegisterConcrete(MsgYourMessage{}, "arbot/YourMessage", nil)
}

// RegisterInterfaces registers the x/arbot module's interface types
func RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	// Register any interfaces and implementations here.
	// For example, if you have an interface and implementations:
	// registry.RegisterInterface(
	//     "osmosis.arbot.v1beta1.MyInterface",
	//     (*MyInterface)(nil),
	//     &MyImplementation{},
	// )
}

func init() {
	RegisterLegacyAminoCodec(ModuleCdc)
	// Seal the codec to prevent further modifications
	ModuleCdc.Seal()
}
