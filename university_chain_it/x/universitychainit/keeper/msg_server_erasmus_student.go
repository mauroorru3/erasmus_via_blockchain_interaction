package keeper

import (
	"context"

	"university_chain_it/x/universitychainit/types"
	"university_chain_it/x/universitychainit/utilfunc"

	sdk "github.com/cosmos/cosmos-sdk/types"
	clienttypes "github.com/cosmos/ibc-go/v3/modules/core/02-client/types"
)

func (k msgServer) SendErasmusStudent(goCtx context.Context, msg *types.MsgSendErasmusStudent) (*types.MsgSendErasmusStudentResponse, error) {

	ctx := sdk.UnwrapSDKContext(goCtx)

	var packet types.ErasmusStudentPacketData

	stu, found := k.GetStoredStudent(ctx, msg.Index)
	if !found {
		return &types.MsgSendErasmusStudentResponse{
			Status: -1,
		}, types.ErrStudentNotPresent
	} else {

		val := types.StoredStudent{
			Index: stu.Index,
			StudentData: &types.StudentInfo{
				Name:                     stu.StudentData.Name,
				Surname:                  stu.StudentData.Surname,
				CourseType:               stu.StudentData.CourseType,
				CourseOfStudy:            stu.StudentData.CourseOfStudy,
				Status:                   stu.StudentData.Status,
				CurrentYearOfStudy:       stu.StudentData.CurrentYearOfStudy,
				OutOfCourse:              stu.StudentData.OutOfCourse,
				NumberOfYearsOutOfCourse: stu.StudentData.NumberOfYearsOutOfCourse,
				StudentKey:               stu.StudentData.StudentKey,
				CompleteInformation:      stu.StudentData.CompleteInformation,
				UniversityName:           stu.StudentData.UniversityName,
				//ChainName:                stu.StudentData.ChainName,
			},
			//TranscriptData: stu.TranscriptData,
			//PersonalData:   stu.PersonalData,
			//ResidenceData:  stu.ResidenceData,
			//ContactData:    stu.ContactData,
			//TaxesData:      stu.TaxesData,
			//ErasmusData: stu.ErasmusData,
		}

		packet.Student = &val

		//packet.Student = &stu

		// Transmit the packet
		err := k.TransmitErasmusStudentPacket(
			ctx,
			packet,
			msg.Port,
			"channel-0",
			clienttypes.ZeroHeight(),
			timeoutTimestamp,
		)

		if err != nil {
			utilfunc.PrintLogs("SendErasmusStudent " + err.Error())
			return &types.MsgSendErasmusStudentResponse{
				Status: -1,
			}, err
		} else {
			utilfunc.PrintLogs("SendErasmusStudent packet sent")
			return &types.MsgSendErasmusStudentResponse{
				Status: -1,
			}, nil
		}
	}
}
