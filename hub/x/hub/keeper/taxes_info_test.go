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

func createTestTaxesInfo(keeper *keeper.Keeper, ctx sdk.Context) types.TaxesInfo {
	item := types.TaxesInfo{}
	keeper.SetTaxesInfo(ctx, item)
	return item
}

func TestTaxesInfoGet(t *testing.T) {
	keeper, ctx := keepertest.HubKeeper(t)
	item := createTestTaxesInfo(keeper, ctx)
	rst, found := keeper.GetTaxesInfo(ctx)
	require.True(t, found)
	require.Equal(t,
		nullify.Fill(&item),
		nullify.Fill(&rst),
	)
}

func TestTaxesInfoRemove(t *testing.T) {
	keeper, ctx := keepertest.HubKeeper(t)
	createTestTaxesInfo(keeper, ctx)
	keeper.RemoveTaxesInfo(ctx)
	_, found := keeper.GetTaxesInfo(ctx)
	require.False(t, found)
}
