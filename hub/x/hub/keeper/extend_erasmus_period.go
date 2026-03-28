package keeper

import (
	"encoding/json"
	"errors"
	"strconv"

	"hub/x/hub/types"
	"hub/x/hub/utilfunc"

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
		k.HandleAbortAckExtendErasmus(ctx, "", data.ForeignIndex, "", data.DestinationUniversityName)
	}
	utilfunc.GetTransactionStats("OnRecvExtendErasmusPeriodPacket", "", ctx, sizeInt, binArray)

	utilfunc.PrintLogs("OnRecvExtendErasmusPeriodPacket", ctx)
	utilfunc.PrintLogs("OnRecvExtendErasmusPeriodPacket packet received "+data.String(), ctx)

	uniInfo, found := k.GetUniversities(ctx, data.DestinationUniversityName)
	if !found {
		utilfunc.PrintLogs("OnRecvExtendErasmusPeriodPacket "+types.ErrWrongNameUniversity.Error(), ctx)
		return k.HandleAbortAckExtendErasmus(ctx, "", data.ForeignIndex, "", data.DestinationUniversityName)
	} else {

		if utilfunc.TestAckError(ctx, false) {

			return k.HandleAbortAckExtendErasmus(ctx, "", data.ForeignIndex, "", data.DestinationUniversityName)

		} else {

			// Transmit the packet
			err = k.TransmitExtendErasmusPeriodPacket(
				ctx,
				data,
				uniInfo.Port,
				uniInfo.ChannelID,
				clienttypes.ZeroHeight(),
				timeoutTimestamp,
				"")

			if err != nil {
				utilfunc.PrintLogs("OnRecvExtendErasmusPeriodPacket "+err.Error(), ctx)
				return k.HandleAbortAckExtendErasmus(ctx, "", data.ForeignIndex, "", data.DestinationUniversityName)
			} else {
				utilfunc.PrintLogs("OnRecvExtendErasmusPeriodPacket packet sent", ctx)

				packetHash := utilfunc.Hash(binArray)
				err = utilfunc.GetConsumedGas("OnRecvExtendErasmusPeriodPacket", strconv.FormatInt(int64(packetHash), 10), ctx)
				if err != nil {
					return k.HandleAbortAckExtendErasmus(ctx, "", data.ForeignIndex, "", data.DestinationUniversityName)
				} else {
					packetAckBytes, err := types.ModuleCdc.MarshalJSON(&packetAck)
					if err != nil {
						return k.HandleAbortAckExtendErasmus(ctx, "", data.ForeignIndex, "", data.DestinationUniversityName)
					}
					sizeInt := len(packetAckBytes)
					utilfunc.GetTransactionStats("OnRecvExtendErasmusPeriodPacket Hub sending ack", "", ctx, sizeInt, binArray)
					return utilfunc.HandleSuccessAckExtendErasmus(ctx)
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

		// Failed acknowledgement logic
		utilfunc.PrintLogs("OnAcknowledgementExtendErasmusPeriodPacket error "+dispatchedAck.Error, ctx)
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

		// Successful acknowledgement logic

		if packetAck.ErasmusRestrictedInfo != "" {

			utilfunc.PrintLogs("OnAcknowledgementExtendErasmusPeriodPacket packetAck.ErasmusRestrictedInfo != nil", ctx)
			utilfunc.PrintLogs("OnAcknowledgementExtendErasmusPeriodPacket "+packetAck.ErasmusRestrictedInfo, ctx)

			var result map[string]interface{}
			err = json.Unmarshal([]byte(packetAck.ErasmusRestrictedInfo), &result)
			if err != nil {
				utilfunc.PrintLogs("OnAcknowledgementExtendErasmusPeriodPacket unmarshal packetAck.ErasmusRestrictedInfo", ctx)
				return err
			}

			packetID, found := result["p_id"].(string)
			if !found {
				utilfunc.PrintLogs("OnAcknowledgementExtendErasmusPeriodPacket p_id not found", ctx)
				return err
			}

			switch packetID {

			case "22": //extend erasmus confirmation

				utilfunc.PrintLogs("OnAcknowledgementExtendErasmusPeriodPacket case 22", ctx)

				var okPacket utilfunc.SuccessPacket
				err = json.Unmarshal([]byte(packetAck.ErasmusRestrictedInfo), &okPacket)
				if err != nil {
					return err

				}

				var successData utilfunc.SuccessPacket

				utilfunc.PrintLogs("OnAcknowledgementExtendErasmusPeriodPacket case 22 packetack "+packetAck.ErasmusRestrictedInfo, ctx)

				uniInfo, found := k.GetUniversities(ctx, okPacket.HomeUniversity)
				if !found {
					utilfunc.PrintLogs("OnAcknowledgementExtendErasmusPeriodPacket case 22 "+types.ErrWrongNameUniversity.Error(), ctx)
					return types.ErrWrongNameUniversity
				} else {

					successData.HomeIndex = okPacket.HomeIndex
					successData.HomeUniversity = okPacket.HomeUniversity
					successData.PacketID = okPacket.PacketID

					resultByteJSON, err := json.Marshal(successData)
					if err != nil {
						return err
					}

					var packet_to_send types.ErasmusRestictedDataPacketData
					packet_to_send.ErasmusRestrictedInfo = string(resultByteJSON)

					err = k.TransmitErasmusRestictedDataPacket(ctx,
						packet_to_send,
						uniInfo.Port,
						uniInfo.ChannelID,
						clienttypes.ZeroHeight(),
						timeoutTimestamp,
						" OnAcknowledgementExtendErasmusPeriodPacket case 22")

					if err != nil {
						utilfunc.PrintLogs("OnAcknowledgementExtendErasmusPeriodPacket case 22 error "+err.Error(), ctx)
						return err

					}
				}

			case "-3": // extend erasmus application error

				utilfunc.PrintLogs("OnAcknowledgementExtendErasmusPeriodPacket case -3", ctx)

				var abortPacket utilfunc.AbortOperationPacket
				err = json.Unmarshal([]byte(packetAck.ErasmusRestrictedInfo), &abortPacket)
				if err != nil {
					return err

				}

				var abortData utilfunc.AbortOperationPacket

				utilfunc.PrintLogs("OnAcknowledgementExtendErasmusPeriodPacket case -3 packetack "+packetAck.ErasmusRestrictedInfo, ctx)

				uniInfo, found := k.GetUniversities(ctx, abortPacket.HomeUniversity)
				if !found {
					utilfunc.PrintLogs("OnAcknowledgementExtendErasmusPeriodPacket case -3 "+types.ErrWrongNameUniversity.Error(), ctx)
					return types.ErrWrongNameUniversity
				} else {

					abortData.HomeIndex = abortPacket.HomeIndex
					abortData.HomeUniversity = abortPacket.HomeUniversity
					abortData.PacketID = abortPacket.PacketID
					abortData.ForeignUniversity = abortPacket.ForeignUniversity
					abortData.ForeignIndex = abortPacket.ForeignIndex
					resultByteJSON, err := json.Marshal(abortData)
					if err != nil {
						return err
					}

					var packet_to_send types.ErasmusRestictedDataPacketData
					packet_to_send.ErasmusRestrictedInfo = string(resultByteJSON)

					utilfunc.PrintLogs("OnAcknowledgementExtendErasmusPeriodPacket case -3 abortData "+string(resultByteJSON), ctx)

					err = k.TransmitErasmusRestictedDataPacket(ctx,
						packet_to_send,
						uniInfo.Port,
						uniInfo.ChannelID,
						clienttypes.ZeroHeight(),
						timeoutTimestamp,
						" OnAcknowledgementExtendErasmusPeriodPacket case -3")

					if err != nil {
						utilfunc.PrintLogs("OnAcknowledgementExtendErasmusPeriodPacket case -3 error "+err.Error(), ctx)
						return err

					}
				}

			default:
				return nil

			}
		}

		utilfunc.PrintLogs("OnAcknowledgementExtendErasmusPeriodPacket success", ctx)

		packetHash := utilfunc.Hash(binArray)
		err = utilfunc.GetConsumedGas("OnAcknowledgementExtendErasmusPeriodPacket Hub", strconv.FormatInt(int64(packetHash), 10), ctx)
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

// OnTimeoutExtendErasmusPeriodPacket responds to the case where a packet has not been transmitted because of a timeout
func (k Keeper) OnTimeoutExtendErasmusPeriodPacket(ctx sdk.Context, packet channeltypes.Packet, data types.ExtendErasmusPeriodPacketData) error {

	// Packet timeout logic

	utilfunc.PrintLogs("OnTimeoutExtendErasmusPeriodPacket", ctx)

	return nil
}
