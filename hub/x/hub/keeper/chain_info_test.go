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

func createTestChainInfo(keeper *keeper.Keeper, ctx sdk.Context) types.ChainInfo {
	item := types.ChainInfo{}
	keeper.SetChainInfo(ctx, item)
	return item
}

func TestChainInfoGet(t *testing.T) {
	keeper, ctx := keepertest.HubKeeper(t)
	item := createTestChainInfo(keeper, ctx)
	rst, found := keeper.GetChainInfo(ctx)
	require.True(t, found)
	require.Equal(t,
		nullify.Fill(&item),
		nullify.Fill(&rst),
	)
}

func TestChainInfoRemove(t *testing.T) {
	keeper, ctx := keepertest.HubKeeper(t)
	createTestChainInfo(keeper, ctx)
	keeper.RemoveChainInfo(ctx)
	_, found := keeper.GetChainInfo(ctx)
	require.False(t, found)
}
