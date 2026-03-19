package keeper

import (
	"context"

	"hub/x/hub/types"
	"hub/x/hub/utilfunc"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) SendEndErasmusPeriodRequest(goCtx context.Context, msg *types.MsgSendEndErasmusPeriodRequest) (*types.MsgSendEndErasmusPeriodRequestResponse, error) {

	// Logic before transmitting the packet

	ctx := sdk.UnwrapSDKContext(goCtx)
	utilfunc.PrintLogs("SendEndErasmusPeriodRequest", ctx)

	return &types.MsgSendEndErasmusPeriodRequestResponse{
		Status: -1,
	}, types.ErrNonCallableFunction

}
