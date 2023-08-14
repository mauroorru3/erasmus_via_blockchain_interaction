package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	keepertest "university_chain_it/testutil/keeper"
	"university_chain_it/testutil/nullify"
	"university_chain_it/x/universitychainit/keeper"
	"university_chain_it/x/universitychainit/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNForeignUniversities(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.ForeignUniversities {
	items := make([]types.ForeignUniversities, n)
	for i := range items {
		items[i].UniversityName = strconv.Itoa(i)

		keeper.SetForeignUniversities(ctx, items[i])
	}
	return items
}

func TestForeignUniversitiesGet(t *testing.T) {
	keeper, ctx := keepertest.UniversitychainitKeeper(t)
	items := createNForeignUniversities(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetForeignUniversities(ctx,
			item.UniversityName,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestForeignUniversitiesRemove(t *testing.T) {
	keeper, ctx := keepertest.UniversitychainitKeeper(t)
	items := createNForeignUniversities(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveForeignUniversities(ctx,
			item.UniversityName,
		)
		_, found := keeper.GetForeignUniversities(ctx,
			item.UniversityName,
		)
		require.False(t, found)
	}
}

func TestForeignUniversitiesGetAll(t *testing.T) {
	keeper, ctx := keepertest.UniversitychainitKeeper(t)
	items := createNForeignUniversities(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllForeignUniversities(ctx)),
	)
}
