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

func createTestTaxesInfo(keeper *keeper.Keeper, ctx sdk.Context) types.TaxesInfo {
	item := types.TaxesInfo{}
	keeper.SetTaxesInfo(ctx, item)
	return item
}

func TestTaxesInfoGet(t *testing.T) {
	keeper, ctx := keepertest.UniversitychaindeKeeper(t)
	item := createTestTaxesInfo(keeper, ctx)
	rst, found := keeper.GetTaxesInfo(ctx)
	require.True(t, found)
	require.Equal(t,
		nullify.Fill(&item),
		nullify.Fill(&rst),
	)
}

func TestTaxesInfoRemove(t *testing.T) {
	keeper, ctx := keepertest.UniversitychaindeKeeper(t)
	createTestTaxesInfo(keeper, ctx)
	keeper.RemoveTaxesInfo(ctx)
	_, found := keeper.GetTaxesInfo(ctx)
	require.False(t, found)
}
