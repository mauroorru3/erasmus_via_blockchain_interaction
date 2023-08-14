package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	keepertest "hub/testutil/keeper"
	"hub/testutil/nullify"
	"hub/x/hub/keeper"
	"hub/x/hub/types"
)

func createTestPersonalInfo(keeper *keeper.Keeper, ctx sdk.Context) types.PersonalInfo {
	item := types.PersonalInfo{}
	keeper.SetPersonalInfo(ctx, item)
	return item
}

func TestPersonalInfoGet(t *testing.T) {
	keeper, ctx := keepertest.HubKeeper(t)
	item := createTestPersonalInfo(keeper, ctx)
	rst, found := keeper.GetPersonalInfo(ctx)
	require.True(t, found)
	require.Equal(t,
		nullify.Fill(&item),
		nullify.Fill(&rst),
	)
}

func TestPersonalInfoRemove(t *testing.T) {
	keeper, ctx := keepertest.HubKeeper(t)
	createTestPersonalInfo(keeper, ctx)
	keeper.RemovePersonalInfo(ctx)
	_, found := keeper.GetPersonalInfo(ctx)
	require.False(t, found)
}
