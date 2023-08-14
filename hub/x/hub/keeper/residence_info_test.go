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

func createTestResidenceInfo(keeper *keeper.Keeper, ctx sdk.Context) types.ResidenceInfo {
	item := types.ResidenceInfo{}
	keeper.SetResidenceInfo(ctx, item)
	return item
}

func TestResidenceInfoGet(t *testing.T) {
	keeper, ctx := keepertest.HubKeeper(t)
	item := createTestResidenceInfo(keeper, ctx)
	rst, found := keeper.GetResidenceInfo(ctx)
	require.True(t, found)
	require.Equal(t,
		nullify.Fill(&item),
		nullify.Fill(&rst),
	)
}

func TestResidenceInfoRemove(t *testing.T) {
	keeper, ctx := keepertest.HubKeeper(t)
	createTestResidenceInfo(keeper, ctx)
	keeper.RemoveResidenceInfo(ctx)
	_, found := keeper.GetResidenceInfo(ctx)
	require.False(t, found)
}
