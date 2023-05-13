package keeper

import (
	"errors"
	"fmt"
	"os"
	"time"

	"university_chain_it/x/universitychainit/types"
	"university_chain_it/x/universitychainit/utilfunc"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	clienttypes "github.com/cosmos/ibc-go/v3/modules/core/02-client/types"
	channeltypes "github.com/cosmos/ibc-go/v3/modules/core/04-channel/types"
	host "github.com/cosmos/ibc-go/v3/modules/core/24-host"
)

// TransmitErasmusStudentPacket transmits the packet over IBC with the specified source port and source channel
func (k Keeper) TransmitErasmusStudentPacket(
	ctx sdk.Context,
	packetData types.ErasmusStudentPacketData,
	sourcePort,
	sourceChannel string,
	timeoutHeight clienttypes.Height,
	timeoutTimestamp uint64,
) error {

	sourceChannelEnd, found := k.ChannelKeeper.GetChannel(ctx, sourcePort, sourceChannel)
	if !found {
		return sdkerrors.Wrapf(channeltypes.ErrChannelNotFound, "port ID (%s) channel ID (%s)", sourcePort, sourceChannel)
	}

	destinationPort := sourceChannelEnd.GetCounterparty().GetPortID()
	destinationChannel := sourceChannelEnd.GetCounterparty().GetChannelID()

	// get the next sequence
	sequence, found := k.ChannelKeeper.GetNextSequenceSend(ctx, sourcePort, sourceChannel)
	if !found {
		return sdkerrors.Wrapf(
			channeltypes.ErrSequenceSendNotFound,
			"source port: %s, source channel: %s", sourcePort, sourceChannel,
		)
	}

	channelCap, ok := k.ScopedKeeper.GetCapability(ctx, host.ChannelCapabilityPath(sourcePort, sourceChannel))
	if !ok {
		return sdkerrors.Wrap(channeltypes.ErrChannelCapabilityNotFound, "module does not own channel capability")
	}

	packetBytes, err := packetData.GetBytes()
	if err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, "cannot marshal the packet: "+err.Error())
	}

	packet := channeltypes.NewPacket(
		packetBytes,
		sequence,
		sourcePort,
		sourceChannel,
		destinationPort,
		destinationChannel,
		timeoutHeight,
		timeoutTimestamp,
	)

	if err := k.ChannelKeeper.SendPacket(ctx, channelCap, packet); err != nil {
		return err
	}

	return nil
}

// OnRecvErasmusStudentPacket processes packet reception
func (k Keeper) OnRecvErasmusStudentPacket(ctx sdk.Context, packet channeltypes.Packet, data types.ErasmusStudentPacketData) (packetAck types.ErasmusStudentPacketAck, err error) {
	// validate packet data upon receiving
	if err := data.ValidateBasic(); err != nil {
		return packetAck, err
	}

	// TODO: packet reception logic

	return packetAck, nil
}

// OnAcknowledgementErasmusStudentPacket responds to the the success or failure of a packet
// acknowledgement written on the receiving chain.
func (k Keeper) OnAcknowledgementErasmusStudentPacket(ctx sdk.Context, packet channeltypes.Packet, data types.ErasmusStudentPacketData, ack channeltypes.Acknowledgement) error {
	switch dispatchedAck := ack.Response.(type) {
	case *channeltypes.Acknowledgement_Error:

		// TODO: failed acknowledgement logic

		file, err := os.OpenFile("data/logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

		if err != nil {
			fmt.Println("Could not open logs.txt - OnAcknowledgementErasmusStudentPacket")
			return nil
		}

		defer file.Close()

		dt := time.Now()

		_, err2 := file.WriteString("OnAcknowledgementErasmusStudentPacket error " + dispatchedAck.Error + " " + utilfunc.FormatDeadline(dt) + "\n")

		if err2 != nil {
			fmt.Println("Could not write text to logs.txt - OnAcknowledgementErasmusStudentPacket")

		} else {
			fmt.Println("Operation successful! Text has been appended to logs.txt - OnAcknowledgementErasmusStudentPacket")
		}

		_ = dispatchedAck.Error

		return nil
	case *channeltypes.Acknowledgement_Result:
		// Decode the packet acknowledgment
		var packetAck types.ErasmusStudentPacketAck

		if err := types.ModuleCdc.UnmarshalJSON(dispatchedAck.Result, &packetAck); err != nil {
			// The counter-party module doesn't implement the correct acknowledgment format
			return errors.New("cannot unmarshal acknowledgment")
		}

		// TODO: successful acknowledgement logic
		file, err := os.OpenFile("data/logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

		if err != nil {
			fmt.Println("Could not open logs.txt - OnAcknowledgementErasmusStudentPacket")
			return nil
		}

		defer file.Close()

		dt := time.Now()

		_, err2 := file.WriteString("OnAcknowledgementErasmusStudentPacket success " + utilfunc.FormatDeadline(dt) + "\n")

		if err2 != nil {
			fmt.Println("Could not write text to logs.txt - OnAcknowledgementErasmusStudentPacket")

		} else {
			fmt.Println("Operation successful! Text has been appended to logs.txt - OnAcknowledgementErasmusStudentPacket")
		}

		return nil
	default:
		// The counter-party module doesn't implement the correct acknowledgment format
		return errors.New("invalid acknowledgment format")
	}
}

// OnTimeoutErasmusStudentPacket responds to the case where a packet has not been transmitted because of a timeout
func (k Keeper) OnTimeoutErasmusStudentPacket(ctx sdk.Context, packet channeltypes.Packet, data types.ErasmusStudentPacketData) error {

	// TODO: packet timeout logic

	file, err := os.OpenFile("data/logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		fmt.Println("Could not open logs.txt - OnTimeoutErasmusStudentPacket")
		return nil
	}

	defer file.Close()

	dt := time.Now()

	_, err2 := file.WriteString("OnTimeoutErasmusStudentPacket " + utilfunc.FormatDeadline(dt) + "\n")

	if err2 != nil {
		fmt.Println("Could not write text to logs.txt - OnTimeoutErasmusStudentPacket")

	} else {
		fmt.Println("Operation successful! Text has been appended to logs.txt - sOnTimeoutErasmusStudentPacket")
	}

	return nil
}
