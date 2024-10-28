package keeper

import (
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"

	gammkeeper "github.com/osmosis-labs/osmosis/v26/x/gamm/keeper"
	gammtypes "github.com/osmosis-labs/osmosis/v26/x/gamm/types"
)

type Keeper struct {
	cdc        codec.Codec
	storeKey   storetypes.StoreKey
	paramstore paramtypes.Subspace

	// Keepers from other modules
	GammKeeper gammkeeper.Keeper
}

func NewKeeper(
	cdc codec.Codec,
	storeKey storetypes.StoreKey,
	paramstore paramtypes.Subspace,
	gammKeeper gammkeeper.Keeper,
) Keeper {
	return Keeper{
		cdc:        cdc,
		storeKey:   storeKey,
		paramstore: paramstore,
		GammKeeper: gammKeeper,
	}
}

// ListPools returns all pools from the GAMM module
func (k Keeper) ListPools(ctx sdk.Context) ([]gammtypes.CFMMPoolI, error) {
	pools, err := k.GammKeeper.GetPoolsAndPoke(ctx)
	if err != nil {
		return nil, err
	}
	return pools, nil
}
