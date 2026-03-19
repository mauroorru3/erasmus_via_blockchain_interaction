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

func createTestErasmusInfo(keeper *keeper.Keeper, ctx sdk.Context) types.ErasmusInfo {
	item := types.ErasmusInfo{}
	keeper.SetErasmusInfo(ctx, item)
	return item
}

func TestErasmusInfoGet(t *testing.T) {
	keeper, ctx := keepertest.HubKeeper(t)
	item := createTestErasmusInfo(keeper, ctx)
	rst, found := keeper.GetErasmusInfo(ctx)
	require.True(t, found)
	require.Equal(t,
		nullify.Fill(&item),
		nullify.Fill(&rst),
	)
}

func TestErasmusInfoRemove(t *testing.T) {
	keeper, ctx := keepertest.HubKeeper(t)
	createTestErasmusInfo(keeper, ctx)
	keeper.RemoveErasmusInfo(ctx)
	_, found := keeper.GetErasmusInfo(ctx)
	require.False(t, found)
}
