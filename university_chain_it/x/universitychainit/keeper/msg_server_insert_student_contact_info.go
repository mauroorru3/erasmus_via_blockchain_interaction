package keeper

import (
	"context"
	"fmt"
	"os"

	"university_chain_it/x/universitychainit/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) InsertStudentContactInfo(goCtx context.Context, msg *types.MsgInsertStudentContactInfo) (*types.MsgInsertStudentContactInfoResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	chainInfo, found := k.Keeper.GetChainInfo(ctx)
	if !found {
		panic("ChainInfo not found")
	} else {
		if !chainInfo.InitStatus {
			fmt.Fprintln(os.Stderr, "The initial configuration of the chain has not yet been performed.")
			return &types.MsgInsertStudentContactInfoResponse{
				Status: -1,
			}, types.ErrChainConfigurationNotDone
		} else {

			_, found := k.Keeper.GetUniversityInfo(ctx, msg.University)
			if !found {
				return &types.MsgInsertStudentContactInfoResponse{
					Status: -1,
				}, types.ErrWrongNameUniversity
			} else {

				searchedStudent, found := k.Keeper.GetStoredStudent(ctx, msg.GetUniversity()+"_"+msg.GetStudentIndex())
				if !found {
					return &types.MsgInsertStudentContactInfoResponse{
						Status: -1,
					}, types.ErrStudentNotPresent
				} else {
					if searchedStudent.GetStudentData().GetStudentKey() != msg.Creator {
						return &types.MsgInsertStudentContactInfoResponse{
							Status: -1,
						}, types.ErrKeyEnteredMismatchStudent
					} else {
						searchedStudent.ContactData = &types.ContactInfo{
							ContactAddress: msg.ContactAddress,
							Email:          msg.Email,
							MobilePhone:    msg.MobilePhone,
						}
						err := searchedStudent.ContactData.Validate()
						if err != nil {
							return &types.MsgInsertStudentContactInfoResponse{
								Status: -1,
							}, err
						} else {

							searchedStudent.StudentData.CompleteInformation[2] = 1
							k.Keeper.SetStoredStudent(ctx, searchedStudent)

							return &types.MsgInsertStudentContactInfoResponse{
								Status: 0,
							}, nil
						}
					}
				}
			}
		}
	}
}
