package keeper

import (
	"errors"
	"university_chain_de/x/universitychainde/types"
	"university_chain_de/x/universitychainde/utilfunc"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	clienttypes "github.com/cosmos/ibc-go/v3/modules/core/02-client/types"
	channeltypes "github.com/cosmos/ibc-go/v3/modules/core/04-channel/types"
	host "github.com/cosmos/ibc-go/v3/modules/core/24-host"
)

// TransmitErasmusIndexPacket transmits the packet over IBC with the specified source port and source channel
func (k Keeper) TransmitErasmusIndexPacket(
	ctx sdk.Context,
	packetData types.ErasmusIndexPacketData,
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

// OnRecvErasmusIndexPacket processes packet reception
func (k Keeper) OnRecvErasmusIndexPacket(ctx sdk.Context, packet channeltypes.Packet, data types.ErasmusIndexPacketData) (packetAck types.ErasmusIndexPacketAck, err error) {
	// validate packet data upon receiving
	err = data.ValidateBasic()

	if err != nil {
		return packetAck, err
	} else {

		// TODO: packet reception logic

		utilfunc.PrintLogs("OnRecvErasmusIndexPacket")

		searchedStudent, found := k.GetStoredStudent(ctx, data.Index)
		if !found {
			return packetAck, types.ErrStudentNotPresent
		} else {

			utilfunc.SetForeignIndex(&searchedStudent, data.ForeignIndex)

			k.SetStoredStudent(ctx, searchedStudent)

			return packetAck, nil
		}
	}
}

// OnAcknowledgementErasmusIndexPacket responds to the the success or failure of a packet
// acknowledgement written on the receiving chain.
func (k Keeper) OnAcknowledgementErasmusIndexPacket(ctx sdk.Context, packet channeltypes.Packet, data types.ErasmusIndexPacketData, ack channeltypes.Acknowledgement) error {
	switch dispatchedAck := ack.Response.(type) {
	case *channeltypes.Acknowledgement_Error:

		// TODO: failed acknowledgement logic
		_ = dispatchedAck.Error

		utilfunc.PrintLogs("OnAcknowledgementErasmusIndexPacket error " + dispatchedAck.Error)

		return nil
	case *channeltypes.Acknowledgement_Result:
		// Decode the packet acknowledgment
		var packetAck types.ErasmusIndexPacketAck

		if err := types.ModuleCdc.UnmarshalJSON(dispatchedAck.Result, &packetAck); err != nil {
			// The counter-party module doesn't implement the correct acknowledgment format
			return errors.New("cannot unmarshal acknowledgment")
		}

		// TODO: successful acknowledgement logic

		utilfunc.PrintLogs("OnAcknowledgementErasmusIndexPacket success")

		return nil
	default:
		// The counter-party module doesn't implement the correct acknowledgment format
		return errors.New("invalid acknowledgment format")
	}
}

// OnTimeoutErasmusIndexPacket responds to the case where a packet has not been transmitted because of a timeout
func (k Keeper) OnTimeoutErasmusIndexPacket(ctx sdk.Context, packet channeltypes.Packet, data types.ErasmusIndexPacketData) error {

	// TODO: packet timeout logic

	utilfunc.PrintLogs("OnTimeoutErasmusIndexPacket")

	return nil
}
