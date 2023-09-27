package keeper

import (
	"encoding/json"
	"errors"
	"strconv"

	"university_chain_it/x/universitychainit/types"
	"university_chain_it/x/universitychainit/utilfunc"

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
		return packetAck, err
	}
	utilfunc.GetTransactionStats("OnRecvErasmusRestictedDataPacket", "", ctx, sizeInt, binArray)

	utilfunc.PrintLogs("OnRecvErasmusStudentPacket")
	utilfunc.PrintData(data.String())

	// TODO: packet reception logic

	var result map[string]interface{}
	err = json.Unmarshal([]byte(data.ErasmusRestrictedInfo), &result)
	if err != nil {
		return packetAck, err
	}

	packetID, found := result["p_id"].(string)
	if !found {
		return packetAck, sdkerrors.Error{}
	}

	switch packetID {
	case "1":

		utilfunc.PrintLogs("OnRecvErasmusStudentPacket case 1")
		var homeIndexPacket utilfunc.StudentInfoRestrictedHomeIndexPacket
		err = json.Unmarshal([]byte(data.ErasmusRestrictedInfo), &homeIndexPacket)
		if err != nil {
			return packetAck, err
		}

		allStudents := k.GetAllStoredStudent(ctx)
		found = false
		i := 0
		for i < len(allStudents) && !found {
			foreignIndex, err := utilfunc.GetForeignIndex(allStudents[i])
			if err != nil {
				return packetAck, err
			} else {
				if foreignIndex == homeIndexPacket.HomeIndex {
					found = true
				} else {
					i++
				}
			}
		}

		if found {
			return packetAck, types.ErrStudentAlreadyPresent
		} else {

			erasmusData, err := utilfunc.IntializeErasmusStructForeign(homeIndexPacket.HomeUniversity, homeIndexPacket.HomeIndex)
			if err != nil {
				return packetAck, err
			} else {
				uniInfo, found := k.GetUniversityInfo(ctx, homeIndexPacket.ForeignUniversity)
				if !found {
					return packetAck, types.ErrWrongNameUniversity
				} else {
					student := types.StoredStudent{
						Index: uniInfo.UniversityName + "_" + strconv.FormatUint(uint64(uniInfo.NextStudentId), 10),
						StudentData: &types.StudentInfo{
							UniversityName: homeIndexPacket.ForeignUniversity,
						},
						TranscriptData: &types.TranscriptOfRecords{},
						PersonalData:   &types.PersonalInfo{},
						ResidenceData:  &types.ResidenceInfo{},
						ContactData:    &types.ContactInfo{},
						TaxesData:      &types.TaxesInfo{},
						ErasmusData: &types.ErasmusInfo{
							ErasmusStudent: "Incoming",
							Career:         erasmusData,
						},
					}
					uniInfo.NextStudentId++
					resAck, err := utilfunc.CreateAnswerJSONPacketFromStudentData(student)
					if err != nil {
						return packetAck, err
					} else {

						packetAck.ErasmusRestrictedInfo = resAck
						utilfunc.PrintData(packetAck.ErasmusRestrictedInfo)

						err = utilfunc.SetForeignIndex(&student, homeIndexPacket.HomeIndex)
						if err != nil {
							return packetAck, err
						} else {
							homeUni, _ := k.GetForeignUniversities(ctx, homeIndexPacket.HomeUniversity)
							utilfunc.SetHomeUniversityInfo(&student, homeUni.ChainName, homeUni.ForeignUniversitiesCountry)
							k.SetStoredStudent(ctx, student)
							k.SetUniversityInfo(ctx, uniInfo)
							utilfunc.PrintLogs("OnRecvErasmusStudentPacket ack sent")

							err = utilfunc.GetConsumedGas("OnRecvErasmusRestictedDataPacket IT - case 1", homeIndexPacket.HomeIndex, ctx)
							if err != nil {
								return packetAck, err
							} else {
								if err != nil {
									return packetAck, err
								} else {

									packetAckBytes, err := types.ModuleCdc.MarshalJSON(&packetAck)
									if err != nil {
										return packetAck, err
									}
									sizeInt := len(packetAckBytes)
									utilfunc.GetTransactionStats("OnRecvErasmusRestictedDataPacket sending ack", " - case 1", ctx, sizeInt, binArray)

									return packetAck, nil
								}

							}
						}
					}
				}

			}
		}
	case "2":
		utilfunc.PrintLogs("OnRecvErasmusStudentPacket case 2")
		utilfunc.PrintData("OnRecvErasmusStudentPacket case 2 " + data.String())
		var nameSurnamePacket utilfunc.StudentInfoRestrictedNameSurnamePacket
		err = json.Unmarshal([]byte(data.ErasmusRestrictedInfo), &nameSurnamePacket)
		if err != nil {
			return packetAck, err
		}

		student, found := k.GetStoredStudent(ctx, nameSurnamePacket.ForeignIndex)
		if !found {
			return packetAck, types.ErrStudentNotPresent
		} else {
			utilfunc.PrintLogs("OnRecvErasmusStudentPacket case 2 qui")
			student.StudentData.Name = nameSurnamePacket.Name
			student.StudentData.Surname = nameSurnamePacket.Surname
			k.SetStoredStudent(ctx, student)
			packetAck.ErasmusRestrictedInfo = ""

			stringIndex, err := utilfunc.GetForeignIndex(student)
			if err != nil {
				return packetAck, err
			} else {
				err = utilfunc.GetConsumedGas("OnRecvErasmusRestictedDataPacket IT - case 2", stringIndex, ctx)
				if err != nil {
					return packetAck, err
				} else {

					packetAckBytes, err := types.ModuleCdc.MarshalJSON(&packetAck)
					if err != nil {
						return packetAck, err
					}
					sizeInt := len(packetAckBytes)
					utilfunc.GetTransactionStats("OnRecvErasmusRestictedDataPacket sending ack", " - case 2", ctx, sizeInt, binArray)
					return packetAck, nil
				}

			}
		}

	case "3":
		utilfunc.PrintLogs("OnRecvErasmusStudentPacket case 3")
		var studentKeyPacket utilfunc.StudentInfoRestrictedStudentKeyPacket
		err = json.Unmarshal([]byte(data.ErasmusRestrictedInfo), &studentKeyPacket)
		if err != nil {
			return packetAck, err
		}

		student, found := k.GetStoredStudent(ctx, studentKeyPacket.ForeignIndex)
		if !found {
			return packetAck, types.ErrStudentNotPresent
		} else {
			student.StudentData.StudentKey = studentKeyPacket.StudentKey
			k.SetStoredStudent(ctx, student)
			packetAck.ErasmusRestrictedInfo = ""

			stringIndex, err := utilfunc.GetForeignIndex(student)
			if err != nil {
				return packetAck, err
			} else {
				err = utilfunc.GetConsumedGas("OnRecvErasmusRestictedDataPacket IT - case 3", stringIndex, ctx)
				if err != nil {
					return packetAck, err
				} else {

					packetAckBytes, err := types.ModuleCdc.MarshalJSON(&packetAck)
					if err != nil {
						return packetAck, err
					}
					sizeInt := len(packetAckBytes)
					utilfunc.GetTransactionStats("OnRecvErasmusRestictedDataPacket sending ack", " - case 3", ctx, sizeInt, binArray)

					return packetAck, nil
				}

			}
		}
	case "4":
		utilfunc.PrintLogs("OnRecvErasmusStudentPacket case 4")
		var studentKeyPacket utilfunc.StudentInfoRestrictedStudentKeyPacket
		err = json.Unmarshal([]byte(data.ErasmusRestrictedInfo), &studentKeyPacket)
		if err != nil {
			return packetAck, err
		}

		student, found := k.GetStoredStudent(ctx, studentKeyPacket.ForeignIndex)
		if !found {
			return packetAck, types.ErrStudentNotPresent
		} else {
			student.StudentData.StudentKey = student.StudentData.StudentKey + studentKeyPacket.StudentKey
			k.SetStoredStudent(ctx, student)
			packetAck.ErasmusRestrictedInfo = ""

			stringIndex, err := utilfunc.GetForeignIndex(student)
			if err != nil {
				return packetAck, err
			} else {
				err = utilfunc.GetConsumedGas("OnRecvErasmusRestictedDataPacket IT - case 4", stringIndex, ctx)
				if err != nil {
					return packetAck, err
				} else {

					packetAckBytes, err := types.ModuleCdc.MarshalJSON(&packetAck)
					if err != nil {
						return packetAck, err
					}
					sizeInt := len(packetAckBytes)
					utilfunc.GetTransactionStats("OnRecvErasmusRestictedDataPacket sending ack", " - case 4", ctx, sizeInt, binArray)

					return packetAck, nil
				}

			}
		}
	case "5":
		utilfunc.PrintLogs("OnRecvErasmusStudentPacket case 5")
		var startDatePacket utilfunc.StudentInfoRestrictedStartDatePacket
		err = json.Unmarshal([]byte(data.ErasmusRestrictedInfo), &startDatePacket)
		if err != nil {
			return packetAck, err
		}

		student, found := k.GetStoredStudent(ctx, startDatePacket.ForeignIndex)
		if !found {
			return packetAck, types.ErrStudentNotPresent
		} else {

			var erasmusCareer []utilfunc.ErasmusCareerStruct
			err = json.Unmarshal([]byte(student.ErasmusData.Career), &erasmusCareer)
			if err != nil {
				return packetAck, err
			}
			lenCareer := len(erasmusCareer)
			erasmusCareer[lenCareer-1].Start_date = startDatePacket.StartDate

			resultByteJSON, err := json.Marshal(erasmusCareer)
			if err != nil {
				return packetAck, err
			}
			student.ErasmusData.Career = string(resultByteJSON)

			utilfunc.PrintLogs("OnRecvErasmusStudentPacket case 5 packet " + startDatePacket.StartDate)
			utilfunc.PrintLogs("OnRecvErasmusStudentPacket case 5 packet " + startDatePacket.PacketID)
			utilfunc.PrintLogs("OnRecvErasmusStudentPacket case 5 student " + student.ErasmusData.Career)
			utilfunc.PrintLogs("OnRecvErasmusStudentPacket case 5 data " + data.String())

			k.SetStoredStudent(ctx, student)
			packetAck.ErasmusRestrictedInfo = ""

			stringIndex, err := utilfunc.GetForeignIndex(student)
			if err != nil {
				return packetAck, err
			} else {
				err = utilfunc.GetConsumedGas("OnRecvErasmusRestictedDataPacket IT - case 5", stringIndex, ctx)
				if err != nil {
					return packetAck, err
				} else {

					packetAckBytes, err := types.ModuleCdc.MarshalJSON(&packetAck)
					if err != nil {
						return packetAck, err
					}
					sizeInt := len(packetAckBytes)
					utilfunc.GetTransactionStats("OnRecvErasmusRestictedDataPacket sending ack", " - case 5", ctx, sizeInt, binArray)

					return packetAck, nil
				}

			}
		}
	case "6":
		utilfunc.PrintLogs("OnRecvErasmusStudentPacket case 6")
		var endDatePacket utilfunc.StudentInfoRestrictedEndDatePacket
		err = json.Unmarshal([]byte(data.ErasmusRestrictedInfo), &endDatePacket)
		if err != nil {
			return packetAck, err
		}

		student, found := k.GetStoredStudent(ctx, endDatePacket.ForeignIndex)
		if !found {
			return packetAck, types.ErrStudentNotPresent
		} else {

			var erasmusCareer []utilfunc.ErasmusCareerStruct
			err = json.Unmarshal([]byte(student.ErasmusData.Career), &erasmusCareer)
			if err != nil {
				return packetAck, err
			}
			lenCareer := len(erasmusCareer)
			erasmusCareer[lenCareer-1].End_date = endDatePacket.EndDate

			resultByteJSON, err := json.Marshal(erasmusCareer)
			if err != nil {
				return packetAck, err
			}
			student.ErasmusData.Career = string(resultByteJSON)

			utilfunc.PrintLogs("OnRecvErasmusStudentPacket case 6 " + student.ErasmusData.Career)

			k.SetStoredStudent(ctx, student)
			packetAck.ErasmusRestrictedInfo = ""

			stringIndex, err := utilfunc.GetForeignIndex(student)
			if err != nil {
				return packetAck, err
			} else {
				err = utilfunc.GetConsumedGas("OnRecvErasmusRestictedDataPacket IT - case 6", stringIndex, ctx)
				if err != nil {
					return packetAck, err
				} else {

					packetAckBytes, err := types.ModuleCdc.MarshalJSON(&packetAck)
					if err != nil {
						return packetAck, err
					}
					sizeInt := len(packetAckBytes)
					utilfunc.GetTransactionStats("OnRecvErasmusRestictedDataPacket sending ack", " - case 6", ctx, sizeInt, binArray)

					return packetAck, nil

				}
			}
		}
	case "7":
		utilfunc.PrintLogs("OnRecvErasmusStudentPacket case 7")
		var durationPacket utilfunc.StudentInfoRestrictedDurationPacket
		err = json.Unmarshal([]byte(data.ErasmusRestrictedInfo), &durationPacket)
		if err != nil {
			return packetAck, err
		}

		student, found := k.GetStoredStudent(ctx, durationPacket.ForeignIndex)
		if !found {
			return packetAck, types.ErrStudentNotPresent
		} else {

			var erasmusCareer []utilfunc.ErasmusCareerStruct
			err = json.Unmarshal([]byte(student.ErasmusData.Career), &erasmusCareer)
			if err != nil {
				return packetAck, err
			}
			lenCareer := len(erasmusCareer)
			elem, _ := strconv.ParseUint(durationPacket.DurationInMonths, 10, 32)
			erasmusCareer[lenCareer-1].Duration_in_months = uint8(elem)

			resultByteJSON, err := json.Marshal(erasmusCareer)
			if err != nil {
				return packetAck, err
			}
			student.ErasmusData.Career = string(resultByteJSON)
			k.SetStoredStudent(ctx, student)
			packetAck.ErasmusRestrictedInfo = ""

			stringIndex, err := utilfunc.GetForeignIndex(student)
			if err != nil {
				return packetAck, err
			} else {
				err = utilfunc.GetConsumedGas("OnRecvErasmusRestictedDataPacket IT - case 7", stringIndex, ctx)
				if err != nil {
					return packetAck, err
				} else {

					packetAckBytes, err := types.ModuleCdc.MarshalJSON(&packetAck)
					if err != nil {
						return packetAck, err
					}
					sizeInt := len(packetAckBytes)
					utilfunc.GetTransactionStats("OnRecvErasmusRestictedDataPacket sending ack", " - case 7", ctx, sizeInt, binArray)
					return packetAck, nil
				}

			}
		}
	case "8":
		utilfunc.PrintLogs("OnRecvErasmusStudentPacket case 8")
		var courseDetailsPacket utilfunc.StudentInfoRestrictedCourseDetailsPacket
		err = json.Unmarshal([]byte(data.ErasmusRestrictedInfo), &courseDetailsPacket)
		if err != nil {
			return packetAck, err
		}

		student, found := k.GetStoredStudent(ctx, courseDetailsPacket.ForeignIndex)
		if !found {
			return packetAck, types.ErrStudentNotPresent
		} else {

			student.StudentData.CourseOfStudy = courseDetailsPacket.CourseOfStudy
			student.StudentData.CourseType = courseDetailsPacket.CourseType
			k.SetStoredStudent(ctx, student)
			packetAck.ErasmusRestrictedInfo = ""

			stringIndex, err := utilfunc.GetForeignIndex(student)
			if err != nil {
				return packetAck, err
			} else {
				err = utilfunc.GetConsumedGas("OnRecvErasmusRestictedDataPacket IT - case 8", stringIndex, ctx)
				if err != nil {
					return packetAck, err
				} else {

					packetAckBytes, err := types.ModuleCdc.MarshalJSON(&packetAck)
					if err != nil {
						return packetAck, err
					}
					sizeInt := len(packetAckBytes)
					utilfunc.GetTransactionStats("OnRecvErasmusRestictedDataPacket sending ack", " - case 8", ctx, sizeInt, binArray)
					return packetAck, nil
				}

			}
		}
	case "9":
		utilfunc.PrintLogs("OnRecvErasmusStudentPacket case 9")
		var departmentPacket utilfunc.StudentInfoRestrictedDepartmentPacket
		err = json.Unmarshal([]byte(data.ErasmusRestrictedInfo), &departmentPacket)
		if err != nil {
			return packetAck, err
		}

		student, found := k.GetStoredStudent(ctx, departmentPacket.ForeignIndex)
		if !found {
			return packetAck, types.ErrStudentNotPresent
		} else {
			student.StudentData.DepartmentName = departmentPacket.DepartmentName
			k.SetStoredStudent(ctx, student)
			packetAck.ErasmusRestrictedInfo = ""

			stringIndex, err := utilfunc.GetForeignIndex(student)
			if err != nil {
				return packetAck, err
			} else {
				err = utilfunc.GetConsumedGas("OnRecvErasmusRestictedDataPacket IT - case 9", stringIndex, ctx)
				if err != nil {
					return packetAck, err
				} else {

					packetAckBytes, err := types.ModuleCdc.MarshalJSON(&packetAck)
					if err != nil {
						return packetAck, err
					}
					sizeInt := len(packetAckBytes)
					utilfunc.GetTransactionStats("OnRecvErasmusRestictedDataPacket sending ack", " - case 9", ctx, sizeInt, binArray)
					return packetAck, nil
				}

			}
		}
	case "10":
		utilfunc.PrintLogs("OnRecvErasmusStudentPacket case 10")
		var erasmusTypePacket utilfunc.StudentInfoRestrictedErasmusTypePacket
		err = json.Unmarshal([]byte(data.ErasmusRestrictedInfo), &erasmusTypePacket)
		if err != nil {
			return packetAck, err
		}

		student, found := k.GetStoredStudent(ctx, erasmusTypePacket.ForeignIndex)
		if !found {
			return packetAck, types.ErrStudentNotPresent
		} else {
			utilfunc.SetErasmusType(&student, erasmusTypePacket.ErasmusType)
			k.SetStoredStudent(ctx, student)
			packetAck.ErasmusRestrictedInfo = ""

			stringIndex, err := utilfunc.GetForeignIndex(student)
			if err != nil {
				return packetAck, err
			} else {
				err = utilfunc.GetConsumedGas("OnRecvErasmusRestictedDataPacket IT - case 10", stringIndex, ctx)
				if err != nil {
					return packetAck, err
				} else {

					packetAckBytes, err := types.ModuleCdc.MarshalJSON(&packetAck)
					if err != nil {
						return packetAck, err
					}
					sizeInt := len(packetAckBytes)
					utilfunc.GetTransactionStats("OnRecvErasmusRestictedDataPacket sending ack", " - case 10", ctx, sizeInt, binArray)
					return packetAck, nil
				}

			}
		}
	case "11":
		utilfunc.PrintLogs("OnRecvErasmusStudentPacket case 11")
		var examsPacket utilfunc.StudentInfoRestrictedExamsPacket
		err = json.Unmarshal([]byte(data.ErasmusRestrictedInfo), &examsPacket)
		if err != nil {
			return packetAck, err
		}

		student, found := k.GetStoredStudent(ctx, examsPacket.ForeignIndex)
		if !found {
			return packetAck, types.ErrStudentNotPresent
		} else {

			mapExams := make(map[string]utilfunc.ExamStruct)

			examsJSON, _, err := utilfunc.GetJSONFromCourseExams(examsPacket.ForeignUniversity, student.StudentData.DepartmentName, student.StudentData.CourseType, student.StudentData.CourseOfStudy)
			if err != nil {
				return packetAck, err
			}

			err = json.Unmarshal([]byte(examsJSON), &mapExams)
			if err != nil {
				return packetAck, err
			}

			var erasmusCareer []utilfunc.ErasmusCareerStruct
			err = json.Unmarshal([]byte(student.ErasmusData.Career), &erasmusCareer)
			if err != nil {
				return packetAck, err
			}
			lenCareer := len(erasmusCareer)

			var resultByteJSON []byte
			mapExamsErasmus := make(map[string]utilfunc.ExamStruct)

			var val utilfunc.ExamStruct
			for i := 0; i < len(examsPacket.ExamsData); i++ {
				val = mapExams[examsPacket.ExamsData[i]]
				mapExamsErasmus[examsPacket.ExamsData[i]] = val
				erasmusCareer[lenCareer-1].Total_credits += mapExams[examsPacket.ExamsData[i]].Credits
				student.ErasmusData.TotalCredits += uint32(mapExams[examsPacket.ExamsData[i]].Credits)
			}

			erasmusCareer[lenCareer-1].Total_exams = uint8(len(examsPacket.ExamsData))
			erasmusCareer[lenCareer-1].Status = "In progress"
			student.ErasmusData.TotalExams = uint32(len(examsPacket.ExamsData))

			resultByteJSON, err = json.Marshal(mapExamsErasmus)
			if err != nil {
				return packetAck, err
			}

			erasmusCareer[lenCareer-1].Exams_data = string(resultByteJSON)

			resultByteJSON, err = json.Marshal(erasmusCareer)
			if err != nil {
				return packetAck, err
			}
			student.ErasmusData.Career = string(resultByteJSON)

			k.SetStoredStudent(ctx, student)
			packetAck.ErasmusRestrictedInfo = ""

			stringIndex, err := utilfunc.GetForeignIndex(student)
			if err != nil {
				return packetAck, err
			} else {
				err = utilfunc.GetConsumedGas("OnRecvErasmusRestictedDataPacket IT - case 11", stringIndex, ctx)
				if err != nil {
					return packetAck, err
				} else {
					packetAckBytes, err := types.ModuleCdc.MarshalJSON(&packetAck)
					if err != nil {
						return packetAck, err
					}
					sizeInt := len(packetAckBytes)
					utilfunc.GetTransactionStats("OnRecvErasmusRestictedDataPacket sending ack", " - case 11", ctx, sizeInt, binArray)
					return packetAck, nil
				}

			}
		}
	}

	return packetAck, err

}

