package keeper

import (
	"context"

	"university_chain_de/x/universitychainde/types"
	"university_chain_de/x/universitychainde/utilfunc"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) InsertErasmusExam(goCtx context.Context, msg *types.MsgInsertErasmusExam) (*types.MsgInsertErasmusExamResponse, error) {

	ctx := sdk.UnwrapSDKContext(goCtx)

	chainInfo, found := k.Keeper.GetChainInfo(ctx)
	if !found {
		panic("ChainInfo not found")
	} else {
		if !chainInfo.InitStatus {
			return &types.MsgInsertErasmusExamResponse{
				Status: -1,
			}, types.ErrChainConfigurationNotDone
		} else {

			uniInfo, found := k.Keeper.GetUniversityInfo(ctx, msg.University)
			if !found {
				return &types.MsgInsertErasmusExamResponse{
					Status: -1,
				}, types.ErrWrongNameUniversity
			} else {

				searchedStudent, found := k.Keeper.GetStoredStudent(ctx, msg.GetUniversity()+"_"+msg.GetStudentIndex())

				if !found {
					return &types.MsgInsertErasmusExamResponse{
						Status: -1,
					}, types.ErrStudentNotPresent
				} else {
					if searchedStudent.GetStudentData().GetStudentKey() != msg.Creator {
						return &types.MsgInsertErasmusExamResponse{
							Status: -1,
						}, types.ErrKeyEnteredMismatchStudent
					} else {

						if searchedStudent.ErasmusData.ErasmusStudent != "Incoming" && searchedStudent.ErasmusData.ErasmusStudent != "Incoming completed" {

							err := utilfunc.CheckCompleteInformation(searchedStudent)

							if err != nil {
								return &types.MsgInsertErasmusExamResponse{
									Status: -1,
								}, types.ErrIncompleteStudentInformation
							} else {

								ok, err := utilfunc.CheckTaxPayment(searchedStudent)
								//var ok bool = true
								//var err error = nil

								if err != nil {
									return &types.MsgInsertErasmusExamResponse{
										Status: -1,
									}, err
								} else {
									if !ok {
										return &types.MsgInsertErasmusExamResponse{
											Status: -1,
										}, types.ErrUnpaidTaxes
									} else {

										err := utilfunc.CheckErasmusStatus(searchedStudent, "insert erasmus exam")
										if err != nil {
											return &types.MsgInsertErasmusExamResponse{
												Status: -1,
											}, err
										} else {
											err := utilfunc.InsertErasmusExam(&searchedStudent, msg.ExamName, uniInfo.MaxErasmusExams)
											if err != nil {
												return &types.MsgInsertErasmusExamResponse{
													Status: -1,
												}, err
											} else {
												k.Keeper.SetStoredStudent(ctx, searchedStudent)

												err = utilfunc.GetConsumedGas("InsertErasmusExam DE", searchedStudent.Index, ctx)
												if err != nil {
													return &types.MsgInsertErasmusExamResponse{
														Status: -1,
													}, err
												} else {
													return &types.MsgInsertErasmusExamResponse{
														Status: 0,
													}, nil
												}
											}
										}
									}
								}

							}
						} else {
							err := utilfunc.CheckErasmusStatus(searchedStudent, "insert erasmus exam")
							if err != nil {
								return &types.MsgInsertErasmusExamResponse{
									Status: -1,
								}, err
							} else {

								err = utilfunc.GetConsumedGas("InsertErasmusExam DE", searchedStudent.Index, ctx)
								if err != nil {
									return &types.MsgInsertErasmusExamResponse{
										Status: -1,
									}, err
								} else {
									return &types.MsgInsertErasmusExamResponse{
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
