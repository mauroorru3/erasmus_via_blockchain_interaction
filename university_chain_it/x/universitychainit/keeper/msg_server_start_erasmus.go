package keeper

import (
	"context"

	"university_chain_it/x/universitychainit/types"
	"university_chain_it/x/universitychainit/utilfunc"

	sdk "github.com/cosmos/cosmos-sdk/types"
	clienttypes "github.com/cosmos/ibc-go/v3/modules/core/02-client/types"
)

const timeoutTimestamp uint64 = 17999999999999999999 //1683974783938252484

func (k msgServer) StartErasmus(goCtx context.Context, msg *types.MsgStartErasmus) (*types.MsgStartErasmusResponse, error) {

	ctx := sdk.UnwrapSDKContext(goCtx)

	chainInfo, found := k.Keeper.GetChainInfo(ctx)
	if !found {
		panic("ChainInfo not found")
	} else {
		if !chainInfo.InitStatus {
			return &types.MsgStartErasmusResponse{
				Status: -1,
			}, types.ErrChainConfigurationNotDone
		} else {

			uniInfo, found := k.Keeper.GetUniversityInfo(ctx, msg.University)
			if !found {
				return &types.MsgStartErasmusResponse{
					Status: -1,
				}, types.ErrWrongNameUniversity
			} else {

				searchedStudent, found := k.Keeper.GetStoredStudent(ctx, msg.GetUniversity()+"_"+msg.GetStudentIndex())

				if !found {
					return &types.MsgStartErasmusResponse{
						Status: -1,
					}, types.ErrStudentNotPresent
				} else {
					if searchedStudent.GetStudentData().GetStudentKey() != msg.Creator {
						return &types.MsgStartErasmusResponse{
							Status: -1,
						}, types.ErrKeyEnteredMismatchStudent
					} else {

						if searchedStudent.ErasmusData.ErasmusStudent != "Incoming" && searchedStudent.ErasmusData.ErasmusStudent != "Incoming completed" {

							err := utilfunc.CheckCompleteInformation(searchedStudent)

							if err != nil {
								return &types.MsgStartErasmusResponse{
									Status: -1,
								}, types.ErrIncompleteStudentInformation
							} else {

								ok, err := utilfunc.CheckTaxPayment(searchedStudent)
								//var ok bool = true
								//var err error = nil

								if err != nil {
									return &types.MsgStartErasmusResponse{
										Status: -1,
									}, err
								} else {
									if !ok {
										return &types.MsgStartErasmusResponse{
											Status: -1,
										}, types.ErrUnpaidTaxes
									} else {

										err := utilfunc.CheckErasmusStatus(searchedStudent, "start erasmus")
										if err != nil {
											return &types.MsgStartErasmusResponse{
												Status: -1,
											}, err
										} else {

											err := utilfunc.StartErasmus(ctx, &searchedStudent, &uniInfo)
											if err != nil {
												return &types.MsgStartErasmusResponse{
													Status: -1,
												}, err
											} else {

												err := k.Keeper.CollectAndSendErasmusContribution(ctx, &searchedStudent, uniInfo.UniversityKey)
												if err != nil {
													return &types.MsgStartErasmusResponse{
														Status: -1,
													}, err
												} else {

													k.Keeper.InsertInTheErasmusFIFOQueue(ctx, &searchedStudent, &uniInfo)

													var packet types.ErasmusRestictedDataPacketData

													data, err := utilfunc.CreateHomeIndexJSONPacketFromStudentData(searchedStudent)
													if err != nil {
														return &types.MsgStartErasmusResponse{
															Status: -1,
														}, err
													}

													packet.ErasmusRestrictedInfo = data

													err = k.TransmitErasmusRestictedDataPacket(
														ctx,
														packet,
														"universitychainit",
														"channel-0",
														clienttypes.ZeroHeight(),
														timeoutTimestamp,
														" StartErasmus",
													)
													if err != nil {
														utilfunc.PrintLogs("TransmitErasmusStudentPacket " + err.Error())
														return nil, err
													} else {
														utilfunc.PrintLogs("TransmitErasmusStudentPacket packet sent")
														k.Keeper.SetStoredStudent(ctx, searchedStudent)
														k.Keeper.SetUniversityInfo(ctx, uniInfo)

														err = utilfunc.GetConsumedGas("StartErasmus IT", searchedStudent.Index, ctx)
														if err != nil {
															return &types.MsgStartErasmusResponse{
																Status: -1,
															}, err
														} else {

															return &types.MsgStartErasmusResponse{
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
							err := utilfunc.CheckErasmusStatus(searchedStudent, "start erasmus")
							if err != nil {
								return &types.MsgStartErasmusResponse{
									Status: -1,
								}, err
							} else {
								err = utilfunc.GetConsumedGas("StartErasmus IT", searchedStudent.Index, ctx)
								if err != nil {
									return &types.MsgStartErasmusResponse{
										Status: -1,
									}, err
								} else {
									return &types.MsgStartErasmusResponse{
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
