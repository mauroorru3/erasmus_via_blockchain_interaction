package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"hub/x/hub/keeper"
	"hub/x/hub/types"
)

func SimulateMsgConfigureChain(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgConfigureChain{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the ConfigureChain simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "ConfigureChain simulation not implemented"), nil, nil
	}
}
