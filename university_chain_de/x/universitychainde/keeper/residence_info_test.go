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

func createTestResidenceInfo(keeper *keeper.Keeper, ctx sdk.Context) types.ResidenceInfo {
	item := types.ResidenceInfo{}
	keeper.SetResidenceInfo(ctx, item)
	return item
}

func TestResidenceInfoGet(t *testing.T) {
	keeper, ctx := keepertest.UniversitychaindeKeeper(t)
	item := createTestResidenceInfo(keeper, ctx)
	rst, found := keeper.GetResidenceInfo(ctx)
	require.True(t, found)
	require.Equal(t,
		nullify.Fill(&item),
		nullify.Fill(&rst),
	)
}

func TestResidenceInfoRemove(t *testing.T) {
	keeper, ctx := keepertest.UniversitychaindeKeeper(t)
	createTestResidenceInfo(keeper, ctx)
	keeper.RemoveResidenceInfo(ctx)
	_, found := keeper.GetResidenceInfo(ctx)
	require.False(t, found)
}
