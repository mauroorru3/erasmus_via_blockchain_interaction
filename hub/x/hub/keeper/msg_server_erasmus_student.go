package keeper

import (
	"context"

	"hub/x/hub/types"
	"hub/x/hub/utilfunc"
)

func (k msgServer) SendErasmusStudent(goCtx context.Context, msg *types.MsgSendErasmusStudent) (*types.MsgSendErasmusStudentResponse, error) {

	utilfunc.PrintLogs("SendErasmusStudent")

	return &types.MsgSendErasmusStudentResponse{
		Status: -1,
	}, types.ErrNonCallableFunction

}
