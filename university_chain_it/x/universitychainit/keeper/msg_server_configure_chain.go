package keeper

import (
	"context"
	"fmt"
	"os"

	"university_chain_it/x/universitychainit/types"
	"university_chain_it/x/universitychainit/utilfunc"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) ConfigureChain(goCtx context.Context, msg *types.MsgConfigureChain) (*types.MsgConfigureChainResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	chainInfo, found := k.Keeper.GetChainInfo(ctx)
	if !found {
		panic("ChainInfo not found")
	} else {
		if chainInfo.InitStatus {
			fmt.Fprintln(os.Stderr, "The chain configuration has already been performed.")
			return &types.MsgConfigureChainResponse{
				Status: -1,
			}, types.ErrChainConfigurationAlreadyDone
		} else {
			if msg.GetCreator() != chainInfo.ChainAdministratorKey {
				fmt.Fprintln(os.Stderr, "The key entered does not match the chain administrator's key.")
				return &types.MsgConfigureChainResponse{
					Status: -1,
				}, types.ErrKeyEnteredMismatchAdminChain
			} else {
				// set the flag to one in order to do the initialization just one time
				chainInfo.InitStatus = true
				k.Keeper.SetChainInfo(ctx, chainInfo)
				// get the universities name and key and insert these info into a map
				foreignUniversitiesList, err := utilfunc.ReadForeignUniversityInfo()
				if err != nil {
					return &types.MsgConfigureChainResponse{
						Status: -1,
					}, err
				} else {
					i := 0
					for i = 0; i < len(foreignUniversitiesList); i++ {

						foreignUniversity := types.ForeignUniversities{
							UniversityName:             foreignUniversitiesList[i].Name,
							ForeignUniversitiesCountry: foreignUniversitiesList[i].Country,
							ForeignUniversitiesKey:     foreignUniversitiesList[i].Address,
						}
						err = foreignUniversity.Validate()
						if err != nil {
							return &types.MsgConfigureChainResponse{
								Status: -1,
							}, err
						} else {
							k.Keeper.SetForeignUniversities(ctx, foreignUniversity)
						}
					}

					universityInfoList, err := utilfunc.ReadUniversitiesInfo()
					if err != nil {
						return &types.MsgConfigureChainResponse{
							Status: -1,
						}, err
					} else {

						for i = 0; i < len(universityInfoList); i++ {

							taxesBracketsJSON, err := utilfunc.GetJSONFromTaxesBrackets(universityInfoList[i].Taxes_brackets)
							if err != nil {
								return &types.MsgConfigureChainResponse{
									Status: -1,
								}, err
							} else {

								// to change in the future

								//deadlineErasmus := ctx.BlockTime().Add(time.Duration(3 * 30 * 24 * time.Hour))
								//deadlineTaxes := ctx.BlockTime().Add(time.Duration(3 * 30 * 24 * time.Hour))

								k.Keeper.SetUniversityInfo(ctx, types.UniversityInfo{
									UniversityName:  universityInfoList[i].Name,
									NextStudentId:   1,
									SecretariatKey:  universityInfoList[i].Secretariat_key,
									UniversityKey:   universityInfoList[i].University_key,
									CaiKey:          "",
									FifoHeadErasmus: "",
									FifoTailErasmus: "",
									DeadlineTaxes:   universityInfoList[i].Deadline_taxes,
									DeadlineErasmus: universityInfoList[i].Deadline_erasmus,
									//DeadlineTaxes:   utilfunc.FormatDeadline(deadlineTaxes),
									//DeadlineErasmus: utilfunc.FormatDeadline(deadlineErasmus),
									MaxErasmusExams: universityInfoList[i].Max_erasmus_exams,
									TaxesBrackets:   taxesBracketsJSON,
								})

								j := 0

								for j = 0; j < len(universityInfoList[i].DepartmentList); j++ {
									q := 0
									for q = 0; q < len(universityInfoList[i].DepartmentList[j].CoursesTypeList); q++ {
										w := 0
										for w = 0; w < len(universityInfoList[i].DepartmentList[j].CoursesTypeList[q].Courses); w++ {
											z := 0
											for z = 0; z < len(universityInfoList[i].DepartmentList[j].CoursesTypeList[q].Courses[w].Exams); z++ {

												k.Keeper.SetProfessorsExams(ctx, types.ProfessorsExams{
													ExamName:      universityInfoList[i].Name + "_" + universityInfoList[i].DepartmentList[j].CoursesTypeList[q].Courses[w].Exams[z].ExamName,
													ProfessorName: universityInfoList[i].DepartmentList[j].CoursesTypeList[q].Courses[w].Exams[z].ProfessorName,
													ProfessorId:   "",
													ProfessorKey:  universityInfoList[i].DepartmentList[j].CoursesTypeList[q].Courses[w].Exams[z].ProfessorAddress,
												})
											}

										}

									}
								}

							}
						}
						return &types.MsgConfigureChainResponse{
							Status: 0,
						}, nil
					}
				}
			}
		}
	}
}
