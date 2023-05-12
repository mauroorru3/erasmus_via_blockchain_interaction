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

			_, found = k.GetUniversityInfo(ctx, msg.University)
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
							err := utilfunc.CheckCompleteInformation(searchedStudent)

							if err != nil {
								return &types.MsgInsertExamGradeResponse{
									Status: -1,
								}, types.ErrIncompleteStudentInformation
							} else {
								ok, err := utilfunc.CheckTaxPayment(searchedStudent)
								//var ok bool = true
								//var err error = nil

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

										JSONExams, credits, err := utilfunc.SetExamGrade(searchedStudent.TranscriptData.ExamsData, msg.ExamName, msg.Grade)
										if err != nil {
											return &types.MsgInsertExamGradeResponse{
												Status: -1,
											}, err
										} else {
											searchedStudent.TranscriptData.ExamsData = JSONExams
											searchedStudent.TranscriptData.AchievedCredits += uint32(credits)
											searchedStudent.TranscriptData.ExamsPassed += 1
											k.Keeper.SetStoredStudent(ctx, searchedStudent)

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
}
