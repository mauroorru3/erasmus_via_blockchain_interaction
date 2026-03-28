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

// TransmitErasmusRestictedDataPacket transmits the packet over IBC with the specified source port and source channel
func (k Keeper) TransmitErasmusRestictedDataPacket(
	ctx sdk.Context,
	packetData types.ErasmusRestictedDataPacketData,
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
	utilfunc.GetTransactionStats("TransmitErasmusRestictedDataPacket", details, ctx, sizeInt, packetBytes)

	if err := k.ChannelKeeper.SendPacket(ctx, channelCap, packet); err != nil {
		return err
	}

	return nil
}

// OnRecvErasmusRestictedDataPacket processes packet reception
func (k Keeper) OnRecvErasmusRestictedDataPacket(ctx sdk.Context, packet channeltypes.Packet, data types.ErasmusRestictedDataPacketData) (packetAck types.ErasmusRestictedDataPacketAck, err error) {

	sizeInt := packet.Size()
	binArray, err := data.GetBytes()
	if err != nil {
		return k.HandleAbortAckStartErasmus(ctx, data.ErasmusRestrictedInfo, packet)
	}
	utilfunc.GetTransactionStats("OnRecvErasmusRestictedDataPacket", "", ctx, sizeInt, binArray)

	// validate packet data upon receiving
	if err := data.ValidateBasic(); err != nil {
		return k.HandleAbortAckStartErasmus(ctx, data.ErasmusRestrictedInfo, packet)
	}

	utilfunc.PrintLogs("OnRecvErasmusRestictedDataPacket", ctx)

	utilfunc.PrintData(data.String(), ctx)

	var result map[string]interface{}
	err = json.Unmarshal([]byte(data.ErasmusRestrictedInfo), &result)
	if err != nil {
		return k.HandleAbortAckStartErasmus(ctx, data.ErasmusRestrictedInfo, packet)
	}

	utilfunc.PrintLogs("OnRecvErasmusRestictedDataPacket res "+data.ErasmusRestrictedInfo, ctx)

	uniStr, found := result["f_uni"].(string)
	if !found {
		return k.HandleAbortAckStartErasmus(ctx, data.ErasmusRestrictedInfo, packet)
	} else {

		uniInfo, found := k.GetUniversities(ctx, uniStr)
		if !found {
			return k.HandleAbortAckStartErasmus(ctx, data.ErasmusRestrictedInfo, packet)
		} else {

			utilfunc.PrintLogs("OnRecvErasmusRestictedDataPacket res dentro "+data.ErasmusRestrictedInfo, ctx)

			if utilfunc.TestAckErrorStartErasmus(ctx, false, 1, data.ErasmusRestrictedInfo) {

				return k.HandleAbortAckStartErasmus(ctx, data.ErasmusRestrictedInfo, packet)

			} else {

				var packet_to_send types.ErasmusRestictedDataPacketData

				packet_to_send.ErasmusRestrictedInfo = data.ErasmusRestrictedInfo

				err := k.TransmitErasmusRestictedDataPacket(
					ctx,
					packet_to_send,
					uniInfo.Port,
					uniInfo.ChannelID,
					clienttypes.ZeroHeight(),
					timeoutTimestamp,
					"")

				if err != nil {
					return k.HandleAbortAckStartErasmus(ctx, data.ErasmusRestrictedInfo, packet)
				} else {
					packetHash := utilfunc.Hash(binArray)
					err = utilfunc.GetConsumedGas("OnRecvErasmusRestictedDataPacket Hub", strconv.FormatInt(int64(packetHash), 10), ctx)
					if err != nil {
						return packetAck, err
					} else {
						packetAckBytes, err := types.ModuleCdc.MarshalJSON(&packetAck)
						if err != nil {
							return k.HandleAbortAckStartErasmus(ctx, data.ErasmusRestrictedInfo, packet)
						}
						sizeInt := len(packetAckBytes)
						utilfunc.GetTransactionStats("OnRecvErasmusRestictedDataPacket sending ack", "", ctx, sizeInt, binArray)

						return utilfunc.HandleSuccessAckStartErasmus(ctx)

					}

				}
			}

		}

	}

}

