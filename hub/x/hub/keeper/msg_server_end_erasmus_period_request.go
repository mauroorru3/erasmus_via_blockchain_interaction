package keeper

import (
	"context"

	"hub/x/hub/types"
	"hub/x/hub/utilfunc"
)

func (k msgServer) SendEndErasmusPeriodRequest(goCtx context.Context, msg *types.MsgSendEndErasmusPeriodRequest) (*types.MsgSendEndErasmusPeriodRequestResponse, error) {

	// TODO: logic before transmitting the packet

	utilfunc.PrintLogs("SendEndErasmusPeriodRequest")

	return &types.MsgSendEndErasmusPeriodRequestResponse{
		Status: -1,
	}, types.ErrNonCallableFunction

}
