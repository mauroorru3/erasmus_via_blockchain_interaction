package keeper

import (
	"context"

	"hub/x/hub/types"
	"hub/x/hub/utilfunc"
)

func (k msgServer) SendFinalErasmusData(goCtx context.Context, msg *types.MsgSendFinalErasmusData) (*types.MsgSendFinalErasmusDataResponse, error) {

	utilfunc.PrintLogs("SendFinalErasmusData")

	return &types.MsgSendFinalErasmusDataResponse{
		Status: -1,
	}, types.ErrNonCallableFunction

	/*

		ctx := sdk.UnwrapSDKContext(goCtx)

		// TODO: logic before transmitting the packet

		// Construct the packet
		var packet types.FinalErasmusDataPacketData

		packet.ErasmusData = msg.ErasmusData
		packet.HomeIndex = msg.HomeIndex

		// Transmit the packet
		err := k.TransmitFinalErasmusDataPacket(
			ctx,
			packet,
			msg.Port,
			msg.ChannelID,
			clienttypes.ZeroHeight(),
			msg.TimeoutTimestamp,
		)
		if err != nil {
			return nil, err
		}

		return &types.MsgSendFinalErasmusDataResponse{}, nil

	*/
}