// OnAcknowledgementErasmusRestictedDataPacket responds to the the success or failure of a packet
// acknowledgement written on the receiving chain.
func (k Keeper) OnAcknowledgementErasmusRestictedDataPacket(ctx sdk.Context, packet channeltypes.Packet, data types.ErasmusRestictedDataPacketData, ack channeltypes.Acknowledgement) error {

	switch dispatchedAck := ack.Response.(type) {
	case *channeltypes.Acknowledgement_Error:

		// Failed acknowledgement logic
		_ = dispatchedAck.Error
		utilfunc.PrintLogs("OnAcknowledgementErasmusRestictedDataPacket error "+dispatchedAck.Error, ctx)

		return nil

	case *channeltypes.Acknowledgement_Result:
		// Decode the packet acknowledgment
		var packetAck types.ErasmusRestictedDataPacketAck

		sizeInt := len(dispatchedAck.Result)
		binArray, err := data.GetBytes()
		if err != nil {
			return err
		}
		utilfunc.GetTransactionStats("OnAcknowledgementErasmusRestictedDataPacket", "", ctx, sizeInt, binArray)

		if err := types.ModuleCdc.UnmarshalJSON(dispatchedAck.Result, &packetAck); err != nil {
			// The counter-party module doesn't implement the correct acknowledgment format
			utilfunc.PrintLogs("OnAcknowledgementErasmusRestictedDataPacket "+err.Error(), ctx)
			return errors.New("cannot unmarshal acknowledgment " + err.Error())
		}

		// Successful acknowledgement logic
		utilfunc.PrintLogs("OnAcknowledgementErasmusRestictedDataPacket", ctx)

		if packetAck.ErasmusRestrictedInfo != "" {

			utilfunc.PrintLogs("OnAcknowledgementErasmusRestictedDataPacket packetAck.ErasmusRestrictedInfo != nil", ctx)
			utilfunc.PrintLogs("OnAcknowledgementErasmusRestictedDataPacket ack "+packetAck.ErasmusRestrictedInfo, ctx)
			utilfunc.PrintLogs("OnAcknowledgementErasmusRestictedDataPacket previosu packet sent "+data.ErasmusRestrictedInfo, ctx)

			var result map[string]interface{}
			err = json.Unmarshal([]byte(packetAck.ErasmusRestrictedInfo), &result)
			if err != nil {
				utilfunc.PrintLogs("OnAcknowledgementErasmusRestictedDataPacket unmarshal packetAck.ErasmusRestrictedInfo", ctx)
				return err
			}

			packetID, found := result["p_id"].(string)
			if !found {
				utilfunc.PrintLogs("OnAcknowledgementErasmusRestictedDataPacket p_id not found", ctx)
				return errors.New("p_id not found or not a string")
			}

			switch packetID {

			case "-1": // error in the index packet of the start erasmus

				utilfunc.PrintLogs("OnAcknowledgementErasmusRestictedDataPacket case -1", ctx)

				var homeIndexPacket utilfunc.StudentInfoRestrictedHomeIndexPacket
				err = json.Unmarshal([]byte(data.ErasmusRestrictedInfo), &homeIndexPacket)
				if err != nil {
					return err
				}

				uniInfo, found := k.GetUniversities(ctx, homeIndexPacket.HomeUniversity)
				if !found {
					utilfunc.PrintLogs("OnAcknowledgementErasmusRestictedDataPacket case -1 "+types.ErrWrongNameUniversity.Error(), ctx)
					//return types.ErrWrongNameUniversity
				} else {

					utilfunc.PrintLogs("OnAcknowledgementErasmusRestictedDataPacket case - 1", ctx)
					var abortData utilfunc.AbortOperationPacket

					utilfunc.PrintLogs("OnAcknowledgementErasmusRestictedDataPacket case -1 data "+data.ErasmusRestrictedInfo, ctx)
					utilfunc.PrintLogs("OnAcknowledgementErasmusRestictedDataPacket case -1 packetack "+packetAck.ErasmusRestrictedInfo, ctx)

					abortData.PacketID = "-1"
					abortData.HomeIndex = homeIndexPacket.HomeIndex
					abortData.HomeUniversity = homeIndexPacket.HomeUniversity
					abortData.ForeignUniversity = homeIndexPacket.ForeignUniversity

					resultByteJSON, err := json.Marshal(abortData)
					if err != nil {
						return err
					}

					var packet_to_send types.ErasmusRestictedDataPacketData
					packet_to_send.ErasmusRestrictedInfo = string(resultByteJSON)

					utilfunc.PrintLogs("OnAcknowledgementErasmusRestictedDataPacket case -1 packettosend "+packet_to_send.ErasmusRestrictedInfo, ctx)

					err = k.TransmitErasmusRestictedDataPacket(ctx,
						packet_to_send,
						uniInfo.Port,
						uniInfo.ChannelID,
						clienttypes.ZeroHeight(),
						timeoutTimestamp,
						" OnAcknowledgementErasmusRestictedDataPacket case -1")

					if err != nil {
						utilfunc.PrintLogs("OnAcknowledgementErasmusRestictedDataPacket case -1 error "+err.Error(), ctx)
						return err

					}
					return nil
				}

			case "-4": // error in the other 10 packets of the start erasmus

				utilfunc.PrintLogs("OnAcknowledgementErasmusRestictedDataPacket case -4", ctx)

				var abortPacket utilfunc.AbortOperationPacket
				err = json.Unmarshal([]byte(packetAck.ErasmusRestrictedInfo), &abortPacket)
				if err != nil {
					return err
				}

				uniInfo, found := k.GetUniversities(ctx, abortPacket.HomeUniversity)
				if !found {
					utilfunc.PrintLogs("OnAcknowledgementErasmusRestictedDataPacket case -4 "+types.ErrWrongNameUniversity.Error(), ctx)
					return types.ErrWrongNameUniversity
				} else {

					utilfunc.PrintLogs("OnAcknowledgementErasmusRestictedDataPacket case - 4", ctx)
					var abortData utilfunc.AbortOperationPacket

					utilfunc.PrintLogs("OnAcknowledgementErasmusRestictedDataPacket case -4 data "+data.ErasmusRestrictedInfo, ctx)
					utilfunc.PrintLogs("OnAcknowledgementErasmusRestictedDataPacket case -4 packetack "+packetAck.ErasmusRestrictedInfo, ctx)

					abortData.PacketID = abortPacket.PacketID
					abortData.HomeIndex = abortPacket.HomeIndex
					abortData.HomeUniversity = abortPacket.HomeUniversity
					abortData.ForeignUniversity = abortPacket.ForeignUniversity
					abortData.ForeignIndex = abortPacket.ForeignIndex

					resultByteJSON, err := json.Marshal(abortData)
					if err != nil {
						return err
					}

					var packet_to_send types.ErasmusRestictedDataPacketData
					packet_to_send.ErasmusRestrictedInfo = string(resultByteJSON)

					utilfunc.PrintLogs("OnAcknowledgementErasmusRestictedDataPacket case -4 packettosend "+packet_to_send.ErasmusRestrictedInfo, ctx)

					err = k.TransmitErasmusRestictedDataPacket(ctx,
						packet_to_send,
						uniInfo.Port,
						uniInfo.ChannelID,
						clienttypes.ZeroHeight(),
						timeoutTimestamp,
						" OnAcknowledgementErasmusRestictedDataPacket case -4")

					if err != nil {
						utilfunc.PrintLogs("OnAcknowledgementErasmusRestictedDataPacket case -4 error "+err.Error(), ctx)
						return err

					}
					return nil
				}

			case "20": // index packet of the start erasmus

				utilfunc.PrintLogs("OnAcknowledgementErasmusRestictedDataPacket case 20", ctx)

				var homeIndexPacket utilfunc.StudentInfoRestrictedAnswerPacket
				utilfunc.PrintLogs("OnAcknowledgementErasmusRestictedDataPacket case 20 - 1", ctx)
				err = json.Unmarshal([]byte(packetAck.ErasmusRestrictedInfo), &homeIndexPacket)
				utilfunc.PrintLogs("OnAcknowledgementErasmusRestictedDataPacket case 20 - 2", ctx)
				if err != nil {
					utilfunc.PrintLogs("OnAcknowledgementErasmusRestictedDataPacket err json.Unmarshal", ctx)
					return err

				}

				utilfunc.PrintLogs("OnAcknowledgementErasmusRestictedDataPacket pre unmarshal data.ErasmusRestrictedInfo", ctx)
				var packetSent map[string]string
				err = json.Unmarshal([]byte(data.ErasmusRestrictedInfo), &packetSent)
				if err != nil {
					utilfunc.PrintLogs("OnAcknowledgementErasmusRestictedDataPacket err unmarshal data.ErasmusRestrictedInfo", ctx)
					return err

				}

				utilfunc.PrintLogs("OnAcknowledgementErasmusRestictedDataPacket pre uni", ctx)
				uniStr := packetSent["h_uni"]

				uniInfo, found := k.GetUniversities(ctx, uniStr)
				if !found {
					utilfunc.PrintLogs("OnAcknowledgementErasmusRestictedDataPacket case 20 "+types.ErrWrongNameUniversity.Error(), ctx)
					return types.ErrWrongNameUniversity
				} else {

					utilfunc.PrintLogs("OnAcknowledgementErasmusRestictedDataPacket case 20 - 1", ctx)
					var packet_to_send types.ErasmusIndexPacketData

					utilfunc.PrintLogs("OnAcknowledgementErasmusRestictedDataPacket case 20 data "+data.ErasmusRestrictedInfo, ctx)
					utilfunc.PrintLogs("OnAcknowledgementErasmusRestictedDataPacket case 20 packetack "+packetAck.ErasmusRestrictedInfo, ctx)

					packet_to_send.ForeignIndex = homeIndexPacket.ForeignIndex
					packet_to_send.Index = packetSent["h_id"]

					err := k.TransmitErasmusIndexPacket(ctx,
						packet_to_send,
						uniInfo.Port,
						uniInfo.ChannelID,
						clienttypes.ZeroHeight(),
						timeoutTimestamp,
						" OnAcknowledgementErasmusRestictedDataPacket case 20")

					if err != nil {
						utilfunc.PrintLogs("OnAcknowledgementErasmusRestictedDataPacket case 20 error "+err.Error(), ctx)
						return err

					} else {

						utilfunc.PrintLogs("OnAcknowledgementErasmusRestictedDataPacket case 20 packet sent", ctx)

						err = utilfunc.GetConsumedGas("OnAcknowledgementErasmusRestictedDataPacket case 20 Hub", packet_to_send.Index, ctx)
						if err != nil {
							utilfunc.PrintLogs("OnAcknowledgementErasmusRestictedDataPacket case 20 error "+err.Error(), ctx)
							return err
						} else {

							return nil
						}
					}
				}

			case "21": // start erasmus confirmation

				utilfunc.PrintLogs("OnAcknowledgementErasmusRestictedDataPacket case 21", ctx)

				var okPacket utilfunc.SuccessPacket
				err = json.Unmarshal([]byte(packetAck.ErasmusRestrictedInfo), &okPacket)
				if err != nil {
					return err

				}

				var successData utilfunc.SuccessPacket

				utilfunc.PrintLogs("OnAcknowledgementErasmusRestictedDataPacket case 21 data "+data.ErasmusRestrictedInfo, ctx)
				utilfunc.PrintLogs("OnAcknowledgementErasmusRestictedDataPacket case 21 packetack "+packetAck.ErasmusRestrictedInfo, ctx)

				uniInfo, found := k.GetUniversities(ctx, okPacket.HomeUniversity)
				if !found {
					utilfunc.PrintLogs("OnAcknowledgementErasmusRestictedDataPacket case 21 "+types.ErrWrongNameUniversity.Error(), ctx)
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
						" OnAcknowledgementErasmusRestictedDataPacket case 21")

					if err != nil {
						utilfunc.PrintLogs("OnAcknowledgementErasmusRestictedDataPacket case 21 error "+err.Error(), ctx)
						return err

					}
					return nil
				}

			case "22": //extend erasmus confirmation

				utilfunc.PrintLogs("OnAcknowledgementErasmusRestictedDataPacket case 22", ctx)

				var okPacket utilfunc.SuccessPacket
				err = json.Unmarshal([]byte(packetAck.ErasmusRestrictedInfo), &okPacket)
				if err != nil {
					return err

				}

				var successData utilfunc.SuccessPacket

				utilfunc.PrintLogs("OnAcknowledgementErasmusRestictedDataPacket case 22 data "+data.ErasmusRestrictedInfo, ctx)
				utilfunc.PrintLogs("OnAcknowledgementErasmusRestictedDataPacket case 22 packetack "+packetAck.ErasmusRestrictedInfo, ctx)

				uniInfo, found := k.GetUniversities(ctx, okPacket.HomeUniversity)
				if !found {
					utilfunc.PrintLogs("OnAcknowledgementErasmusRestictedDataPacket case 22 "+types.ErrWrongNameUniversity.Error(), ctx)
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
						" OnAcknowledgementErasmusRestictedDataPacket case 22")

					if err != nil {
						utilfunc.PrintLogs("OnAcknowledgementErasmusRestictedDataPacket case 22 error "+err.Error(), ctx)
						return err

					}

					return nil
				}

			default:
				return nil

			}

		}

		return nil
	default:
		// The counter-party module doesn't implement the correct acknowledgment format
		return errors.New("invalid acknowledgment format")
	}
}

// OnTimeoutErasmusRestictedDataPacket responds to the case where a packet has not been transmitted because of a timeout
func (k Keeper) OnTimeoutErasmusRestictedDataPacket(ctx sdk.Context, packet channeltypes.Packet, data types.ErasmusRestictedDataPacketData) error {

	// Packet timeout logic
	utilfunc.PrintLogs("OnTimeoutErasmusRestictedDataPacket", ctx)
	return nil

}
