package keeper

import (
	"encoding/json"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"hub/x/hub/types"
	"hub/x/hub/utilfunc"

	clienttypes "github.com/cosmos/ibc-go/v3/modules/core/02-client/types"
	channeltypes "github.com/cosmos/ibc-go/v3/modules/core/04-channel/types"
)

// ***Timeout and error management of the Start Erasmus operation***

// Function that creates the content of the packet that sends the abort operation request
// for that considering the outgoing student and the Start Erasmus function

func (k Keeper) CreateAbortOperationString(student_index string, student_foreign_index string, student_home_uni string, student_foreign_uni string) (abort_op_JSON string, err error) {

	var abort_op utilfunc.AbortOperationPacket

	abort_op.HomeIndex = student_index
	abort_op.ForeignIndex = student_foreign_index
	abort_op.ForeignUniversity = student_foreign_uni
	abort_op.HomeUniversity = student_home_uni
	abort_op.PacketID = "-1" // Value that identifies the packet related to the abort operation of the Start Erasmus operation

	resultByteJSON, err := json.Marshal(abort_op)
	if err != nil {
		return abort_op_JSON, err
	}

	abort_op_JSON = string(resultByteJSON)

	return abort_op_JSON, err
}

// Function that sends the abort operation packet

func (k Keeper) SendAbortPacketToPortChannel(ctx sdk.Context, data string, port_id string, channel_id string) (err error) {

	utilfunc.PrintLogs("SendAbortPacketToPortChannel", ctx)

	var packetToSend types.ErasmusRestictedDataPacketData
	packetToSend.ErasmusRestrictedInfo = data

	err = k.TransmitErasmusRestictedDataPacket(
		ctx,
		packetToSend,
		port_id,
		channel_id,
		clienttypes.ZeroHeight(),
		timeoutTimestamp,
		" SendAbortPacketToPortChannel",
	)

	if err != nil {
		utilfunc.PrintLogs("SendAbortPacketToPortChannel "+err.Error(), ctx)
		return err
	}
	return nil
}

// Function that sends the abort operation packet

func (k Keeper) SendAbortPacket(ctx sdk.Context, data string, packet channeltypes.Packet) (err error) {

	utilfunc.PrintLogs("SendAbortPacket", ctx)

	var packetToSend types.ErasmusRestictedDataPacketData
	packetToSend.ErasmusRestrictedInfo = data

	err = k.TransmitErasmusRestictedDataPacket(
		ctx,
		packetToSend,
		packet.DestinationPort,
		packet.DestinationChannel,
		clienttypes.ZeroHeight(),
		timeoutTimestamp,
		" SendAbortPacket",
	)

	if err != nil {
		utilfunc.PrintLogs("SendAbortPacket "+err.Error(), ctx)
		return err
	}
	return nil
}

// Function that constructs the ack contents and returns the ack related to the abort operation

func (k Keeper) HandleAbortAckStartErasmus(ctx sdk.Context, data string, packet channeltypes.Packet) (packetAck types.ErasmusRestictedDataPacketAck, err error) {

	utilfunc.PrintLogs("HandleAbortAck", ctx)

	var result map[string]interface{}
	err = json.Unmarshal([]byte(data), &result)
	if err != nil {
		return packetAck, err
	}

	student_index, _ := result["h_id"].(string)
	student_foreign_index, _ := result["f_id"].(string)
	student_foreign_uni, _ := result["f_uni"].(string)
	student_home_uni, _ := result["h_uni"].(string)

	new_data_info, err := k.CreateAbortOperationString(student_index, student_foreign_index, student_home_uni, student_foreign_uni)
	if err != nil {
		return packetAck, err
	} else {

		packetAck.ErasmusRestrictedInfo = new_data_info
		return packetAck, err
	}
}

// ***Timeout and error management of the End Erasmus and the EndErasmusBeforeDeadline operations***

// Function that allows the construction of the packet content that allows the abort
// operation related to the end of the Erasmus

func (k Keeper) CreateAbortOperationStringEndErasmus(student_index string, student_foreign_index string, student_home_uni string, student_foreign_uni string) (abort_op_JSON string, err error) {

	var abort_op utilfunc.AbortOperationPacket

	abort_op.HomeIndex = student_index
	abort_op.ForeignIndex = student_foreign_index
	abort_op.ForeignUniversity = student_foreign_uni
	abort_op.HomeUniversity = student_home_uni
	abort_op.PacketID = "-2"

	resultByteJSON, err := json.Marshal(abort_op)
	if err != nil {
		return abort_op_JSON, err
	}

	abort_op_JSON = string(resultByteJSON)

	return abort_op_JSON, err
}

// Function that constructs the ack contents and returns the ack related to the abort
// operation of the end erasmus

func (k Keeper) HandleAbortAckEndErasmus(ctx sdk.Context, student_index string, student_foreign_index string, student_home_uni string, student_foreign_uni string, packet channeltypes.Packet) (packetAck types.EndErasmusPeriodRequestPacketAck, err error) {

	utilfunc.PrintLogs("HandleAbortAckEndErasmus", ctx)

	new_data_info, err := k.CreateAbortOperationStringEndErasmus(student_index, student_foreign_index, student_home_uni, student_foreign_uni)
	if err != nil {
		return packetAck, err
	} else {

		packetAck.ErasmusRestrictedInfo = new_data_info
		return packetAck, err
	}
}

// ***Timeout and error management of the Extend Erasmus operation***

// Function that allows the construction of the packet content that allows the abort
// operation related to the extend of the Erasmus

func (k Keeper) CreateAbortOperationStringExtendErasmus(student_index string, student_foreign_index string, student_home_uni string, student_foreign_uni string) (abort_op_JSON string, err error) {

	var abort_op utilfunc.AbortOperationPacket

	abort_op.HomeIndex = student_index
	abort_op.ForeignIndex = student_foreign_index
	abort_op.ForeignUniversity = student_foreign_uni
	abort_op.HomeUniversity = student_home_uni
	abort_op.PacketID = "-3"

	resultByteJSON, err := json.Marshal(abort_op)
	if err != nil {
		return abort_op_JSON, err
	}

	abort_op_JSON = string(resultByteJSON)

	return abort_op_JSON, err
}

// Function that constructs the ack contents and returns the ack related to the abort
// operation of the end erasmus

func (k Keeper) HandleAbortAckExtendErasmus(ctx sdk.Context, student_index string, student_foreign_index string, student_home_uni string, student_foreign_uni string) (packetAck types.ExtendErasmusPeriodPacketAck, err error) {

	utilfunc.PrintLogs("HandleAbortAckExtendErasmus", ctx)

	new_data_info, err := k.CreateAbortOperationStringExtendErasmus(student_index, student_foreign_index, student_home_uni, student_foreign_uni)
	if err != nil {
		utilfunc.PrintLogs("HandleAbortAckExtendErasmus error "+err.Error(), ctx)
		return packetAck, err
	} else {

		packetAck.ErasmusRestrictedInfo = new_data_info
		utilfunc.PrintLogs("HandleAbortAckExtendErasmus "+new_data_info, ctx)
		utilfunc.PrintLogs("HandleAbortAckExtendErasmus "+packetAck.ErasmusRestrictedInfo, ctx)
		return packetAck, err

	}
}
