package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	keepertest "university_chain_it/testutil/keeper"
	"university_chain_it/testutil/nullify"
	"university_chain_it/x/universitychainit/keeper"
	"university_chain_it/x/universitychainit/types"
)

func createTestTaxesInfo(keeper *keeper.Keeper, ctx sdk.Context) types.TaxesInfo {
	item := types.TaxesInfo{}
	keeper.SetTaxesInfo(ctx, item)
	return item
}

func TestTaxesInfoGet(t *testing.T) {
	keeper, ctx := keepertest.UniversitychainitKeeper(t)
	item := createTestTaxesInfo(keeper, ctx)
	rst, found := keeper.GetTaxesInfo(ctx)
	require.True(t, found)
	require.Equal(t,
		nullify.Fill(&item),
		nullify.Fill(&rst),
	)
}

func TestTaxesInfoRemove(t *testing.T) {
	keeper, ctx := keepertest.UniversitychainitKeeper(t)
	createTestTaxesInfo(keeper, ctx)
	keeper.RemoveTaxesInfo(ctx)
	_, found := keeper.GetTaxesInfo(ctx)
	require.False(t, found)
}
