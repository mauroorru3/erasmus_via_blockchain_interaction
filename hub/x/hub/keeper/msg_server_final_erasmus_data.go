package keeper

import (
	"context"

	"hub/x/hub/types"
	"hub/x/hub/utilfunc"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) SendFinalErasmusData(goCtx context.Context, msg *types.MsgSendFinalErasmusData) (*types.MsgSendFinalErasmusDataResponse, error) {

	ctx := sdk.UnwrapSDKContext(goCtx)

	utilfunc.PrintLogs("SendFinalErasmusData", ctx)

	return &types.MsgSendFinalErasmusDataResponse{
		Status: -1,
	}, types.ErrNonCallableFunction

	/*



		// Logic before transmitting the packet

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
