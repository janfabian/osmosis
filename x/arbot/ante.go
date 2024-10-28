package arbot

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type ArbotAnteDecorator struct{}

// NewArbotAnteDecorator creates a new ArbotAnteDecorator instance.
func NewArbotAnteDecorator() ArbotAnteDecorator {
	return ArbotAnteDecorator{}
}

// AnteHandle logs transactions as they enter the mempool.
func (aad ArbotAnteDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (sdk.Context, error) {
	// // Log transaction hash and messages
	// txBytes := ctx.TxBytes()
	// tx, error := sdk.TxDecoder(txBytes)

	// txHash := fmt.Sprintf("%X", sdk.TxDecoder(txBytes).Hash())
	// ctx.Logger().Info("Arbot module: New transaction in mempool", "txHash", txHash)
	ctx.Logger().Error("Arbot module: New transaction in mempool")

	// Access transaction messages
	msgs := tx.GetMsgs()
	for _, msg := range msgs {
		ctx.Logger().Error("Arbot module: Transaction message", "msgType", sdk.MsgTypeURL(msg))
	}

	// Proceed to the next AnteDecorator
	return next(ctx, tx, simulate)
}
