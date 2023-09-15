package keeper

import (
	"context"
	"strconv"

	"university_chain_it/x/universitychainit/types"
	"university_chain_it/x/universitychainit/utilfunc"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) InsertExamGrade(goCtx context.Context, msg *types.MsgInsertExamGrade) (*types.MsgInsertExamGradeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	chainInfo, found := k.Keeper.GetChainInfo(ctx)
	if !found {
		panic("ChainInfo not found")
	} else {
		if !chainInfo.InitStatus {
			return &types.MsgInsertExamGradeResponse{
				Status: -1,
			}, types.ErrChainConfigurationNotDone
		} else {

			_, found := k.GetUniversityInfo(ctx, msg.University)
			if !found {
				return &types.MsgInsertExamGradeResponse{
					Status: -1,
				}, types.ErrWrongNameUniversity
			} else {

				examInfo, found := k.Keeper.GetProfessorsExams(ctx, msg.University+"_"+msg.ExamName)
				if !found {
					return &types.MsgInsertExamGradeResponse{
						Status: -1,
					}, types.ErrWrongExamName
				} else {
					if msg.Creator != examInfo.ProfessorKey {
						return &types.MsgInsertExamGradeResponse{
							Status: -1,
						}, types.ErrUnauthorisedUser
					} else {

						searchedStudent, found := k.Keeper.GetStoredStudent(ctx, msg.GetUniversity()+"_"+msg.GetStudentIndex())
						if !found {
							return &types.MsgInsertExamGradeResponse{
								Status: -1,
							}, types.ErrStudentNotPresent
						} else {
							if searchedStudent.ErasmusData.ErasmusStudent != "Incoming" && searchedStudent.ErasmusData.ErasmusStudent != "Incoming completed" {
								err := utilfunc.CheckCompleteInformation(searchedStudent)

								if err != nil {
									return &types.MsgInsertExamGradeResponse{
										Status: -1,
									}, types.ErrIncompleteStudentInformation
								} else {
									ok, err := utilfunc.CheckTaxPayment(searchedStudent)
									if err != nil {
										return &types.MsgInsertExamGradeResponse{
											Status: -1,
										}, err
									} else {
										if !ok {
											return &types.MsgInsertExamGradeResponse{
												Status: -1,
											}, types.ErrUnpaidTaxes
										} else {

											err := utilfunc.CheckErasmusStatus(searchedStudent, "insert exam grade")
											if err != nil {
												return &types.MsgInsertExamGradeResponse{
													Status: -1,
												}, err
											} else {

												gradeNum, err := strconv.ParseUint(msg.Grade, 10, 8)
												if err != nil {
													if msg.Grade != "30L" {
														return &types.MsgInsertExamGradeResponse{
															Status: -1,
														}, types.ErrWrongExamGrade
													}
												} else {
													if gradeNum > 30 || gradeNum < 18 {
														return &types.MsgInsertExamGradeResponse{
															Status: -1,
														}, types.ErrWrongExamGrade
													}
												}
												var credits uint8 = 0
												var JSONExams string = ""
												JSONExams, credits, err = utilfunc.SetExamGrade(searchedStudent.TranscriptData.ExamsData, msg.ExamName, msg.Grade)
												if err != nil {
													return &types.MsgInsertExamGradeResponse{
														Status: -1,
													}, err
												} else {
													searchedStudent.TranscriptData.ExamsData = JSONExams

												}

												searchedStudent.TranscriptData.AchievedCredits += uint32(credits)
												searchedStudent.TranscriptData.ExamsPassed += 1

												k.Keeper.SetStoredStudent(ctx, searchedStudent)

												err = utilfunc.GetConsumedGas("InsertExamGrade IT", searchedStudent.Index, ctx)
												if err != nil {
													return &types.MsgInsertExamGradeResponse{
														Status: -1,
													}, err
												} else {

													return &types.MsgInsertExamGradeResponse{
														Status: 0,
													}, nil

												}
											}
										}

									}
								}
							} else {
								err := utilfunc.CheckErasmusStatus(searchedStudent, "insert exam grade")
								if err != nil {
									return &types.MsgInsertExamGradeResponse{
										Status: -1,
									}, err
								} else {

									gradeNum, err := strconv.ParseUint(msg.Grade, 10, 8)
									if err != nil {
										if msg.Grade != "30L" {
											return &types.MsgInsertExamGradeResponse{
												Status: -1,
											}, types.ErrWrongExamGrade
										}
									} else {
										if gradeNum > 30 || gradeNum < 18 {
											return &types.MsgInsertExamGradeResponse{
												Status: -1,
											}, types.ErrWrongExamGrade
										}
									}
									var credits uint8 = 0
									credits, err = utilfunc.SetErasmusExamGrade(&searchedStudent, msg.ExamName, msg.Grade)
									if err != nil {
										return &types.MsgInsertExamGradeResponse{
											Status: -1,
										}, err
									}

									searchedStudent.TranscriptData.AchievedCredits += uint32(credits)
									searchedStudent.TranscriptData.ExamsPassed += 1

									k.Keeper.SetStoredStudent(ctx, searchedStudent)

									err = utilfunc.GetConsumedGas("InsertExamGrade IT", searchedStudent.Index, ctx)
									if err != nil {
										return &types.MsgInsertExamGradeResponse{
											Status: -1,
										}, err
									} else {

										return &types.MsgInsertExamGradeResponse{
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
