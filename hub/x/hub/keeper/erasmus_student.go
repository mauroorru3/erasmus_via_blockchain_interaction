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

const timeoutTimestamp uint64 = 17999999999999999999

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

	utilfunc.PrintLogs("OnRecvErasmusStudentPacket")

	//---------------------------------------

	utilfunc.PrintData(data.String())

	packetAck.ForeignIndex = "-1"

	foreignUni, err := utilfunc.GetForeignUniversityName(*data.Student)
	if err != nil {
		return packetAck, err
	} else {

		utilfunc.PrintLogs("OnRecvErasmusStudentPacket " + foreignUni)

		uniInfo, found := k.GetUniversities(ctx, foreignUni)
		if !found {
			return packetAck, types.ErrWrongNameUniversity
		} else {

			var packet_to_send types.ErasmusStudentPacketData

			packet_to_send.Student = data.Student

			err := k.TransmitErasmusStudentPacket(
				ctx,
				packet_to_send,
				uniInfo.Port,
				uniInfo.ChannelID,
				clienttypes.ZeroHeight(),
				timeoutTimestamp)

			if err != nil {
				utilfunc.PrintLogs("OnRecvErasmusStudentPacket " + err.Error())
				return packetAck, err
			} else {

				utilfunc.PrintLogs("OnRecvErasmusStudentPacket packet sent")
				return packetAck, nil
			}
		}
	}

}

// OnAcknowledgementErasmusStudentPacket responds to the the success or failure of a packet
// acknowledgement written on the receiving chain.
func (k Keeper) OnAcknowledgementErasmusStudentPacket(ctx sdk.Context, packet channeltypes.Packet, data types.ErasmusStudentPacketData, ack channeltypes.Acknowledgement) error {
	switch dispatchedAck := ack.Response.(type) {
	case *channeltypes.Acknowledgement_Error:

		// TODO: failed acknowledgement logic

		utilfunc.PrintLogs("OnAcknowledgementErasmusStudentPacket error " + dispatchedAck.Error)

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

		utilfunc.PrintLogs("OnAcknowledgementErasmusStudentPacket success")

		foreignUni, err := utilfunc.GetForeignUniversityName(*data.Student)
		if err != nil {
			return err
		} else {

			uniInfo, found := k.GetUniversities(ctx, foreignUni)
			if !found {
				return types.ErrWrongNameUniversity
			} else {

				var packet_to_send types.ErasmusIndexPacketData

				packet_to_send.Index = data.Student.Index
				packet_to_send.ForeignIndex = packetAck.ForeignIndex

				err := k.TransmitErasmusIndexPacket(ctx,
					packet_to_send,
					uniInfo.Port,
					uniInfo.ChannelID,
					clienttypes.ZeroHeight(),
					timeoutTimestamp)

				if err != nil {
					utilfunc.PrintLogs("OnAcknowledgementErasmusStudentPacket error " + err.Error())
					return err
				} else {

					utilfunc.PrintLogs("OnAcknowledgementErasmusStudentPacket packet sent")
					return nil
				}
			}
		}

	default:
		// The counter-party module doesn't implement the correct acknowledgment format
		return errors.New("invalid acknowledgment format")
	}
}

// OnTimeoutErasmusStudentPacket responds to the case where a packet has not been transmitted because of a timeout
func (k Keeper) OnTimeoutErasmusStudentPacket(ctx sdk.Context, packet channeltypes.Packet, data types.ErasmusStudentPacketData) error {

	// TODO: packet timeout logic

	utilfunc.PrintLogs("OnTimeoutErasmusStudentPacket")

	return nil
}
