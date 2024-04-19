package keeper

import (
	"context"

	"hub/x/hub/types"
	"hub/x/hub/utilfunc"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) SendErasmusIndex(goCtx context.Context, msg *types.MsgSendErasmusIndex) (*types.MsgSendErasmusIndexResponse, error) {

	ctx := sdk.UnwrapSDKContext(goCtx)
	utilfunc.PrintLogs("SendErasmusIndex", ctx)

	return &types.MsgSendErasmusIndexResponse{
		Status: -1,
	}, types.ErrNonCallableFunction

}
