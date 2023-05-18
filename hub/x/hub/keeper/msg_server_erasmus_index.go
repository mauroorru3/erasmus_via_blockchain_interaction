package keeper

import (
	"context"

	"hub/x/hub/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	clienttypes "github.com/cosmos/ibc-go/v3/modules/core/02-client/types"
)

func (k msgServer) SendErasmusIndex(goCtx context.Context, msg *types.MsgSendErasmusIndex) (*types.MsgSendErasmusIndexResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: logic before transmitting the packet

	// Construct the packet
	var packet types.ErasmusIndexPacketData

	packet.Index = msg.Index
	packet.ForeignIndex = msg.ForeignIndex

	// Transmit the packet
	err := k.TransmitErasmusIndexPacket(
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

	return &types.MsgSendErasmusIndexResponse{}, nil
}
