package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	keepertest "university_chain_de/testutil/keeper"
	"university_chain_de/testutil/nullify"
	"university_chain_de/x/universitychainde/keeper"
	"university_chain_de/x/universitychainde/types"
)

func createTestPersonalInfo(keeper *keeper.Keeper, ctx sdk.Context) types.PersonalInfo {
	item := types.PersonalInfo{}
	keeper.SetPersonalInfo(ctx, item)
	return item
}

func TestPersonalInfoGet(t *testing.T) {
	keeper, ctx := keepertest.UniversitychaindeKeeper(t)
	item := createTestPersonalInfo(keeper, ctx)
	rst, found := keeper.GetPersonalInfo(ctx)
	require.True(t, found)
	require.Equal(t,
		nullify.Fill(&item),
		nullify.Fill(&rst),
	)
}

func TestPersonalInfoRemove(t *testing.T) {
	keeper, ctx := keepertest.UniversitychaindeKeeper(t)
	createTestPersonalInfo(keeper, ctx)
	keeper.RemovePersonalInfo(ctx)
	_, found := keeper.GetPersonalInfo(ctx)
	require.False(t, found)
}
