package keeper

import (
	"context"

	"university_chain_it/x/universitychainit/types"
	"university_chain_it/x/universitychainit/utilfunc"

	sdk "github.com/cosmos/cosmos-sdk/types"
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

									res, err := utilfunc.CheckErasmusStatus(searchedStudent)
									if err != nil {
										return &types.MsgStartErasmusResponse{
											Status: -1,
										}, err
									} else {
										if res == "" {
											return &types.MsgStartErasmusResponse{
												Status: -1,
											}, types.ErrNoErasmusRequest
										} else if res == "in progress" {
											return &types.MsgStartErasmusResponse{
												Status: -1,
											}, types.ErrPreviousRequestInProgress
										} else if res == "terminated" {
											return &types.MsgStartErasmusResponse{
												Status: -1,
											}, types.ErrPreviousRequestCompleted
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

													k.Keeper.SendToFifoTail(ctx, &searchedStudent, &uniInfo)

													_, err = k.SendErasmusStudent(goCtx, types.NewMsgSendErasmusStudent(
														msg.Creator, "hub", "channel-0", timeoutTimestamp, msg.GetUniversity()+"_"+msg.GetStudentIndex()))
													if err != nil {
														return nil, err
													} else {
														k.Keeper.SetStoredStudent(ctx, searchedStudent)
														k.Keeper.SetUniversityInfo(ctx, uniInfo)
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
				}
			}
		}
	}

}
