package keeper

import (
	"context"

	"university_chain_it/x/universitychainit/types"
	"university_chain_it/x/universitychainit/utilfunc"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) InsertErasmusRequest(goCtx context.Context, msg *types.MsgInsertErasmusRequest) (*types.MsgInsertErasmusRequestResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	chainInfo, found := k.Keeper.GetChainInfo(ctx)
	if !found {
		panic("ChainInfo not found")
	} else {
		if !chainInfo.InitStatus {
			return &types.MsgInsertErasmusRequestResponse{
				Status: -1,
			}, types.ErrChainConfigurationNotDone
		} else {

			uniInfo, found := k.GetUniversityInfo(ctx, msg.University)
			if !found {
				return &types.MsgInsertErasmusRequestResponse{
					Status: -1,
				}, types.ErrWrongNameUniversity
			} else {

				searchedStudent, found := k.Keeper.GetStoredStudent(ctx, msg.GetUniversity()+"_"+msg.GetStudentIndex())

				if !found {
					return &types.MsgInsertErasmusRequestResponse{
						Status: -1,
					}, types.ErrStudentNotPresent
				} else {
					if searchedStudent.GetStudentData().GetStudentKey() != msg.Creator {
						return &types.MsgInsertErasmusRequestResponse{
							Status: -1,
						}, types.ErrKeyEnteredMismatchStudent
					} else {

						err := utilfunc.CheckCompleteInformation(searchedStudent)

						if err != nil {
							return &types.MsgInsertErasmusRequestResponse{
								Status: -1,
							}, types.ErrIncompleteStudentInformation
						} else {

							ok, err := utilfunc.CheckTaxPayment(searchedStudent)
							//var ok bool = true
							//var err error = nil

							if err != nil {
								return &types.MsgInsertErasmusRequestResponse{
									Status: -1,
								}, err
							} else {
								if !ok {
									return &types.MsgInsertErasmusRequestResponse{
										Status: -1,
									}, types.ErrUnpaidTaxes
								} else {

									foreignUni, found := k.Keeper.GetForeignUniversities(ctx, msg.ForeignUniversityName)
									if !found {
										return &types.MsgInsertErasmusRequestResponse{
											Status: -1,
										}, types.ErrWrongForeignUniversity
									} else {

										res, err := utilfunc.CheckErasmusStatus(searchedStudent)
										if err != nil {
											return &types.MsgInsertErasmusRequestResponse{
												Status: -1,
											}, err
										} else {
											if res == "in progress" {
												return &types.MsgInsertErasmusRequestResponse{
													Status: -1,
												}, types.ErrPreviousRequestInProgress
											} else if res == "to start" {
												return &types.MsgInsertErasmusRequestResponse{
													Status: -1,
												}, types.ErrPreviousRequestStartup
											} else {

												err := utilfunc.CheckErasmusDeadline(ctx, uniInfo.DeadlineErasmus)
												if err != nil {
													return &types.MsgInsertErasmusRequestResponse{
														Status: -1,
													}, err
												} else {

													err = utilfunc.CheckErasmusParams(msg.DurationInMonths, msg.ErasmusType, &searchedStudent, msg.ForeignUniversityName, foreignUni.ForeignUniversitiesCountry)
													if err != nil {
														return &types.MsgInsertErasmusRequestResponse{
															Status: -1,
														}, err
													} else {
														k.Keeper.SetStoredStudent(ctx, searchedStudent)
														return &types.MsgInsertErasmusRequestResponse{
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
