package keeper

import (
	"context"

	"university_chain_it/x/universitychainit/types"
	"university_chain_it/x/universitychainit/utilfunc"

	sdk "github.com/cosmos/cosmos-sdk/types"
	clienttypes "github.com/cosmos/ibc-go/v3/modules/core/02-client/types"
)

func (k msgServer) EndErasmusBeforeDeadline(goCtx context.Context, msg *types.MsgEndErasmusBeforeDeadline) (*types.MsgEndErasmusBeforeDeadlineResponse, error) {

	ctx := sdk.UnwrapSDKContext(goCtx)

	chainInfo, found := k.Keeper.GetChainInfo(ctx)
	if !found {
		panic("ChainInfo not found")
	} else {
		if !chainInfo.InitStatus {
			return &types.MsgEndErasmusBeforeDeadlineResponse{
				Status: -1,
			}, types.ErrChainConfigurationNotDone
		} else {

			uniInfo, found := k.Keeper.GetUniversityInfo(ctx, msg.University)
			if !found {
				return &types.MsgEndErasmusBeforeDeadlineResponse{
					Status: -1,
				}, types.ErrWrongNameUniversity
			} else {

				searchedStudent, found := k.Keeper.GetStoredStudent(ctx, msg.GetUniversity()+"_"+msg.GetStudentIndex())

				if !found {
					return &types.MsgEndErasmusBeforeDeadlineResponse{
						Status: -1,
					}, types.ErrStudentNotPresent
				} else {
					if searchedStudent.GetStudentData().GetStudentKey() != msg.Creator {
						return &types.MsgEndErasmusBeforeDeadlineResponse{
							Status: -1,
						}, types.ErrKeyEnteredMismatchStudent
					} else {

						if searchedStudent.ErasmusData.ErasmusStudent != "Incoming" && searchedStudent.ErasmusData.ErasmusStudent != "Incoming completed" {

							err := utilfunc.CheckCompleteInformation(searchedStudent)

							if err != nil {
								return &types.MsgEndErasmusBeforeDeadlineResponse{
									Status: -1,
								}, types.ErrIncompleteStudentInformation
							} else {

								ok, err := utilfunc.CheckTaxPayment(searchedStudent)

								if err != nil {
									return &types.MsgEndErasmusBeforeDeadlineResponse{
										Status: -1,
									}, err
								} else {
									if !ok {
										return &types.MsgEndErasmusBeforeDeadlineResponse{
											Status: -1,
										}, types.ErrUnpaidTaxes
									} else {

										err := utilfunc.CheckErasmusStatus(searchedStudent, "end erasmus before deadline")
										if err != nil {
											return &types.MsgEndErasmusBeforeDeadlineResponse{
												Status: -1,
											}, err
										} else {

											k.RemoveFromFifo(ctx, &searchedStudent, &uniInfo)

											k.SetStoredStudent(ctx, searchedStudent)
											k.SetUniversityInfo(ctx, uniInfo)

											var packet types.EndErasmusPeriodRequestPacketData

											packet.StartingUniversityName = msg.University
											packet.Index = searchedStudent.Index

											foreignUniversityName, err := utilfunc.GetForeignUniversityName(searchedStudent)
											if err != nil {
												return &types.MsgEndErasmusBeforeDeadlineResponse{
													Status: -1,
												}, err
											} else {

												packet.DestinationUniversityName = foreignUniversityName
												foreignIndex, err := utilfunc.GetForeignIndex(searchedStudent)

												if err != nil {
													return &types.MsgEndErasmusBeforeDeadlineResponse{
														Status: -1,
													}, err
												} else {

													packet.ForeignIndex = foreignIndex

													err = k.TransmitEndErasmusPeriodRequestPacket(
														ctx,
														packet,
														"universitychainit",
														"channel-0",
														clienttypes.ZeroHeight(),
														timeoutTimestamp,
														"EndErasmusBeforeDeadline",
													)
													if err != nil {
														return &types.MsgEndErasmusBeforeDeadlineResponse{
															Status: -1,
														}, err
													} else {
														err = utilfunc.GetConsumedGas("EndErasmusBeforeDeadline IT", searchedStudent.Index, ctx)
														if err != nil {
															return &types.MsgEndErasmusBeforeDeadlineResponse{
																Status: -1,
															}, err
														} else {

															return &types.MsgEndErasmusBeforeDeadlineResponse{
																Status: 0,
															}, nil
														}
													}
												}
											}
										}
									}
								}
							}
						} else {
							err := utilfunc.CheckErasmusStatus(searchedStudent, "end erasmus before deadline")
							if err != nil {
								return &types.MsgEndErasmusBeforeDeadlineResponse{
									Status: -1,
								}, err
							} else {

								err = utilfunc.GetConsumedGas("EndErasmusBeforeDeadline IT", searchedStudent.Index, ctx)
								if err != nil {
									return &types.MsgEndErasmusBeforeDeadlineResponse{
										Status: -1,
									}, err
								} else {
									return &types.MsgEndErasmusBeforeDeadlineResponse{
										Status: 0,
									}, nil
								}
							}
						}
					}
				}
			}
		}
	}
}
