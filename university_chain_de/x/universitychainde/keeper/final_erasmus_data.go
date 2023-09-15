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

// TransmitFinalErasmusDataPacket transmits the packet over IBC with the specified source port and source channel
func (k Keeper) TransmitFinalErasmusDataPacket(
	ctx sdk.Context,
	packetData types.FinalErasmusDataPacketData,
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

	sizeInt := packet.Size()
	utilfunc.GetTransactionStats("TransmitEndErasmusPeriodRequestPacket", ctx, sizeInt, packetBytes)

	if err := k.ChannelKeeper.SendPacket(ctx, channelCap, packet); err != nil {
		return err
	}

	return nil
}

// OnRecvFinalErasmusDataPacket processes packet reception
func (k Keeper) OnRecvFinalErasmusDataPacket(ctx sdk.Context, packet channeltypes.Packet, data types.FinalErasmusDataPacketData) (packetAck types.FinalErasmusDataPacketAck, err error) {

	sizeInt := packet.Size()
	binArray, err := data.GetBytes()
	if err != nil {
		return packetAck, err
	}
	utilfunc.GetTransactionStats("OnRecvFinalErasmusDataPacket", ctx, sizeInt, binArray)

	// validate packet data upon receiving
	if err := data.ValidateBasic(); err != nil {
		return packetAck, err
	}

	utilfunc.PrintLogs("OnRecvFinalErasmusDataPacket")

	// TODO: packet reception logic

	searchedStudent, found := k.GetStoredStudent(ctx, data.HomeIndex)
	if !found {
		return packetAck, types.ErrStudentNotPresent
	} else {
		utilfunc.PrintLogs("OnRecvFinalErasmusDataPacket " + data.ErasmusRestrictedInfo)
		err = utilfunc.UpdateErasmusData(&searchedStudent, data.ErasmusRestrictedInfo)
		if err != nil {
			return packetAck, err
		}
		utilfunc.PrintLogs("OnRecvFinalErasmusDataPacket finish")
		k.SetStoredStudent(ctx, searchedStudent)

		err = utilfunc.GetConsumedGas("OnRecvFinalErasmusDataPacket DE", data.HomeIndex, ctx)
		if err != nil {
			return packetAck, err
		} else {
			utilfunc.GetTransactionStats("OnRecvFinalErasmusDataPacket sending ack", ctx, sizeInt, binArray)
			return packetAck, nil
		}
	}

}

// OnAcknowledgementFinalErasmusDataPacket responds to the the success or failure of a packet
// acknowledgement written on the receiving chain.
func (k Keeper) OnAcknowledgementFinalErasmusDataPacket(ctx sdk.Context, packet channeltypes.Packet, data types.FinalErasmusDataPacketData, ack channeltypes.Acknowledgement) error {

	sizeInt := packet.Size()
	binArray, err := data.GetBytes()
	if err != nil {
		return err
	}
	utilfunc.GetTransactionStats("OnAcknowledgementFinalErasmusDataPacket", ctx, sizeInt, binArray)

	switch dispatchedAck := ack.Response.(type) {
	case *channeltypes.Acknowledgement_Error:

		// TODO: failed acknowledgement logic
		_ = dispatchedAck.Error

		utilfunc.PrintLogs("OnAcknowledgementFinalErasmusDataPacket error " + dispatchedAck.Error)

		return nil
	case *channeltypes.Acknowledgement_Result:
		// Decode the packet acknowledgment
		var packetAck types.FinalErasmusDataPacketAck

		if err := types.ModuleCdc.UnmarshalJSON(dispatchedAck.Result, &packetAck); err != nil {
			// The counter-party module doesn't implement the correct acknowledgment format
			return errors.New("cannot unmarshal acknowledgment")
		}

		// TODO: successful acknowledgement logic
		utilfunc.PrintLogs("OnAcknowledgementFinalErasmusDataPacket success")

		err = utilfunc.GetConsumedGas("OnAcknowledgementFinalErasmusDataPacket DE", data.HomeIndex, ctx)
		if err != nil {
			return err
		} else {

			return nil

		}
	default:
		// The counter-party module doesn't implement the correct acknowledgment format
		return errors.New("invalid acknowledgment format")
	}
}

// OnTimeoutFinalErasmusDataPacket responds to the case where a packet has not been transmitted because of a timeout
func (k Keeper) OnTimeoutFinalErasmusDataPacket(ctx sdk.Context, packet channeltypes.Packet, data types.FinalErasmusDataPacketData) error {

	// TODO: packet timeout logic

	utilfunc.PrintLogs("OnTimeoutFinalErasmusDataPacket")

	return nil
}
