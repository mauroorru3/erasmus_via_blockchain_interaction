package keeper

import (
	"encoding/json"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"university_chain_it/x/universitychainit/types"
	"university_chain_it/x/universitychainit/utilfunc"

	clienttypes "github.com/cosmos/ibc-go/v3/modules/core/02-client/types"
)

// ***Timeout and error management of the Start Erasmus operation***

//Function that resets the Erasmus career

func (k Keeper) ClearErasmusCareer(ctx sdk.Context, studentIndex string) (err error) {

	utilfunc.PrintLogs("ClearErasmusCareer", ctx)

	student, found := k.GetStoredStudent(ctx, studentIndex)
	if !found {
		return types.ErrStudentNotPresent
	} else {

		if student.ErasmusData.ErasmusStudent == "Outgoing" {

			uniInfo, found := k.GetUniversityInfo(ctx, student.StudentData.UniversityName)
			if !found {
				return types.ErrWrongNameUniversity
			} else {

				k.RemoveFromFifo(ctx, &student, &uniInfo)
				k.SetUniversityInfo(ctx, uniInfo)

				var erasmusCareer []utilfunc.ErasmusCareerStruct

				err = json.Unmarshal([]byte(student.ErasmusData.Career), &erasmusCareer)
				if err != nil {
					return err
				}

				lenCareer := len(erasmusCareer)

				student.ErasmusData.ErasmusStudent = "No"
				student.ErasmusData.NumberTimes--
				student.ErasmusData.NumberMonths = student.ErasmusData.NumberMonths - uint32(erasmusCareer[lenCareer-1].Duration_in_months)
				student.ErasmusData.TotalExams = student.ErasmusData.TotalExams - uint32(erasmusCareer[lenCareer-1].Total_exams)
				//student.ErasmusData.ExamsPassed -> it doesn't change because the first packet was lost and therefore Erasmus didn't start
				student.ErasmusData.TotalCredits = student.ErasmusData.TotalCredits - uint32(erasmusCareer[lenCareer-1].Total_credits)
				//student.ErasmusData.AchievedCredits -> it doesn't change because the first packet was lost and therefore Erasmus didn't start
				//student.ErasmusData.PreviousStudentFifo = ""-> it is already done inside the RemoveFromFifo function
				//student.ErasmusData.NextStudentFifo = "" -> it is already done inside the RemoveFromFifo function

				var contributionInfo = erasmusCareer[lenCareer-1].Contribution

				erasmusInfo := utilfunc.ErasmusCareerStruct{
					Duration_in_months:            0,
					Start_date:                    "",
					End_date:                      "",
					Erasmus_type:                  "",
					Total_credits:                 0,
					Achieved_credits:              0,
					Total_exams:                   0,
					Exams_passed:                  0,
					Foreign_university_name:       "",
					Foreign_university_country:    "",
					Foreign_university_student_id: "",
					Foreign_chain_name:            "",
					Status:                        "",
					Contribution:                  contributionInfo,
					Exams_data:                    "",
				}

				erasmusCareer[lenCareer-1] = erasmusInfo

				resultByteJSON, err := json.Marshal(erasmusCareer)
				if err != nil {
					return err
				}

				erasmusJSON := string(resultByteJSON)

				student.ErasmusData.Career = erasmusJSON
				k.SetStoredStudent(ctx, student)
			}
		}

		err = k.ClearOperationQueue(ctx, &student)
		if err != nil {
			return err
		}

	}
	return nil
}

// Function that creates the content of the packet that sends the abort operation request
// for that considering the outgoing student and the Start Erasmus function

func (k Keeper) CreateAbortOperationString(student types.StoredStudent) (abort_op_JSON string, err error) {

	var abort_op utilfunc.AbortOperationPacket

	abort_op.HomeIndex = student.Index
	abort_op.ForeignIndex, _ = utilfunc.GetForeignIndex(student)
	abort_op.ForeignUniversity, _ = utilfunc.GetForeignUniversityName(student)
	abort_op.HomeUniversity = student.StudentData.UniversityName
	abort_op.PacketID = "-1" // Value that identifies the packet related to the abort operation of the Start Erasmus operation

	resultByteJSON, err := json.Marshal(abort_op)
	if err != nil {
		return abort_op_JSON, err
	}

	abort_op_JSON = string(resultByteJSON)

	return abort_op_JSON, err
}

// Function that sends the abort operation packet

