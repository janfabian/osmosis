package arbot

import (
	"context"
	"encoding/json"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/gorilla/mux"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/spf13/cobra"

	"github.com/osmosis-labs/osmosis/v26/x/arbot/keeper"
	"github.com/osmosis-labs/osmosis/v26/x/arbot/types"
	pooltypes "github.com/osmosis-labs/osmosis/v26/x/poolmanager/types"
)

var (
	_ module.AppModule      = AppModule{}
	_ module.AppModuleBasic = AppModuleBasic{}
)

// AppModuleBasic defines the basic application module used by the arbot module.
type AppModuleBasic struct {
	cdc codec.BinaryCodec
}

func NewAppModuleBasic(cdc codec.BinaryCodec) AppModuleBasic {
	return AppModuleBasic{cdc: cdc}
}

// Name returns the arbot module's name.
func (AppModuleBasic) Name() string { return types.ModuleName }

// RegisterLegacyAminoCodec registers the arbot module's types for the legacy Amino codec.
func (AppModuleBasic) RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	types.RegisterLegacyAminoCodec(cdc)
}

// RegisterInterfaces registers the module's interface types.
func (a AppModuleBasic) RegisterInterfaces(reg cdctypes.InterfaceRegistry) {
	types.RegisterInterfaces(reg)
}

// DefaultGenesis returns default genesis state as raw bytes for the arbot module.
func (AppModuleBasic) DefaultGenesis(cdc codec.JSONCodec) json.RawMessage {
	return []byte("{}")
}

// ValidateGenesis performs genesis state validation for the arbot module.
func (AppModuleBasic) ValidateGenesis(cdc codec.JSONCodec, config client.TxEncodingConfig, bz json.RawMessage) error {
	return nil
}

// RegisterRESTRoutes registers the REST routes for the arbot module.
func (AppModuleBasic) RegisterRESTRoutes(ctx client.Context, rtr *mux.Router) {
	// REST routes (if any)
}

// RegisterGRPCGatewayRoutes registers the gRPC Gateway routes for the arbot module.
func (AppModuleBasic) RegisterGRPCGatewayRoutes(clientCtx client.Context, mux *runtime.ServeMux) {
	// gRPC Gateway routes (if any)
}

// GetTxCmd returns the root tx command for the arbot module.
func (AppModuleBasic) GetTxCmd() *cobra.Command {
	return nil // No tx commands
}

// GetQueryCmd returns the root query command for the arbot module.
func (AppModuleBasic) GetQueryCmd() *cobra.Command {
	return nil // No query commands
}

//____________________________________________________________________________

// AppModule implements an application module for the arbot module.
type AppModule struct {
	AppModuleBasic

	keeper     keeper.Keeper
	gammKeeper pooltypes.PoolModuleI
}

// IsAppModule implements the appmodule.AppModule interface.
func (am AppModule) IsAppModule() {}

// IsOnePerModuleType is a marker function just indicates that this is a one-per-module type.
func (am AppModule) IsOnePerModuleType() {}

// NewAppModule creates a new AppModule object
func NewAppModule(cdc codec.Codec, k keeper.Keeper, gammKeeper pooltypes.PoolModuleI) AppModule {
	return AppModule{
		AppModuleBasic: NewAppModuleBasic(cdc),
		keeper:         k,
		gammKeeper:     gammKeeper,
	}
}

// Name returns the arbot module's name.
func (AppModule) Name() string { return types.ModuleName }

// RegisterInvariants registers the arbot module invariants.
func (am AppModule) RegisterInvariants(ir sdk.InvariantRegistry) {}

// QuerierRoute returns the arbot module's querier route name.
func (AppModule) QuerierRoute() string { return types.QuerierRoute }

// RegisterServices registers module services.
func (am AppModule) RegisterServices(cfg module.Configurator) {}

// InitGenesis performs genesis initialization for the arbot module.
func (am AppModule) InitGenesis(ctx sdk.Context, cdc codec.JSONCodec, data json.RawMessage) {

}

// ExportGenesis returns the exported genesis state as raw bytes for the arbot module.
func (am AppModule) ExportGenesis(ctx sdk.Context, cdc codec.JSONCodec) json.RawMessage {
	return []byte("{}")

}

// BeginBlock executes all ABCI BeginBlock logic respective to the capability module.
func (am AppModule) BeginBlock(context context.Context) error {
	ctx := sdk.UnwrapSDKContext(context)

	BeginBlocker(ctx, am.keeper)

	return nil
}

func (am AppModule) EndBlock(context context.Context) error {
	return nil
}
