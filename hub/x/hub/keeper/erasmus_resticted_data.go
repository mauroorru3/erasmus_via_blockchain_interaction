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
	utilfunc.GetTransactionStats("TransmitErasmusRestictedDataPacket", ctx, sizeInt, packetBytes)

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
		return packetAck, err
	}
	utilfunc.GetTransactionStats("OnRecvErasmusRestictedDataPacket", ctx, sizeInt, binArray)

	// validate packet data upon receiving
	if err := data.ValidateBasic(); err != nil {
		return packetAck, err
	}

	utilfunc.PrintLogs("OnRecvErasmusRestictedDataPacket")

	utilfunc.PrintData(data.String())

	var result map[string]interface{}
	err = json.Unmarshal([]byte(data.ErasmusRestrictedInfo), &result)
	if err != nil {
		return packetAck, err
	}

	uniStr, found := result["f_uni"].(string)
	if !found {
		return packetAck, sdkerrors.Error{}
	} else {

		uniInfo, found := k.GetUniversities(ctx, uniStr)
		if !found {
			return packetAck, types.ErrWrongNameUniversity
		} else {

			var packet_to_send types.ErasmusRestictedDataPacketData

			packet_to_send.ErasmusRestrictedInfo = data.ErasmusRestrictedInfo

			err := k.TransmitErasmusRestictedDataPacket(
				ctx,
				packet_to_send,
				uniInfo.Port,
				uniInfo.ChannelID,
				clienttypes.ZeroHeight(),
				timeoutTimestamp)

			if err != nil {
				return packetAck, err
			} else {
				packetHash := utilfunc.Hash(binArray)
				err = utilfunc.GetConsumedGas("OnRecvErasmusRestictedDataPacket Hub", strconv.FormatInt(int64(packetHash), 10), ctx)
				if err != nil {
					return packetAck, err
				} else {
					utilfunc.GetTransactionStats("OnRecvErasmusRestictedDataPacket sending ack", ctx, sizeInt, binArray)
					return packetAck, nil
				}

			}
		}

	}

}

// OnAcknowledgementErasmusRestictedDataPacket responds to the the success or failure of a packet
// acknowledgement written on the receiving chain.
func (k Keeper) OnAcknowledgementErasmusRestictedDataPacket(ctx sdk.Context, packet channeltypes.Packet, data types.ErasmusRestictedDataPacketData, ack channeltypes.Acknowledgement) error {

	sizeInt := packet.Size()
	binArray, err := data.GetBytes()
	if err != nil {
		return err
	}
	utilfunc.GetTransactionStats("OnAcknowledgementErasmusRestictedDataPacket", ctx, sizeInt, binArray)

	switch dispatchedAck := ack.Response.(type) {
	case *channeltypes.Acknowledgement_Error:

		// TODO: failed acknowledgement logic
		_ = dispatchedAck.Error
		utilfunc.PrintLogs("OnAcknowledgementErasmusRestictedDataPacket error " + dispatchedAck.Error)

		return nil
	case *channeltypes.Acknowledgement_Result:
		// Decode the packet acknowledgment
		var packetAck types.ErasmusRestictedDataPacketAck

		if err := types.ModuleCdc.UnmarshalJSON(dispatchedAck.Result, &packetAck); err != nil {
			// The counter-party module doesn't implement the correct acknowledgment format
			utilfunc.PrintLogs("OnAcknowledgementErasmusRestictedDataPacket " + err.Error())
			return errors.New("cannot unmarshal acknowledgment " + err.Error())
		}

		// TODO: successful acknowledgement logic
		utilfunc.PrintLogs("OnAcknowledgementErasmusRestictedDataPacket")

		if packetAck.ErasmusRestrictedInfo != "" {

			var packetSent map[string]string
			err := json.Unmarshal([]byte(data.ErasmusRestrictedInfo), &packetSent)
			if err != nil {
				return err
			}
			uniStr := packetSent["h_uni"]

			var result map[string]string
			err = json.Unmarshal([]byte(packetAck.ErasmusRestrictedInfo), &result)
			if err != nil {
				return err
			}

			uniInfo, found := k.GetUniversities(ctx, uniStr)
			if !found {
				utilfunc.PrintLogs("OnRecvErasmusRestictedDataPacket " + types.ErrWrongNameUniversity.Error())
				return types.ErrWrongNameUniversity
			} else {

				var packet_to_send types.ErasmusIndexPacketData

				utilfunc.PrintData("OnAcknowledgementErasmusRestictedDataPacket " + data.String())

				packet_to_send.ForeignIndex = result["f_id"]
				packet_to_send.Index = packetSent["h_id"]

				utilfunc.PrintData("OnAcknowledgementErasmusRestictedDataPacket " + packet_to_send.String())

				err := k.TransmitErasmusIndexPacket(ctx,
					packet_to_send,
					uniInfo.Port,
					uniInfo.ChannelID,
					clienttypes.ZeroHeight(),
					timeoutTimestamp)

				if err != nil {
					utilfunc.PrintLogs("OnAcknowledgementErasmusRestictedDataPacket error " + err.Error())
					return err
				} else {

					utilfunc.PrintLogs("OnAcknowledgementErasmusRestictedDataPacket packet sent")

					err = utilfunc.GetConsumedGas("OnAcknowledgementErasmusRestictedDataPacket Hub", packet_to_send.Index, ctx)
					if err != nil {
						utilfunc.PrintLogs("OnAcknowledgementErasmusRestictedDataPacket error " + err.Error())
						return err
					} else {

						return nil
					}

				}
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

	// TODO: packet timeout logic
	utilfunc.PrintLogs("OnAcknowledgementErasmusRestictedDataPacket")

	return nil
}
