package keeper

import (
	"context"

	"university_chain_it/x/universitychainit/types"
	"university_chain_it/x/universitychainit/utilfunc"
)

func (k msgServer) SendEndErasmusPeriodRequest(goCtx context.Context, msg *types.MsgSendEndErasmusPeriodRequest) (*types.MsgSendEndErasmusPeriodRequestResponse, error) {

	utilfunc.PrintLogs("SendEndErasmusPeriodRequest")

	return &types.MsgSendEndErasmusPeriodRequestResponse{
		Status: -1,
	}, types.ErrNonCallableFunction

}
