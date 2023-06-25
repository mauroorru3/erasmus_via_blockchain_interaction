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

func createTestStudentInfo(keeper *keeper.Keeper, ctx sdk.Context) types.StudentInfo {
	item := types.StudentInfo{}
	keeper.SetStudentInfo(ctx, item)
	return item
}

func TestStudentInfoGet(t *testing.T) {
	keeper, ctx := keepertest.UniversitychaindeKeeper(t)
	item := createTestStudentInfo(keeper, ctx)
	rst, found := keeper.GetStudentInfo(ctx)
	require.True(t, found)
	require.Equal(t,
		nullify.Fill(&item),
		nullify.Fill(&rst),
	)
}

func TestStudentInfoRemove(t *testing.T) {
	keeper, ctx := keepertest.UniversitychaindeKeeper(t)
	createTestStudentInfo(keeper, ctx)
	keeper.RemoveStudentInfo(ctx)
	_, found := keeper.GetStudentInfo(ctx)
	require.False(t, found)
}
