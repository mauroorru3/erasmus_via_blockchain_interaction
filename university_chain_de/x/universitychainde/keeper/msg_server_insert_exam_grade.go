package keeper

import (
	"context"
	"strconv"
	"strings"

	"university_chain_de/x/universitychainde/types"
	"university_chain_de/x/universitychainde/utilfunc"

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
							if searchedStudent.ErasmusData.ErasmusStudent != "Incoming" && searchedStudent.ErasmusData.ErasmusStudent != "Incoming completed" {
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

											err := utilfunc.CheckErasmusStatus(searchedStudent, "insert exam grade")
											if err != nil {
												return &types.MsgInsertExamGradeResponse{
													Status: -1,
												}, err
											} else {

												extendedGrade := strings.Split(msg.Grade, ",")
												if len(extendedGrade) > 1 {
													if len(extendedGrade[1]) > 1 {
														return &types.MsgInsertExamGradeResponse{
															Status: -1,
														}, types.ErrWrongExamGrade
													}
												}
												extendedGrade = strings.Split(msg.Grade, ".")
												if len(extendedGrade) > 1 {
													if len(extendedGrade[1]) > 1 {
														return &types.MsgInsertExamGradeResponse{
															Status: -1,
														}, types.ErrWrongExamGrade
													}
												}

												gradeNum, err := strconv.ParseFloat(msg.Grade, 32)
												if err != nil {
													return &types.MsgInsertExamGradeResponse{
														Status: -1,
													}, types.ErrWrongExamGrade

												} else {
													if gradeNum > 4 || gradeNum < 1 {
														return &types.MsgInsertExamGradeResponse{
															Status: -1,
														}, types.ErrWrongExamGrade
													}
												}

												var credits uint8 = 0
												var JSONExams string = ""
												JSONExams, credits, err = utilfunc.SetExamGrade(searchedStudent.TranscriptData.ExamsData, msg.ExamName, strconv.FormatFloat(gradeNum, 'f', 1, 64))
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

												return &types.MsgInsertExamGradeResponse{
													Status: 0,
												}, nil

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

									extendedGrade := strings.Split(msg.Grade, ",")
									if len(extendedGrade) > 1 {
										if len(extendedGrade[1]) > 1 {
											return &types.MsgInsertExamGradeResponse{
												Status: -1,
											}, types.ErrWrongExamGrade
										}
									}
									extendedGrade = strings.Split(msg.Grade, ".")
									if len(extendedGrade) > 1 {
										if len(extendedGrade[1]) > 1 {
											return &types.MsgInsertExamGradeResponse{
												Status: -1,
											}, types.ErrWrongExamGrade
										}
									}

									gradeNum, err := strconv.ParseFloat(msg.Grade, 32)
									if err != nil {
										return &types.MsgInsertExamGradeResponse{
											Status: -1,
										}, types.ErrWrongExamGrade

									} else {
										if gradeNum > 4 || gradeNum < 1 {
											return &types.MsgInsertExamGradeResponse{
												Status: -1,
											}, types.ErrWrongExamGrade
										}
									}

									var credits uint8 = 0
									credits, err = utilfunc.SetErasmusExamGrade(&searchedStudent, msg.ExamName, strconv.FormatFloat(gradeNum, 'f', 1, 64))
									if err != nil {
										return &types.MsgInsertExamGradeResponse{
											Status: -1,
										}, err
									}

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
