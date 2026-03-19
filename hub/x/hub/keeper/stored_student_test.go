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

func createNStoredStudent(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.StoredStudent {
	items := make([]types.StoredStudent, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetStoredStudent(ctx, items[i])
	}
	return items
}

func TestStoredStudentGet(t *testing.T) {
	keeper, ctx := keepertest.HubKeeper(t)
	items := createNStoredStudent(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetStoredStudent(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestStoredStudentRemove(t *testing.T) {
	keeper, ctx := keepertest.HubKeeper(t)
	items := createNStoredStudent(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveStoredStudent(ctx,
			item.Index,
		)
		_, found := keeper.GetStoredStudent(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestStoredStudentGetAll(t *testing.T) {
	keeper, ctx := keepertest.HubKeeper(t)
	items := createNStoredStudent(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllStoredStudent(ctx)),
	)
}
