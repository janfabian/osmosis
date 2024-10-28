package arbot

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/osmosis-labs/osmosis/v26/x/arbot/keeper"
)

func BeginBlocker(ctx sdk.Context, k keeper.Keeper) {
	// Log a message at the INFO level
	ctx.Logger().Error(fmt.Sprintf("Arbot module: new block height %d", ctx.BlockHeight()))
}
