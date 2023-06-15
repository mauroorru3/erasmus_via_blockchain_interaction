package keeper

import (
	"context"
	"time"
	"university_chain_de/x/universitychainde/types"
	"university_chain_de/x/universitychainde/utilfunc"

	sdk "github.com/cosmos/cosmos-sdk/types"
	clienttypes "github.com/cosmos/ibc-go/v3/modules/core/02-client/types"
)

func (k Keeper) TerminateExpiredErasmusPeriods(goCtx context.Context) {

	ctx := sdk.UnwrapSDKContext(goCtx)

	uniList := k.GetAllUniversityInfo(ctx)
	lenUniList := len(uniList)

	for i := 0; i < lenUniList; i++ {

		studentIndex := uniList[i].FifoHeadErasmus
		finish := false
		count := 0
		for !finish {
			// Finished moving along
			if studentIndex == "" {
				finish = true
			} else {
				storedStudent, found := k.GetStoredStudent(ctx, studentIndex)

				if !found {
					panic("Fifo head game not found " + uniList[i].FifoHeadErasmus)
				}
				deadline, err := utilfunc.GetErasmusDeadline(storedStudent)
				if err != nil {
					panic(err)
				}

				s := utilfunc.FormatDeadline(ctx.BlockTime())
				formattedStartDate, _ := time.Parse(utilfunc.DeadlineLayout, s)

				//utilfunc.PrintLogs("current time " + formattedStartDate.String())
				//utilfunc.PrintLogs("deadline " + deadline.String())

				if deadline.Before(formattedStartDate) {
					//if deadline.Before(ctx.BlockTime()) {

					// Erasmus period is past deadline

					utilfunc.PrintLogs("inside")

					k.RemoveFromFifo(ctx, &storedStudent, &uniList[i])

					k.SetStoredStudent(ctx, storedStudent)

					var packet types.EndErasmusPeriodRequestPacketData

					packet.StartingUniversityName = storedStudent.StudentData.UniversityName
					packet.Index = storedStudent.Index

					foreignUniversityName, err := utilfunc.GetForeignUniversityName(storedStudent)
					if err != nil {
						panic(err)
					} else {

						packet.DestinationUniversityName = foreignUniversityName
						foreignIndex, err := utilfunc.GetForeignIndex(storedStudent)

						if err != nil {
							panic(err)
						} else {

							packet.ForeignIndex = foreignIndex

							/*

								utilfunc.PrintLogs("prima di SendEndErasmusPeriodRequest")

								_, err := msgServer.SendEndErasmusPeriodRequest(goCtx, &types.MsgSendEndErasmusPeriodRequest{
									Creator:                   storedStudent.StudentData.StudentKey,
									Port:                      "hub",
									ChannelID:                 "channel-0",
									TimeoutTimestamp:          timeoutTimestamp,
									StartingUniversityName:    storedStudent.StudentData.UniversityName,
									DestinationUniversityName: foreignUniversityName,
									Index:                     storedStudent.Index,
									ForeignIndex:              foreignIndex,
								})

								if err != nil {
									utilfunc.PrintLogs("SendEndErasmusPeriodRequest " + err.Error())
									panic(err)
								} else {

							*/

							utilfunc.PrintLogs("TransmitEndErasmusPeriodRequestPacket " + packet.ForeignIndex)
							utilfunc.PrintLogs("TransmitEndErasmusPeriodRequestPacket " + packet.DestinationUniversityName)

							err = k.TransmitEndErasmusPeriodRequestPacket(
								ctx,
								packet,
								"hub",
								"channel-0",
								clienttypes.ZeroHeight(),
								timeoutTimestamp,
							)
							if err != nil {
								panic(err)
							}

							utilfunc.PrintLogs("TransmitEndErasmusPeriodRequestPacket packet sent")

							// Move along FIFO
							studentIndex = uniList[i].FifoHeadErasmus

							count++

						}
					}

				} else {
					finish = true
				}

			}
		}
		if count > 0 {
			k.SetUniversityInfo(ctx, uniList[i])
		}

	}

}
