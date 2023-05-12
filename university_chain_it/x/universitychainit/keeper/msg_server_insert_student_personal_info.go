package keeper

import (
	"context"
	"fmt"
	"os"

	"university_chain_it/x/universitychainit/types"
	"university_chain_it/x/universitychainit/utilfunc"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) InsertStudentPersonalInfo(goCtx context.Context, msg *types.MsgInsertStudentPersonalInfo) (*types.MsgInsertStudentPersonalInfoResponse, error) {

	ctx := sdk.UnwrapSDKContext(goCtx)

	chainInfo, found := k.Keeper.GetChainInfo(ctx)
	if !found {
		panic("ChainInfo not found")
	} else {
		if !chainInfo.InitStatus {
			fmt.Fprintln(os.Stderr, "The initial configuration of the chain has not yet been performed.")
			return &types.MsgInsertStudentPersonalInfoResponse{
				Status: -1,
			}, types.ErrChainConfigurationNotDone
		} else {

			uniInfo, found := k.Keeper.GetUniversityInfo(ctx, msg.University)
			if !found {
				return &types.MsgInsertStudentPersonalInfoResponse{
					Status: -1,
				}, types.ErrWrongNameUniversity
			} else {

				searchedStudent, found := k.Keeper.GetStoredStudent(ctx, msg.GetUniversity()+"_"+msg.GetStudentIndex())

				if !found {
					return &types.MsgInsertStudentPersonalInfoResponse{
						Status: -1,
					}, types.ErrStudentNotPresent
				} else {

					if searchedStudent.GetStudentData().GetStudentKey() != msg.Creator {
						return &types.MsgInsertStudentPersonalInfoResponse{
							Status: -1,
						}, types.ErrKeyEnteredMismatchStudent
					} else {
						searchedStudent.PersonalData = &types.PersonalInfo{
							Gender:             msg.Gender,
							DateOfBirth:        msg.DateOfBirth,
							PrimaryNationality: msg.PrimaryNationality,
							CountryOfBirth:     msg.CountryOfBirth,
							ProvinceOfBirth:    msg.ProvinceOfBirth,
							TownOfBirth:        msg.TownOfBirth,
							TaxCode:            msg.TaxCode,
						}
						err := searchedStudent.PersonalData.Validate()
						if err != nil {
							return &types.MsgInsertStudentPersonalInfoResponse{
								Status: -1,
							}, err
						} else {

							var taxesData, erasmusData string
							var err error

							taxesData, err = utilfunc.IntializeTaxesStruct(msg.IncomeBracket, uniInfo.TaxesBrackets)
							if err != nil {
								return &types.MsgInsertStudentPersonalInfoResponse{
									Status: -1,
								}, err
							} else {
								erasmusData, err = utilfunc.IntializeErasmusStruct(msg.IncomeBracket)
								if err != nil {
									return &types.MsgInsertStudentPersonalInfoResponse{
										Status: -1,
									}, err
								} else {

									searchedStudent.TaxesData.TaxesHistory = taxesData
									searchedStudent.ErasmusData.Career = erasmusData
									searchedStudent.StudentData.CompleteInformation[0] = 1

									k.Keeper.SetStoredStudent(ctx, searchedStudent)

									return &types.MsgInsertStudentPersonalInfoResponse{
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
