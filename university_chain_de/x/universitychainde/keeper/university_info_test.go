package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	keepertest "university_chain_de/testutil/keeper"
	"university_chain_de/testutil/nullify"
	"university_chain_de/x/universitychainde/keeper"
	"university_chain_de/x/universitychainde/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNUniversityInfo(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.UniversityInfo {
	items := make([]types.UniversityInfo, n)
	for i := range items {
		items[i].UniversityName = strconv.Itoa(i)

		keeper.SetUniversityInfo(ctx, items[i])
	}
	return items
}

func TestUniversityInfoGet(t *testing.T) {
	keeper, ctx := keepertest.UniversitychaindeKeeper(t)
	items := createNUniversityInfo(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetUniversityInfo(ctx,
			item.UniversityName,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestUniversityInfoRemove(t *testing.T) {
	keeper, ctx := keepertest.UniversitychaindeKeeper(t)
	items := createNUniversityInfo(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveUniversityInfo(ctx,
			item.UniversityName,
		)
		_, found := keeper.GetUniversityInfo(ctx,
			item.UniversityName,
		)
		require.False(t, found)
	}
}

func TestUniversityInfoGetAll(t *testing.T) {
	keeper, ctx := keepertest.UniversitychaindeKeeper(t)
	items := createNUniversityInfo(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllUniversityInfo(ctx)),
	)
}
