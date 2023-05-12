package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	clienttypes "github.com/cosmos/ibc-go/v3/modules/core/02-client/types"
	"hub/x/hub/types"
)

func (k msgServer) SendErasmusStudent(goCtx context.Context, msg *types.MsgSendErasmusStudent) (*types.MsgSendErasmusStudentResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: logic before transmitting the packet

	// Construct the packet
	var packet types.ErasmusStudentPacketData

	packet.Student = msg.Student

	// Transmit the packet
	err := k.TransmitErasmusStudentPacket(
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

	return &types.MsgSendErasmusStudentResponse{}, nil
}
