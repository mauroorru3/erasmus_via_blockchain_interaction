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

func createTestStudentInfo(keeper *keeper.Keeper, ctx sdk.Context) types.StudentInfo {
	item := types.StudentInfo{}
	keeper.SetStudentInfo(ctx, item)
	return item
}

func TestStudentInfoGet(t *testing.T) {
	keeper, ctx := keepertest.HubKeeper(t)
	item := createTestStudentInfo(keeper, ctx)
	rst, found := keeper.GetStudentInfo(ctx)
	require.True(t, found)
	require.Equal(t,
		nullify.Fill(&item),
		nullify.Fill(&rst),
	)
}

func TestStudentInfoRemove(t *testing.T) {
	keeper, ctx := keepertest.HubKeeper(t)
	createTestStudentInfo(keeper, ctx)
	keeper.RemoveStudentInfo(ctx)
	_, found := keeper.GetStudentInfo(ctx)
	require.False(t, found)
}
