package keeper

import (
	"context"
	"fmt"
	"os"

	"university_chain_it/x/universitychainit/types"
	"university_chain_it/x/universitychainit/utilfunc"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) PayTaxes(goCtx context.Context, msg *types.MsgPayTaxes) (*types.MsgPayTaxesResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	chainInfo, found := k.Keeper.GetChainInfo(ctx)
	if !found {
		panic("ChainInfo not found")
	} else {
		if !chainInfo.InitStatus {
			fmt.Fprintln(os.Stderr, "The initial configuration of the chain has not yet been performed.")
			return &types.MsgPayTaxesResponse{
				Status: -1,
			}, types.ErrChainConfigurationNotDone
		} else {

			uniInfo, found := k.Keeper.GetUniversityInfo(ctx, msg.University)
			if !found {
				return &types.MsgPayTaxesResponse{
					Status: -1,
				}, types.ErrWrongNameUniversity
			} else {

				searchedStudent, found := k.Keeper.GetStoredStudent(ctx, msg.GetUniversity()+"_"+msg.GetStudentIndex())
				if !found {
					return &types.MsgPayTaxesResponse{
						Status: -1,
					}, types.ErrStudentNotPresent
				} else {

					if searchedStudent.GetStudentData().GetStudentKey() != msg.Creator {
						return &types.MsgPayTaxesResponse{
							Status: -1,
						}, types.ErrKeyEnteredMismatchStudent
					} else {

						err := utilfunc.CheckCompleteInformation(searchedStudent)

						if err != nil {
							return &types.MsgPayTaxesResponse{
								Status: -1,
							}, types.ErrIncompleteStudentInformation
						} else {

							err := k.Keeper.CollectAndPayTaxes(ctx, &searchedStudent, uniInfo.UniversityKey)
							if err != nil {
								return &types.MsgPayTaxesResponse{
									Status: -1,
								}, err
							}

							k.Keeper.SetStoredStudent(ctx, searchedStudent)
							return &types.MsgPayTaxesResponse{
								Status: 0,
							}, nil
						}
					}
				}
			}
		}
	}
}
