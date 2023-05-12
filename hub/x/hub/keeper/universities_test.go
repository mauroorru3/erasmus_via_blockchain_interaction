package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	keepertest "hub/testutil/keeper"
	"hub/testutil/nullify"
	"hub/x/hub/keeper"
	"hub/x/hub/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNUniversities(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Universities {
	items := make([]types.Universities, n)
	for i := range items {
		items[i].UniversityName = strconv.Itoa(i)

		keeper.SetUniversities(ctx, items[i])
	}
	return items
}

func TestUniversitiesGet(t *testing.T) {
	keeper, ctx := keepertest.HubKeeper(t)
	items := createNUniversities(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetUniversities(ctx,
			item.UniversityName,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestUniversitiesRemove(t *testing.T) {
	keeper, ctx := keepertest.HubKeeper(t)
	items := createNUniversities(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveUniversities(ctx,
			item.UniversityName,
		)
		_, found := keeper.GetUniversities(ctx,
			item.UniversityName,
		)
		require.False(t, found)
	}
}

func TestUniversitiesGetAll(t *testing.T) {
	keeper, ctx := keepertest.HubKeeper(t)
	items := createNUniversities(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllUniversities(ctx)),
	)
}
