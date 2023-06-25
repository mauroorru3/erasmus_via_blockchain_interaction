package keeper

import (
	"context"

	"university_chain_de/x/universitychainde/types"
)

func (k msgServer) SendExtendErasmusPeriod(goCtx context.Context, msg *types.MsgSendExtendErasmusPeriod) (*types.MsgSendExtendErasmusPeriodResponse, error) {

	return &types.MsgSendExtendErasmusPeriodResponse{}, nil
}
