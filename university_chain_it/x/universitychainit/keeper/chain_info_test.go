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

func createTestChainInfo(keeper *keeper.Keeper, ctx sdk.Context) types.ChainInfo {
	item := types.ChainInfo{}
	keeper.SetChainInfo(ctx, item)
	return item
}

func TestChainInfoGet(t *testing.T) {
	keeper, ctx := keepertest.UniversitychainitKeeper(t)
	item := createTestChainInfo(keeper, ctx)
	rst, found := keeper.GetChainInfo(ctx)
	require.True(t, found)
	require.Equal(t,
		nullify.Fill(&item),
		nullify.Fill(&rst),
	)
}

func TestChainInfoRemove(t *testing.T) {
	keeper, ctx := keepertest.UniversitychainitKeeper(t)
	createTestChainInfo(keeper, ctx)
	keeper.RemoveChainInfo(ctx)
	_, found := keeper.GetChainInfo(ctx)
	require.False(t, found)
}
