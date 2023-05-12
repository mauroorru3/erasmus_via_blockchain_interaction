package keeper

import (
	"context"
	"fmt"
	"os"

	"university_chain_it/x/universitychainit/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) InsertStudentResidenceInfo(goCtx context.Context, msg *types.MsgInsertStudentResidenceInfo) (*types.MsgInsertStudentResidenceInfoResponse, error) {

	ctx := sdk.UnwrapSDKContext(goCtx)

	chainInfo, found := k.Keeper.GetChainInfo(ctx)
	if !found {
		panic("ChainInfo not found")
	} else {
		if !chainInfo.InitStatus {
			fmt.Fprintln(os.Stderr, "The initial configuration of the chain has not yet been performed.")
			return &types.MsgInsertStudentResidenceInfoResponse{
				Status: -1,
			}, types.ErrChainConfigurationNotDone
		} else {

			_, found := k.Keeper.GetUniversityInfo(ctx, msg.University)
			if !found {
				return &types.MsgInsertStudentResidenceInfoResponse{
					Status: -1,
				}, types.ErrWrongNameUniversity
			} else {

				searchedStudent, found := k.Keeper.GetStoredStudent(ctx, msg.GetUniversity()+"_"+msg.GetStudentIndex())
				if !found {
					return &types.MsgInsertStudentResidenceInfoResponse{
						Status: -1,
					}, types.ErrStudentNotPresent
				} else {

					if searchedStudent.GetStudentData().GetStudentKey() != msg.Creator {
						return &types.MsgInsertStudentResidenceInfoResponse{
							Status: -1,
						}, types.ErrKeyEnteredMismatchStudent
					} else {
						searchedStudent.ResidenceData = &types.ResidenceInfo{
							Country:     msg.Country,
							Province:    msg.Province,
							Town:        msg.Town,
							PostCode:    msg.PostCode,
							Address:     msg.Address,
							HouseNumber: msg.HouseNumber,
							HomePhone:   msg.HomePhone,
						}
						err := searchedStudent.ResidenceData.Validate()
						if err != nil {
							return &types.MsgInsertStudentResidenceInfoResponse{
								Status: -1,
							}, err
						} else {

							searchedStudent.StudentData.CompleteInformation[1] = 1

							k.Keeper.SetStoredStudent(ctx, searchedStudent)

							return &types.MsgInsertStudentResidenceInfoResponse{
								Status: 0,
							}, nil

						}
					}

				}
			}
		}
	}
}
