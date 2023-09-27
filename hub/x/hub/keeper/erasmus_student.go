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
	details string,
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

	sizeInt := packet.Size()
	utilfunc.GetTransactionStats("TransmitErasmusStudentPacket", details, ctx, sizeInt, packetBytes)

	if err := k.ChannelKeeper.SendPacket(ctx, channelCap, packet); err != nil {
		return err
	}

	return nil
}

// OnRecvErasmusStudentPacket processes packet reception
func (k Keeper) OnRecvErasmusStudentPacket(ctx sdk.Context, packet channeltypes.Packet, data types.ErasmusStudentPacketData) (packetAck types.ErasmusStudentPacketAck, err error) {

	sizeInt := packet.Size()
	binArray, err := data.GetBytes()
	if err != nil {
		return packetAck, err
	}
	utilfunc.GetTransactionStats("OnRecvErasmusStudentPacket", "", ctx, sizeInt, binArray)

	// validate packet data upon receiving
	if err := data.ValidateBasic(); err != nil {
		return packetAck, err
	}

	// TODO: packet reception logic

	utilfunc.PrintLogs("OnRecvErasmusStudentPacket")

	utilfunc.PrintData(data.String())

	foreignUni, err := utilfunc.GetForeignUniversityName(*data.Student)
	if err != nil {

		utilfunc.PrintLogs("OnRecvErasmusStudentPacket " + err.Error())
		return packetAck, err
	} else {

		uniInfo, found := k.GetUniversities(ctx, foreignUni)
		if !found {
			utilfunc.PrintLogs("OnRecvErasmusStudentPacket " + types.ErrWrongNameUniversity.Error())
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
				timeoutTimestamp,
				"")

			if err != nil {
				utilfunc.PrintLogs("OnRecvErasmusStudentPacket " + err.Error())
				return packetAck, err
			} else {

				err = utilfunc.GetConsumedGas("OnRecvErasmusStudentPacket", data.Student.Index, ctx)
				if err != nil {
					return packetAck, nil
				} else {

					utilfunc.PrintLogs("OnRecvErasmusStudentPacket packet sent")
					packetAckBytes, err := types.ModuleCdc.MarshalJSON(&packetAck)
					if err != nil {
						return packetAck, err
					}
					sizeInt := len(packetAckBytes)
					utilfunc.GetTransactionStats("OnRecvErasmusStudentPacket sending ack", "", ctx, sizeInt, binArray)
					return packetAck, nil
				}
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

		sizeInt := len(dispatchedAck.Result)
		binArray, err := data.GetBytes()
		if err != nil {
			return err
		}
		utilfunc.GetTransactionStats("OnAcknowledgementErasmusStudentPacket", "", ctx, sizeInt, binArray)

		if err := types.ModuleCdc.UnmarshalJSON(dispatchedAck.Result, &packetAck); err != nil {
			// The counter-party module doesn't implement the correct acknowledgment format
			utilfunc.PrintLogs("OnAcknowledgementErasmusStudentPacket cannot unmarshal acknowledgment")
			return errors.New("cannot unmarshal acknowledgment")
		}

		// TODO: successful acknowledgement logic

		utilfunc.PrintLogs("OnAcknowledgementErasmusStudentPacket")

		uniInfo, found := k.GetUniversities(ctx, data.Student.StudentData.UniversityName)
		if !found {
			utilfunc.PrintLogs("OnAcknowledgementErasmusStudentPacket " + types.ErrWrongNameUniversity.Error())
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
				timeoutTimestamp,
				"")

			if err != nil {
				utilfunc.PrintLogs("OnAcknowledgementErasmusStudentPacket error " + err.Error())
				return err
			} else {

				utilfunc.PrintLogs("OnAcknowledgementErasmusStudentPacket packet sent")

				err = utilfunc.GetConsumedGas("OnAcknowledgementErasmusStudentPacket", data.Student.Index, ctx)
				if err != nil {
					return err
				} else {

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
