package keeper

import (
	"context"

	"university_chain_de/x/universitychainde/types"
	"university_chain_de/x/universitychainde/utilfunc"

	sdk "github.com/cosmos/cosmos-sdk/types"
	clienttypes "github.com/cosmos/ibc-go/v3/modules/core/02-client/types"
)

func (k msgServer) ExtendErasmus(goCtx context.Context, msg *types.MsgExtendErasmus) (*types.MsgExtendErasmusResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	chainInfo, found := k.Keeper.GetChainInfo(ctx)
	if !found {
		panic("ChainInfo not found")
	} else {
		if !chainInfo.InitStatus {
			return &types.MsgExtendErasmusResponse{
				Status: -1,
			}, types.ErrChainConfigurationNotDone
		} else {

			uniInfo, found := k.GetUniversityInfo(ctx, msg.University)
			if !found {
				return &types.MsgExtendErasmusResponse{
					Status: -1,
				}, types.ErrWrongNameUniversity
			} else {
				searchedStudent, found := k.Keeper.GetStoredStudent(ctx, msg.GetUniversity()+"_"+msg.GetStudentIndex())

				if !found {
					return &types.MsgExtendErasmusResponse{
						Status: -1,
					}, types.ErrStudentNotPresent
				} else {
					if searchedStudent.GetStudentData().GetStudentKey() != msg.Creator {
						return &types.MsgExtendErasmusResponse{
							Status: -1,
						}, types.ErrKeyEnteredMismatchStudent
					} else {

						err := utilfunc.CheckCompleteInformation(searchedStudent)

						if err != nil {
							return &types.MsgExtendErasmusResponse{
								Status: -1,
							}, types.ErrIncompleteStudentInformation
						} else {

							ok, err := utilfunc.CheckTaxPayment(searchedStudent)

							if err != nil {
								return &types.MsgExtendErasmusResponse{
									Status: -1,
								}, err
							} else {
								if !ok {
									return &types.MsgExtendErasmusResponse{
										Status: -1,
									}, types.ErrUnpaidTaxes
								} else {

									err := utilfunc.CheckErasmusStatus(searchedStudent, "extend erasmus")
									if err != nil {
										return &types.MsgExtendErasmusResponse{
											Status: -1,
										}, err
									} else {

										err := utilfunc.CheckErasmusDeadline(ctx, uniInfo.DeadlineErasmus)
										if err != nil {
											return &types.MsgExtendErasmusResponse{
												Status: -1,
											}, err
										} else {

											additionalDuration, finalDate, err := utilfunc.ExtendErasmus(ctx, msg.DurationInMonths, &searchedStudent)
											if err != nil {
												return &types.MsgExtendErasmusResponse{
													Status: -1,
												}, err
											} else {

												var packet types.ExtendErasmusPeriodPacketData

												foreignUni, err := utilfunc.GetForeignUniversityName(searchedStudent)
												if err != nil {
													return &types.MsgExtendErasmusResponse{
														Status: -1,
													}, err
												} else {

													foreignIndex, err := utilfunc.GetForeignIndex(searchedStudent)
													if err != nil {
														return &types.MsgExtendErasmusResponse{
															Status: -1,
														}, err
													} else {

														packet.DestinationUniversityName = foreignUni
														packet.ForeignIndex = foreignIndex
														packet.DurationInMonths = uint32(additionalDuration)
														packet.FinalDate = finalDate

														// Transmit the packet
														err = k.TransmitExtendErasmusPeriodPacket(
															ctx,
															packet,
															"universitychainde",
															"channel-0",
															clienttypes.ZeroHeight(),
															timeoutTimestamp,
														)
														if err != nil {
															return nil, err
														} else {

															k.CheckAndInCaseMoveStutent(ctx, &searchedStudent, &uniInfo)
															k.Keeper.SetStoredStudent(ctx, searchedStudent)
															k.Keeper.SetUniversityInfo(ctx, uniInfo)
															return &types.MsgExtendErasmusResponse{
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
					}
				}
			}
		}
	}
}
