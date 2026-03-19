package keeper

import (
	"context"

	"university_chain_de/x/universitychainde/types"
)

func (k msgServer) SendEndErasmusPeriodRequest(goCtx context.Context, msg *types.MsgSendEndErasmusPeriodRequest) (*types.MsgSendEndErasmusPeriodRequestResponse, error) {

	return &types.MsgSendEndErasmusPeriodRequestResponse{
		Status: -1,
	}, types.ErrNonCallableFunction
}