// OnAcknowledgementErasmusRestictedDataPacket responds to the the success or failure of a packet
// acknowledgement written on the receiving chain.
func (k Keeper) OnAcknowledgementErasmusRestictedDataPacket(ctx sdk.Context, packet channeltypes.Packet, data types.ErasmusRestictedDataPacketData, ack channeltypes.Acknowledgement) error {

	switch dispatchedAck := ack.Response.(type) {
	case *channeltypes.Acknowledgement_Error:

		// TODO: failed acknowledgement logic
		_ = dispatchedAck.Error

		utilfunc.PrintLogs("OnAcknowledgementErasmusRestictedDataPacket error " + dispatchedAck.Error)

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
			utilfunc.PrintLogs("OnAcknowledgementErasmusRestictedDataPacket cannot unmarshal acknowledgment")
			return errors.New("cannot unmarshal acknowledgment")
		}

		// TODO: successful acknowledgement logic
		utilfunc.PrintLogs("OnAcknowledgementErasmusRestictedDataPacket success")

		packetHash := utilfunc.Hash(binArray)
		err = utilfunc.GetConsumedGas("OnAcknowledgementErasmusRestictedDataPacket IT", strconv.FormatInt(int64(packetHash), 10), ctx)
		if err != nil {
			return err
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
	utilfunc.PrintLogs("OnTimeoutErasmusRestictedDataPacket")

	return nil
}
