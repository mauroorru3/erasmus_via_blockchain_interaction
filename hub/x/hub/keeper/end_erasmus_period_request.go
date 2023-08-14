package keeper

import (
	"errors"

	"hub/x/hub/types"
	"hub/x/hub/utilfunc"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	clienttypes "github.com/cosmos/ibc-go/v3/modules/core/02-client/types"
	channeltypes "github.com/cosmos/ibc-go/v3/modules/core/04-channel/types"
	host "github.com/cosmos/ibc-go/v3/modules/core/24-host"
)

// TransmitEndErasmusPeriodRequestPacket transmits the packet over IBC with the specified source port and source channel
func (k Keeper) TransmitEndErasmusPeriodRequestPacket(
	ctx sdk.Context,
	packetData types.EndErasmusPeriodRequestPacketData,
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

// OnRecvEndErasmusPeriodRequestPacket processes packet reception
func (k Keeper) OnRecvEndErasmusPeriodRequestPacket(ctx sdk.Context, packet channeltypes.Packet, data types.EndErasmusPeriodRequestPacketData) (packetAck types.EndErasmusPeriodRequestPacketAck, err error) {
	// validate packet data upon receiving
	if err := data.ValidateBasic(); err != nil {
		return packetAck, err
	}

	utilfunc.PrintLogs("OnRecvEndErasmusPeriodRequestPacket")

	// TODO: packet reception logic

	packetAck.ErasmusRestrictedInfo = ""

	uniInfo, found := k.GetUniversities(ctx, data.DestinationUniversityName)
	if !found {
		return packetAck, types.ErrWrongNameUniversity
	} else {

		var packet_to_send types.EndErasmusPeriodRequestPacketData = data
		utilfunc.PrintData("OnRecvEndErasmusPeriodRequestPacket " + data.String())

		err := k.TransmitEndErasmusPeriodRequestPacket(
			ctx,
			packet_to_send,
			uniInfo.Port,
			uniInfo.ChannelID,
			clienttypes.ZeroHeight(),
			timeoutTimestamp)

		if err != nil {
			return packetAck, err
		} else {

			return packetAck, nil
		}
	}
}

// OnAcknowledgementEndErasmusPeriodRequestPacket responds to the the success or failure of a packet
// acknowledgement written on the receiving chain.
func (k Keeper) OnAcknowledgementEndErasmusPeriodRequestPacket(ctx sdk.Context, packet channeltypes.Packet, data types.EndErasmusPeriodRequestPacketData, ack channeltypes.Acknowledgement) error {
	switch dispatchedAck := ack.Response.(type) {
	case *channeltypes.Acknowledgement_Error:

		// TODO: failed acknowledgement logic
		_ = dispatchedAck.Error

		utilfunc.PrintLogs("OnAcknowledgementEndErasmusPeriodRequestPacket error " + dispatchedAck.Error)

		return nil
	case *channeltypes.Acknowledgement_Result:
		// Decode the packet acknowledgment
		var packetAck types.EndErasmusPeriodRequestPacketAck

		if err := types.ModuleCdc.UnmarshalJSON(dispatchedAck.Result, &packetAck); err != nil {
			// The counter-party module doesn't implement the correct acknowledgment format
			utilfunc.PrintLogs("OnAcknowledgementEndErasmusPeriodRequestPacket error " + err.Error())
			return errors.New("cannot unmarshal acknowledgment")
		}

		// TODO: successful acknowledgement logic

		utilfunc.PrintLogs("OnAcknowledgementEndErasmusPeriodRequestPacket success")

		uniInfo, found := k.GetUniversities(ctx, data.StartingUniversityName)
		if !found {
			return types.ErrWrongNameUniversity
		} else {

			var packet_to_send types.FinalErasmusDataPacketData

			packet_to_send.ErasmusRestrictedInfo = packetAck.ErasmusRestrictedInfo
			packet_to_send.HomeIndex = data.Index

			err := k.TransmitFinalErasmusDataPacket(ctx,
				packet_to_send,
				uniInfo.Port,
				uniInfo.ChannelID,
				clienttypes.ZeroHeight(),
				timeoutTimestamp)

			if err != nil {
				return err
			} else {
				utilfunc.PrintLogs("OnAcknowledgementEndErasmusPeriodRequestPacket msg sent")
				return nil
			}
		}

	default:
		// The counter-party module doesn't implement the correct acknowledgment format
		return errors.New("invalid acknowledgment format")
	}
}

// OnTimeoutEndErasmusPeriodRequestPacket responds to the case where a packet has not been transmitted because of a timeout
func (k Keeper) OnTimeoutEndErasmusPeriodRequestPacket(ctx sdk.Context, packet channeltypes.Packet, data types.EndErasmusPeriodRequestPacketData) error {

	// TODO: packet timeout logic

	utilfunc.PrintLogs("OnTimeoutEndErasmusPeriodRequestPacket")

	return nil
}
