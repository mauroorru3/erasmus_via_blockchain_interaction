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

// TransmitErasmusIndexPacket transmits the packet over IBC with the specified source port and source channel
func (k Keeper) TransmitErasmusIndexPacket(
	ctx sdk.Context,
	packetData types.ErasmusIndexPacketData,
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

// OnRecvErasmusIndexPacket processes packet reception
func (k Keeper) OnRecvErasmusIndexPacket(ctx sdk.Context, packet channeltypes.Packet, data types.ErasmusIndexPacketData) (packetAck types.ErasmusIndexPacketAck, err error) {
	// validate packet data upon receiving
	err = data.ValidateBasic()

	if err != nil {
		return packetAck, err
	} else {

		// TODO: packet reception logic

		searchedStudent, found := k.GetStoredStudent(ctx, data.Index)
		if !found {
			utilfunc.PrintLogs("OnRecvErasmusIndexPacket " + types.ErrStudentNotPresent.Error())
			return packetAck, types.ErrStudentNotPresent
		} else {

			utilfunc.PrintLogs("OnRecvErasmusIndexPacket success")
			utilfunc.PrintData("OnRecvErasmusIndexPacket " + data.String())

			utilfunc.SetForeignIndex(&searchedStudent, data.ForeignIndex)

			k.SetStoredStudent(ctx, searchedStudent)
			//------------------

			stu, found := k.GetStoredStudent(ctx, data.Index)
			if !found {
				return packetAck, err
			} else {

				data, err := utilfunc.CreateNameSurnameJSONPacketFromStudentData(stu)
				if err != nil {
					return packetAck, err
				}

				var packet types.ErasmusRestictedDataPacketData
				packet.ErasmusRestrictedInfo = data

				err = k.TransmitErasmusRestictedDataPacket(
					ctx,
					packet,
					"universitychainde",
					"channel-0",
					clienttypes.ZeroHeight(),
					timeoutTimestamp,
				)

				if err != nil {
					utilfunc.PrintLogs("SendErasmusStudent " + err.Error())
					return packetAck, err
				} else {

					utilfunc.PrintLogs("SendErasmusStudent CreateNameSurnameJSONPacketFromStudentData sent")

					data, err := utilfunc.CreateStudentKeyPart1JSONPacketFromStudentData(stu)
					if err != nil {
						return packetAck, err
					}

					var packet types.ErasmusRestictedDataPacketData
					packet.ErasmusRestrictedInfo = data

					err = k.TransmitErasmusRestictedDataPacket(
						ctx,
						packet,
						"universitychainde",
						"channel-0",
						clienttypes.ZeroHeight(),
						timeoutTimestamp,
					)

					if err != nil {
						utilfunc.PrintLogs("SendErasmusStudent " + err.Error())
						return packetAck, err
					} else {

						utilfunc.PrintLogs("SendErasmusStudent CreateStudentKeyPart1JSONPacketFromStudentData sent")

						data, err := utilfunc.CreateStudentKeyPart2JSONPacketFromStudentData(stu)
						if err != nil {
							return packetAck, err
						}

						var packet types.ErasmusRestictedDataPacketData
						packet.ErasmusRestrictedInfo = data

						err = k.TransmitErasmusRestictedDataPacket(
							ctx,
							packet,
							"universitychainde",
							"channel-0",
							clienttypes.ZeroHeight(),
							timeoutTimestamp,
						)

						if err != nil {
							utilfunc.PrintLogs("SendErasmusStudent " + err.Error())
							return packetAck, err
						} else {
							utilfunc.PrintLogs("SendErasmusStudent CreateStudentKeyPart2JSONPacketFromStudentData sent")

							data, err := utilfunc.CreateStartDateJSONPacketFromStudentData(stu)
							if err != nil {
								return packetAck, err
							}

							var packet types.ErasmusRestictedDataPacketData
							packet.ErasmusRestrictedInfo = data

							err = k.TransmitErasmusRestictedDataPacket(
								ctx,
								packet,
								"universitychainde",
								"channel-0",
								clienttypes.ZeroHeight(),
								timeoutTimestamp,
							)

							if err != nil {
								utilfunc.PrintLogs("SendErasmusStudent " + err.Error())
								return packetAck, err
							} else {
								utilfunc.PrintLogs("SendErasmusStudent CreateStartDateJSONPacketFromStudentData sent")
								data, err := utilfunc.CreateEndDateJSONPacketFromStudentData(stu)
								if err != nil {
									return packetAck, err
								}

								var packet types.ErasmusRestictedDataPacketData
								packet.ErasmusRestrictedInfo = data

								err = k.TransmitErasmusRestictedDataPacket(
									ctx,
									packet,
									"universitychainde",
									"channel-0",
									clienttypes.ZeroHeight(),
									timeoutTimestamp,
								)

								if err != nil {
									utilfunc.PrintLogs("SendErasmusStudent " + err.Error())
									return packetAck, err
								} else {
									utilfunc.PrintLogs("SendErasmusStudent CreateEndDateJSONPacketFromStudentData sent")
									data, err := utilfunc.CreateDurationJSONPacketFromStudentData(stu)
									if err != nil {
										return packetAck, err
									}

									var packet types.ErasmusRestictedDataPacketData
									packet.ErasmusRestrictedInfo = data

									err = k.TransmitErasmusRestictedDataPacket(
										ctx,
										packet,
										"universitychainde",
										"channel-0",
										clienttypes.ZeroHeight(),
										timeoutTimestamp,
									)

									if err != nil {
										utilfunc.PrintLogs("SendErasmusStudent " + err.Error())
										return packetAck, err
									} else {
										utilfunc.PrintLogs("SendErasmusStudent CreateDurationJSONPacketFromStudentData sent")

										data, err := utilfunc.CreateCourseDetailsJSONPacketFromStudentData(stu)
										if err != nil {
											utilfunc.PrintLogs("SendErasmusStudent " + err.Error())
											return packetAck, err
										}

										var packet types.ErasmusRestictedDataPacketData
										packet.ErasmusRestrictedInfo = data

										err = k.TransmitErasmusRestictedDataPacket(
											ctx,
											packet,
											"universitychainde",
											"channel-0",
											clienttypes.ZeroHeight(),
											timeoutTimestamp,
										)

										if err != nil {
											utilfunc.PrintLogs("SendErasmusStudent " + err.Error())
											return packetAck, err
										} else {
											utilfunc.PrintLogs("SendErasmusStudent CreateCourseDetailsJSONPacketFromStudentData sent")
											data, err := utilfunc.CreateDepartmentJSONPacketFromStudentData(stu)
											if err != nil {
												return packetAck, err
											}

											var packet types.ErasmusRestictedDataPacketData
											packet.ErasmusRestrictedInfo = data

											err = k.TransmitErasmusRestictedDataPacket(
												ctx,
												packet,
												"universitychainde",
												"channel-0",
												clienttypes.ZeroHeight(),
												timeoutTimestamp,
											)

											if err != nil {
												utilfunc.PrintLogs("SendErasmusStudent " + err.Error())
												return packetAck, err
											} else {
												utilfunc.PrintLogs("SendErasmusStudent CreateDepartmentJSONPacketFromStudentData sent")
												data, err := utilfunc.CreateErasmusTypeJSONPacketFromStudentData(stu)
												if err != nil {
													utilfunc.PrintLogs("SendErasmusStudent " + err.Error())
													return packetAck, err
												}

												var packet types.ErasmusRestictedDataPacketData
												packet.ErasmusRestrictedInfo = data

												err = k.TransmitErasmusRestictedDataPacket(
													ctx,
													packet,
													"universitychainde",
													"channel-0",
													clienttypes.ZeroHeight(),
													timeoutTimestamp,
												)

												if err != nil {
													utilfunc.PrintLogs("SendErasmusStudent " + err.Error())
													return packetAck, err
												} else {
													utilfunc.PrintLogs("SendErasmusStudent CreateErasmusTypeJSONPacketFromStudentData sent")

													data, err := utilfunc.CreateExamsJSONPacketFromStudentData(stu)
													if err != nil {
														utilfunc.PrintLogs("SendErasmusStudent " + err.Error())
														return packetAck, err
													}

													var packet types.ErasmusRestictedDataPacketData
													packet.ErasmusRestrictedInfo = data

													err = k.TransmitErasmusRestictedDataPacket(
														ctx,
														packet,
														"universitychainde",
														"channel-0",
														clienttypes.ZeroHeight(),
														timeoutTimestamp,
													)

													if err != nil {
														utilfunc.PrintLogs("SendErasmusStudent " + err.Error())
														return packetAck, err
													} else {
														utilfunc.PrintLogs("SendErasmusStudent CreateExamsJSONPacketFromStudentData sent")
														return packetAck, err
													}
												}
											}
										}
									}
								}
							}

						}
					}
				}
			}

		}
	}

}

// OnAcknowledgementErasmusIndexPacket responds to the the success or failure of a packet
// acknowledgement written on the receiving chain.
func (k Keeper) OnAcknowledgementErasmusIndexPacket(ctx sdk.Context, packet channeltypes.Packet, data types.ErasmusIndexPacketData, ack channeltypes.Acknowledgement) error {
	switch dispatchedAck := ack.Response.(type) {
	case *channeltypes.Acknowledgement_Error:

		// TODO: failed acknowledgement logic
		_ = dispatchedAck.Error

		utilfunc.PrintLogs("OnAcknowledgementErasmusIndexPacket error " + dispatchedAck.Error)

		return nil
	case *channeltypes.Acknowledgement_Result:
		// Decode the packet acknowledgment
		var packetAck types.ErasmusIndexPacketAck

		if err := types.ModuleCdc.UnmarshalJSON(dispatchedAck.Result, &packetAck); err != nil {
			// The counter-party module doesn't implement the correct acknowledgment format
			utilfunc.PrintLogs("OnAcknowledgementErasmusIndexPacket cannot unmarshal acknowledgment")
			return errors.New("cannot unmarshal acknowledgment")
		}

		// TODO: successful acknowledgement logic

		utilfunc.PrintLogs("OnAcknowledgementErasmusIndexPacket success")

		return nil
	default:
		// The counter-party module doesn't implement the correct acknowledgment format
		return errors.New("invalid acknowledgment format")
	}
}

// OnTimeoutErasmusIndexPacket responds to the case where a packet has not been transmitted because of a timeout
func (k Keeper) OnTimeoutErasmusIndexPacket(ctx sdk.Context, packet channeltypes.Packet, data types.ErasmusIndexPacketData) error {

	// TODO: packet timeout logic

	utilfunc.PrintLogs("OnTimeoutErasmusIndexPacket")

	return nil
}
