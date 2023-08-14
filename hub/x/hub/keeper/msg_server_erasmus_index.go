package keeper

import (
	"context"

	"hub/x/hub/types"
	"hub/x/hub/utilfunc"
)

func (k msgServer) SendErasmusIndex(goCtx context.Context, msg *types.MsgSendErasmusIndex) (*types.MsgSendErasmusIndexResponse, error) {

	utilfunc.PrintLogs("SendErasmusIndex")

	return &types.MsgSendErasmusIndexResponse{
		Status: -1,
	}, types.ErrNonCallableFunction

}
