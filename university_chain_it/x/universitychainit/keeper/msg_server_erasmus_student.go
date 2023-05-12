package keeper

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"university_chain_it/x/universitychainit/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	clienttypes "github.com/cosmos/ibc-go/v3/modules/core/02-client/types"
)

func (k msgServer) SendErasmusStudent(goCtx context.Context, msg *types.MsgSendErasmusStudent) (*types.MsgSendErasmusStudentResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: logic before transmitting the packet

	file, err := os.OpenFile("data/logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		fmt.Println("Could not open logs.txt - SendErasmusStudent")
		return nil, err
	}

	defer file.Close()

	file.WriteString(strconv.FormatUint(msg.TimeoutTimestamp, 10))
	_, err2 := file.WriteString("\n")

	if err2 != nil {
		fmt.Println("Could not write text to logs.txt - SendErasmusStudent")

	} else {
		fmt.Println("Operation successful! Text has been appended to logs.txt - SendErasmusStudent")
	}

	student, found := k.GetStoredStudent(ctx, msg.Index)
	if !found {
		return nil, types.ErrStudentNotPresent
	} else {

		// Construct the packet
		var packet types.ErasmusStudentPacketData

		packet.Student = &student

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
		} else {
			return &types.MsgSendErasmusStudentResponse{}, nil
		}
	}
}
