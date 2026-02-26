package keeper

import (
	"encoding/json"
	"errors"
	"strconv"

	"university_chain_de/x/universitychainde/types"
	"university_chain_de/x/universitychainde/utilfunc"

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
	utilfunc.GetTransactionStats("TransmitEndErasmusPeriodRequestPacket", details, ctx, sizeInt, packetBytes)

	if err := k.ChannelKeeper.SendPacket(ctx, channelCap, packet); err != nil {
		return err
	}

	return nil
}

// OnRecvEndErasmusPeriodRequestPacket processes packet reception
func (k Keeper) OnRecvEndErasmusPeriodRequestPacket(ctx sdk.Context, packet channeltypes.Packet, data types.EndErasmusPeriodRequestPacketData) (packetAck types.EndErasmusPeriodRequestPacketAck, err error) {

	sizeInt := packet.Size()
	binArray, err := data.GetBytes()
	if err != nil {
		return packetAck, err
	}
	utilfunc.GetTransactionStats("OnRecvEndErasmusPeriodRequestPacket", "", ctx, sizeInt, binArray)

	// validate packet data upon receiving
	if err := data.ValidateBasic(); err != nil {
		return packetAck, err
	}

	// Packet reception logic

	utilfunc.PrintLogs("OnRecvEndErasmusPeriodRequestPacket", ctx)

	searchedStudent, found := k.GetStoredStudent(ctx, data.ForeignIndex)
	if !found {
		return packetAck, types.ErrStudentNotPresent
	} else {
		searchedStudent.ErasmusData.ErasmusStudent = "Incoming completed"
		err = utilfunc.ConcludeErasmusFlag(ctx, &searchedStudent)
		if err != nil {
			return k.ErrorHandlingEndErasmusAck(ctx, data.ForeignIndex)
		} else {
			k.SetStoredStudent(ctx, searchedStudent)
			stringIndex := data.Index
			data_res, err := utilfunc.GetErasmusExamsResults(searchedStudent)
			if err != nil {
				utilfunc.PrintLogs("SendErasmusStudent "+err.Error(), ctx)
				return k.ErrorHandlingEndErasmusAck(ctx, data.ForeignIndex)
			}
			packetAck.ErasmusRestrictedInfo = data_res

			err = utilfunc.GetConsumedGas("OnRecvEndErasmusPeriodRequestPacket", stringIndex, ctx)
			if err != nil {
				return k.ErrorHandlingEndErasmusAck(ctx, data.ForeignIndex)
			} else {
				packetAckBytes, err := types.ModuleCdc.MarshalJSON(&packetAck)
				if err != nil {
					return k.ErrorHandlingEndErasmusAck(ctx, data.ForeignIndex)
				}
				sizeInt := len(packetAckBytes)
				utilfunc.GetTransactionStats("OnRecvEndErasmusPeriodRequestPacket DE sending ack", "", ctx, sizeInt, binArray)
				return packetAck, nil
			}

		}
	}

}

// OnAcknowledgementEndErasmusPeriodRequestPacket responds to the the success or failure of a packet
// acknowledgement written on the receiving chain.
func (k Keeper) OnAcknowledgementEndErasmusPeriodRequestPacket(ctx sdk.Context, packet channeltypes.Packet, data types.EndErasmusPeriodRequestPacketData, ack channeltypes.Acknowledgement) error {

	switch dispatchedAck := ack.Response.(type) {
	case *channeltypes.Acknowledgement_Error:

		// Failed acknowledgement logic
		err := k.RevertEndErasmus(ctx, data.ForeignIndex)
		if err != nil {
			return err
		}

		utilfunc.PrintLogs("OnAcknowledgementEndErasmusPeriodRequestPacket error "+dispatchedAck.Error, ctx)

		return nil
	case *channeltypes.Acknowledgement_Result:
		// Decode the packet acknowledgment
		var packetAck types.EndErasmusPeriodRequestPacketAck

		sizeInt := len(dispatchedAck.Result)
		binArray, err := data.GetBytes()
		if err != nil {
			return err
		}
		utilfunc.GetTransactionStats("OnAcknowledgementEndErasmusPeriodRequestPacket", "", ctx, sizeInt, binArray)

		if err := types.ModuleCdc.UnmarshalJSON(dispatchedAck.Result, &packetAck); err != nil {
			// The counter-party module doesn't implement the correct acknowledgment format
			return errors.New("cannot unmarshal acknowledgment")
		}

		// Successful acknowledgement logic

		var result map[string]interface{}
		err = json.Unmarshal([]byte(packetAck.ErasmusRestrictedInfo), &result)
		if err != nil {
			return err
		}

		packetID, found := result["p_id"].(string)
		if !found {
			utilfunc.PrintLogs("OnAcknowledgementEndErasmusPeriodRequestPacket packetid not found", ctx)
			return nil
		}

		switch packetID {

		case "-1":

			utilfunc.PrintLogs("OnAcknowledgementEndErasmusPeriodRequestPacket case -1", ctx)

			var abortPacket utilfunc.AbortOperationPacket
			err = json.Unmarshal([]byte(packetAck.ErasmusRestrictedInfo), &abortPacket)
			if err != nil {
				return err
			} else {

				err = k.RevertEndErasmus(ctx, data.ForeignIndex)
				if err != nil {
					return err
				}

				return nil

			}
		default:

			utilfunc.PrintLogs("OnAcknowledgementEndErasmusPeriodRequestPacket success", ctx)

			packetHash := utilfunc.Hash(binArray)
			err = utilfunc.GetConsumedGas("OnAcknowledgementEndErasmusPeriodRequestPacket IT", strconv.FormatInt(int64(packetHash), 10), ctx)
			if err != nil {
				return err
			}

			return nil
		}

	default:
		// The counter-party module doesn't implement the correct acknowledgment format
		return errors.New("invalid acknowledgment format")
	}
}

// OnTimeoutEndErasmusPeriodRequestPacket responds to the case where a packet has not been transmitted because of a timeout
func (k Keeper) OnTimeoutEndErasmusPeriodRequestPacket(ctx sdk.Context, packet channeltypes.Packet, data types.EndErasmusPeriodRequestPacketData) error {

	// Packet timeout logic

	err := k.RevertEndErasmus(ctx, data.ForeignIndex)
	if err != nil {
		return err
	}

	utilfunc.PrintLogs("OnTimeoutEndErasmusPeriodRequestPacket", ctx)

	return nil
}
