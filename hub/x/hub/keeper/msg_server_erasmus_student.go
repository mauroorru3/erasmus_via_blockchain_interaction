package keeper

import (
	"context"

	"hub/x/hub/types"
	"hub/x/hub/utilfunc"

	sdk "github.com/cosmos/cosmos-sdk/types"
	clienttypes "github.com/cosmos/ibc-go/v3/modules/core/02-client/types"
)

func (k msgServer) SendErasmusStudent(goCtx context.Context, msg *types.MsgSendErasmusStudent) (*types.MsgSendErasmusStudentResponse, error) {

	/*
		utilfunc.PrintLogs("SendErasmusStudent")

		return &types.MsgSendErasmusStudentResponse{
			Status: -1,
		}, types.ErrNonCallableFunction

	*/

	utilfunc.PrintLogs("SendErasmusStudent")

	ctx := sdk.UnwrapSDKContext(goCtx)

	var packet types.ErasmusStudentPacketData

	packet.Student = nil

	// Transmit the packet
	err := k.TransmitErasmusStudentPacket(
		ctx,
		packet,
		msg.Port,
		msg.ChannelID,
		clienttypes.ZeroHeight(),
		msg.TimeoutTimestamp,
		"SendErasmusStudent",
	)
	if err != nil {
		return nil, err
	} else {
		return &types.MsgSendErasmusStudentResponse{
			Status: 0,
		}, nil
	}

}
