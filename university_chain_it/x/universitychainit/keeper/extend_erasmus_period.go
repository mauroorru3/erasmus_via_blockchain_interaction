package keeper

import (
	"errors"

	"university_chain_it/x/universitychainit/types"
	"university_chain_it/x/universitychainit/utilfunc"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	clienttypes "github.com/cosmos/ibc-go/v3/modules/core/02-client/types"
	channeltypes "github.com/cosmos/ibc-go/v3/modules/core/04-channel/types"
	host "github.com/cosmos/ibc-go/v3/modules/core/24-host"
)

// TransmitExtendErasmusPeriodPacket transmits the packet over IBC with the specified source port and source channel
func (k Keeper) TransmitExtendErasmusPeriodPacket(
	ctx sdk.Context,
	packetData types.ExtendErasmusPeriodPacketData,
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
	utilfunc.GetTransactionStats("TransmitExtendErasmusPeriodPacket", details, ctx, sizeInt, packetBytes)

	if err := k.ChannelKeeper.SendPacket(ctx, channelCap, packet); err != nil {
		return err
	}

	return nil
}

// OnRecvExtendErasmusPeriodPacket processes packet reception
func (k Keeper) OnRecvExtendErasmusPeriodPacket(ctx sdk.Context, packet channeltypes.Packet, data types.ExtendErasmusPeriodPacketData) (packetAck types.ExtendErasmusPeriodPacketAck, err error) {

	sizeInt := packet.Size()
	binArray, err := data.GetBytes()
	if err != nil {
		return packetAck, err
	}
	utilfunc.GetTransactionStats("OnRecvExtendErasmusPeriodPacket", "", ctx, sizeInt, binArray)

	utilfunc.PrintLogs("OnRecvExtendErasmusPeriodPacket")

	searchedStudent, found := k.GetStoredStudent(ctx, data.ForeignIndex)
	if !found {
		utilfunc.PrintLogs("OnRecvExtendErasmusPeriodPacket " + types.ErrStudentNotPresent.Error())
		return packetAck, types.ErrStudentNotPresent
	} else {

		err := utilfunc.ExtendErasmusForeignStudent(ctx, data.DurationInMonths, data.FinalDate, &searchedStudent)
		if err != nil {
			utilfunc.PrintLogs("OnRecvExtendErasmusPeriodPacket " + err.Error())
			return packetAck, err
		} else {
			k.SetStoredStudent(ctx, searchedStudent)
			utilfunc.PrintLogs("OnRecvExtendErasmusPeriodPacket ack sent")

			stringIndex, err := utilfunc.GetForeignIndex(searchedStudent)
			if err != nil {
				return packetAck, err
			} else {
				err = utilfunc.GetConsumedGas("OnRecvExtendErasmusPeriodPacket", stringIndex, ctx)
				if err != nil {
					return packetAck, err
				} else {
					packetAckBytes, err := types.ModuleCdc.MarshalJSON(&packetAck)
					if err != nil {
						return packetAck, err
					}
					sizeInt := len(packetAckBytes)
					utilfunc.GetTransactionStats("OnRecvExtendErasmusPeriodPacket IT sending ack", "", ctx, sizeInt, binArray)
					return packetAck, nil
				}

			}
		}
	}
}

// OnAcknowledgementExtendErasmusPeriodPacket responds to the the success or failure of a packet
// acknowledgement written on the receiving chain.
func (k Keeper) OnAcknowledgementExtendErasmusPeriodPacket(ctx sdk.Context, packet channeltypes.Packet, data types.ExtendErasmusPeriodPacketData, ack channeltypes.Acknowledgement) error {

	switch dispatchedAck := ack.Response.(type) {
	case *channeltypes.Acknowledgement_Error:

		// TODO: failed acknowledgement logic
		_ = dispatchedAck.Error
		utilfunc.PrintLogs("OnAcknowledgementExtendErasmusPeriodPacket error " + dispatchedAck.Error)

		return nil
	case *channeltypes.Acknowledgement_Result:
		// Decode the packet acknowledgment
		var packetAck types.ExtendErasmusPeriodPacketAck

		sizeInt := len(dispatchedAck.Result)
		binArray, err := data.GetBytes()
		if err != nil {
			return err
		}
		utilfunc.GetTransactionStats("OnAcknowledgementExtendErasmusPeriodPacket", "", ctx, sizeInt, binArray)

		if err := types.ModuleCdc.UnmarshalJSON(dispatchedAck.Result, &packetAck); err != nil {
			// The counter-party module doesn't implement the correct acknowledgment format
			return errors.New("cannot unmarshal acknowledgment")
		}

		// TODO: successful acknowledgement logic
		utilfunc.PrintLogs("OnAcknowledgementExtendErasmusPeriodPacket success")

		searchedStudent, found := k.GetStoredStudent(ctx, data.ForeignIndex)
		if !found {
			utilfunc.PrintLogs("OnAcknowledgementExtendErasmusPeriodPacket " + types.ErrStudentNotPresent.Error())
			return types.ErrStudentNotPresent
		} else {
			stringIndex, err := utilfunc.GetForeignIndex(searchedStudent)
			if err != nil {
				return err
			} else {
				err = utilfunc.GetConsumedGas("OnAcknowledgementExtendErasmusPeriodPacket IT", stringIndex, ctx)
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

// OnTimeoutExtendErasmusPeriodPacket responds to the case where a packet has not been transmitted because of a timeout
func (k Keeper) OnTimeoutExtendErasmusPeriodPacket(ctx sdk.Context, packet channeltypes.Packet, data types.ExtendErasmusPeriodPacketData) error {

	// TODO: packet timeout logic
	utilfunc.PrintLogs("OnTimeoutExtendErasmusPeriodPacket")

	return nil
}