func (k Keeper) SendAbortPacket(ctx sdk.Context, data string) (err error) {

	utilfunc.PrintLogs("SendAbortPacket", ctx)

	var packet types.ErasmusRestictedDataPacketData
	packet.ErasmusRestrictedInfo = data

	utilfunc.PrintLogs("SendAbortPacket data "+data, ctx)

	err = k.TransmitErasmusRestictedDataPacket(
		ctx,
		packet,
		"universitychainit",
		"channel-0",
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

// Function that constructs the packet contents and sends the packet related to the abort operation

func (k Keeper) HandleAbortPacket(ctx sdk.Context, studentIndex string) (err error) {

	utilfunc.PrintLogs("HandleAbortPacket", ctx)

	storedStudent, found := k.GetStoredStudent(ctx, studentIndex)
	if !found {
		return types.ErrStudentNotPresent
	} else {
	}

	new_data_info, err := k.CreateAbortOperationString(storedStudent)
	if err != nil {
		return err
	} else {

		err = k.SendAbortPacket(ctx, new_data_info)
		if err != nil {
			return err
		} else {

			utilfunc.PrintLogs("HandleAbortPacket", ctx)
			return nil
		}
	}
}

// Function that constructs the ack contents and returns the ack related to the abort operation

func (k Keeper) HandleAbortAck(ctx sdk.Context, studentIndex string) (packetAck types.ErasmusRestictedDataPacketAck, err error) {

	utilfunc.PrintLogs("HandleAbortAck", ctx)

	storedStudent, found := k.GetStoredStudent(ctx, studentIndex)
	if !found {
		return packetAck, types.ErrStudentNotPresent
	} else {
	}

	new_data_info, err := k.CreateAbortOperationString(storedStudent)
	if err != nil {
		return packetAck, err
	} else {

		packetAck.ErasmusRestrictedInfo = new_data_info
		return packetAck, err
	}
}

// Function that resets the Erasmus career,
// constructs the ack contents and returns the ack related to the abort operation

func (k Keeper) ErrorHandlingErasmus(ctx sdk.Context, studentIndex string) (packetAck types.ErasmusRestictedDataPacketAck, err error) {

	utilfunc.PrintLogs("ErrorHandlingErasmus", ctx)

	err = k.ClearErasmusCareer(ctx, studentIndex)
	if err != nil {
		return packetAck, err
	} else {

		for {
			packetAck, err = k.HandleAbortAck(ctx, studentIndex)
			if err == nil {
				break
			}
		}
		return packetAck, err
	}
}

// ***Timeout and error management of the End Erasmus and the EndErasmusBeforeDeadline operations***

// Function that takes care of moving the Erasmus period of the outgoing students forward by some
// time in such a way as to allow the deadline timer to send the packet that allows the end of
// the Erasmus again or the user to request the end_erasmus_before_deadline operation again.
// // In the case of the incoming students, this function allows you to reset the Erasmus period
// so that the end of the Erasmus never occurred

func (k Keeper) RevertEndErasmus(ctx sdk.Context, studentIndex string) (err error) {

	utilfunc.PrintLogs("RevertEndErasmus", ctx)

	student, found := k.GetStoredStudent(ctx, studentIndex)
	if !found {
		return types.ErrStudentNotPresent
	} else {

		if student.ErasmusData.ErasmusStudent == "Waiting for updated data from the destination university" {

			uniInfo, found := k.GetUniversityInfo(ctx, student.StudentData.UniversityName)
			if !found {
				return types.ErrWrongNameUniversity
			} else {

				var erasmusCareer []utilfunc.ErasmusCareerStruct

				err = json.Unmarshal([]byte(student.ErasmusData.Career), &erasmusCareer)
				if err != nil {
					return err
				}

				lenCareer := len(erasmusCareer)

				startDate := ctx.BlockTime()

				// I add a few seconds (obviously for testing) so that the Erasmus deadline is postponed and the packet is sent again
				endDate := startDate.Add(time.Duration(600 * time.Second))
				erasmusCareer[lenCareer-1].End_date = utilfunc.FormatDeadline(endDate)
				erasmusCareer[lenCareer-1].Status = "Outgoing"

				resultByteJSON, err := json.Marshal(erasmusCareer)
				if err != nil {
					return err
				}

				erasmusJSON := string(resultByteJSON)

				student.ErasmusData.Career = erasmusJSON
				student.ErasmusData.ErasmusStudent = "Outgoing"
				k.InsertInTheErasmusFIFOQueue(ctx, &student, &uniInfo)
				k.SetStoredStudent(ctx, student)
				k.SetUniversityInfo(ctx, uniInfo)
				return err

			}
		}
		err = k.ClearOperationQueue(ctx, &student)
		if err != nil {
			return err
		}
	}
	return nil
}

// Function that allows the construction of the packet content that allows the abort
// operation related to the end of the Erasmus

func (k Keeper) CreateAbortOperationStringEndErasmusV2(student types.StoredStudent) (abort_op_JSON string, err error) {

	var abort_op utilfunc.AbortOperationPacket

	abort_op.HomeIndex = student.Index
	abort_op.ForeignIndex, _ = utilfunc.GetForeignIndex(student)
	abort_op.ForeignUniversity, _ = utilfunc.GetForeignUniversityName(student)
	abort_op.HomeUniversity = student.StudentData.UniversityName
	abort_op.PacketID = "-2"

	resultByteJSON, err := json.Marshal(abort_op)
	if err != nil {
		return abort_op_JSON, err
	}

	abort_op_JSON = string(resultByteJSON)

	return abort_op_JSON, err
}

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

// Function that constructs packet contents and returns the packet related to the abort operation of the end erasmus

func (k Keeper) HandleAbortEndErasmus(ctx sdk.Context, studentIndex string) (err error) {

	utilfunc.PrintLogs("HandleAbortEndErasmus", ctx)

	storedStudent, found := k.GetStoredStudent(ctx, studentIndex)
	if !found {
		return types.ErrStudentNotPresent
	} else {
	}

	new_data_info, err := k.CreateAbortOperationStringEndErasmusV2(storedStudent)
	if err != nil {
		return err
	} else {

		err = k.SendAbortPacket(ctx, new_data_info)
		if err != nil {
			return err
		} else {

			utilfunc.PrintLogs("HandleAbortEndErasmus", ctx)
			return nil
		}
	}
}

// Function that constructs the ack contents and returns the ack related to the abort operation of the end erasmus

func (k Keeper) HandleAbortAckEndErasmus(ctx sdk.Context, studentIndex string) (packetAck types.EndErasmusPeriodRequestPacketAck, err error) {

	utilfunc.PrintLogs("HandleAbortAckEndErasmus", ctx)

	storedStudent, found := k.GetStoredStudent(ctx, studentIndex)
	if !found {
		return packetAck, types.ErrStudentNotPresent
	} else {
	}

	new_data_info, err := k.CreateAbortOperationStringEndErasmusV2(storedStudent)
	if err != nil {
		return packetAck, err
	} else {

		packetAck.ErasmusRestrictedInfo = new_data_info
		return packetAck, err
	}
}

// Function that resets the end erasmus situation,
// constructs the ack contents and returns the ack related to the abort operation

func (k Keeper) ErrorHandlingEndErasmusAck(ctx sdk.Context, studentIndex string) (packetAck types.EndErasmusPeriodRequestPacketAck, err error) {

	utilfunc.PrintLogs("ErrorHandlingEndErasmusAck", ctx)

	err = k.RevertEndErasmus(ctx, studentIndex)
	if err != nil {
		return packetAck, err
	} else {

		for {
			packetAck, err = k.HandleAbortAckEndErasmus(ctx, studentIndex)
			if err == nil {
				break
			}
		}
		return packetAck, err
	}
}

func (k Keeper) ErrorHandlingEndErasmusAckFinalPacket(ctx sdk.Context, studentIndex string, data string) (packetAck types.FinalErasmusDataPacketAck, err error) {

	utilfunc.PrintLogs("ErrorHandlingEndErasmusAckFinalPacket", ctx)

	student_index := studentIndex
	student_foreign_index := ""
	student_foreign_uni := ""
	student_home_uni := ""

	var result map[string]interface{}
	err = json.Unmarshal([]byte(data), &result)
	if err != nil {
		return packetAck, err
	}
	student_foreign_index, _ = result["f_id"].(string)
	student_foreign_uni, _ = result["f_uni"].(string)
	student_home_uni, _ = result["h_uni"].(string)

	new_data_info, err := k.CreateAbortOperationStringEndErasmus(student_index, student_foreign_index, student_home_uni, student_foreign_uni)
	if err != nil {
		return packetAck, err
	} else {

		err = k.SendAbortPacket(ctx, new_data_info)
		if err != nil {
			return packetAck, err
		} else {

			if err != nil {
				utilfunc.PrintLogs("ErrorHandlingEndErasmusAckIncomingFinalPacket "+err.Error(), ctx)
				return packetAck, err
			}
			return packetAck, err
		}
	}
}

// ***Timeout and error management of the Extend Erasmus***

// Function that reverses the Erasmus extend operation

func (k Keeper) RevertExtendErasmus(ctx sdk.Context, studentIndex string) (err error) {

	utilfunc.PrintLogs("RevertExtendErasmus", ctx)

	student, found := k.GetStoredStudent(ctx, studentIndex)
	if !found {
		utilfunc.PrintLogs("RevertExtendErasmus - error ErrStudentNotPresent", ctx)
		return types.ErrStudentNotPresent
	} else {

		if student.ErasmusData.ErasmusStudent == "Outgoing" {

			utilfunc.PrintLogs("RevertExtendErasmus - case outgoing and incoming", ctx)

			var erasmusCareer []utilfunc.ErasmusCareerStruct

			err = json.Unmarshal([]byte(student.ErasmusData.Career), &erasmusCareer)
			if err != nil {
				return err
			}

			lenCareer := len(erasmusCareer)

			// It is assumed that the Erasmus extension is 6 months, and therefore
			// it is brought back to the previous state by subtracting 6 months

			erasmusCareer[lenCareer-1].Duration_in_months = erasmusCareer[lenCareer-1].Duration_in_months - 3
			student.ErasmusData.NumberMonths = student.ErasmusData.NumberMonths - 3

			current_end_date, err := time.Parse(utilfunc.DeadlineLayout, erasmusCareer[lenCareer-1].End_date)
			if err != nil {
				return err
			}

			// as a test in extend erasmus 1500 seconds were added for each extend, here to do
			// the revert 1500 seconds are subtracted

			endDate := current_end_date.Add(time.Duration(-1500 * time.Second))
			erasmusCareer[lenCareer-1].End_date = endDate.Format(utilfunc.DeadlineLayout)

			resultByteJSON, err := json.Marshal(erasmusCareer)
			if err != nil {
				return err
			}

			erasmusJSON := string(resultByteJSON)

			student.ErasmusData.Career = erasmusJSON

			k.SetStoredStudent(ctx, student)

			uniInfo, found := k.GetUniversityInfo(ctx, student.StudentData.UniversityName)
			if !found {
				return types.ErrWrongNameUniversity
			} else {

				k.CheckAndInCaseMoveStudent(ctx, &student, &uniInfo)
				k.SetUniversityInfo(ctx, uniInfo)
			}

			err = k.ClearOperationQueue(ctx, &student)
			if err != nil {
				return err
			}

		}
	}
	utilfunc.PrintLogs("RevertExtendErasmus - end", ctx)

	return nil
}

// Function that allows the construction of the packet content that allows the abort
// operation related to the extend of the Erasmus

func (k Keeper) CreateAbortOperationStringExtendErasmus(student types.StoredStudent) (abort_op_JSON string, err error) {

	var abort_op utilfunc.AbortOperationPacket

	abort_op.HomeIndex, _ = utilfunc.GetForeignIndex(student)
	abort_op.ForeignIndex = student.Index
	abort_op.ForeignUniversity = student.StudentData.UniversityName
	abort_op.HomeUniversity, _ = utilfunc.GetForeignUniversityName(student)
	abort_op.PacketID = "-3"

	resultByteJSON, err := json.Marshal(abort_op)
	if err != nil {
		return abort_op_JSON, err
	}

	abort_op_JSON = string(resultByteJSON)

	return abort_op_JSON, err
}

// Function that constructs the packet contents and sends the packet related to the abort operation

func (k Keeper) HandleAbortPacketExtendErasmus(ctx sdk.Context, studentIndex string) (err error) {

	utilfunc.PrintLogs("HandleAbortPacketExtendErasmus", ctx)

	storedStudent, found := k.GetStoredStudent(ctx, studentIndex)
	if !found {
		return types.ErrStudentNotPresent
	} else {
	}

	new_data_info, err := k.CreateAbortOperationStringExtendErasmus(storedStudent)
	if err != nil {
		return err
	} else {

		err = k.SendAbortPacket(ctx, new_data_info)
		if err != nil {
			return err
		} else {

			if err != nil {
				utilfunc.PrintLogs("HandleAbortPacketExtendErasmus "+err.Error(), ctx)
				return err
			}
			return nil
		}

	}
}

// Function that allows the construction of the ack content that allows the abort
// operation related to the extend of the Erasmus

func (k Keeper) HandleAbortAckExtendErasmus(ctx sdk.Context, studentIndex string) (packetAck types.ExtendErasmusPeriodPacketAck, err error) {

	utilfunc.PrintLogs("HandleAbortAckExtendErasmus", ctx)

	storedStudent, found := k.GetStoredStudent(ctx, studentIndex)
	if !found {
		return packetAck, types.ErrStudentNotPresent
	} else {
	}

	new_data_info, err := k.CreateAbortOperationStringExtendErasmus(storedStudent)
	if err != nil {
		return packetAck, err
	} else {

		packetAck.ErasmusRestrictedInfo = new_data_info
		return packetAck, err
	}
}
