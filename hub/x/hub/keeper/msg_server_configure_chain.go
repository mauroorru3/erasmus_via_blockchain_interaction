package keeper

import (
	"context"
	"fmt"
	"os"

	"hub/x/hub/types"
	"hub/x/hub/utilfunc"

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
				UniversitiesList, err := utilfunc.ReadForeignUniversityInfo()
				if err != nil {
					return &types.MsgConfigureChainResponse{
						Status: -1,
					}, err
				} else {
					i := 0
					for i = 0; i < len(UniversitiesList); i++ {

						foreignUniversity := types.Universities{
							UniversityName:      UniversitiesList[i].Name,
							UniversitiesCountry: UniversitiesList[i].Country,
							UniversitiesKey:     UniversitiesList[i].Address,
						}
						err = foreignUniversity.Validate()
						if err != nil {
							return &types.MsgConfigureChainResponse{
								Status: -1,
							}, err
						} else {
							k.Keeper.SetUniversities(ctx, foreignUniversity)
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
