package keeper

import (
	"context"

	"university_chain_it/x/universitychainit/types"
)

func (k msgServer) SendEndErasmusPeriodRequest(goCtx context.Context, msg *types.MsgSendEndErasmusPeriodRequest) (*types.MsgSendEndErasmusPeriodRequestResponse, error) {

	/*
		utilfunc.PrintLogs("SendEndErasmusPeriodRequest")

		return &types.MsgSendEndErasmusPeriodRequestResponse{
			Status: -1,
		}, types.ErrNonCallableFunction
	*/

	/*
		ctx := sdk.UnwrapSDKContext(goCtx)

		// TODO: logic before transmitting the packet

		// Construct the packet
		var packet types.EndErasmusPeriodRequestPacketData

		packet.DestinationUniversityName = msg.DestinationUniversityName
		packet.ForeignIndex = msg.ForeignIndex
		packet.StartingUniversityName = msg.StartingUniversityName
		packet.Index = msg.Index

		utilfunc.PrintLogs("prima di TransmitEndErasmusPeriodRequestPacket")

		// Transmit the packet
		err := k.TransmitEndErasmusPeriodRequestPacket(
			ctx,
			packet,
			msg.Port,
			msg.ChannelID,
			clienttypes.ZeroHeight(),
			msg.TimeoutTimestamp,
		)

		utilfunc.PrintLogs("dopo di TransmitEndErasmusPeriodRequestPacket")

		if err != nil {
			return &types.MsgSendEndErasmusPeriodRequestResponse{
				Status: -1,
			}, err
		} else {

			return &types.MsgSendEndErasmusPeriodRequestResponse{
				Status: 0,
			}, nil
		}
	*/

	return &types.MsgSendEndErasmusPeriodRequestResponse{
		Status: 0,
	}, nil
}
