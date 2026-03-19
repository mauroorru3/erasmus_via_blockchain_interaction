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

func createNProfessorsExams(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.ProfessorsExams {
	items := make([]types.ProfessorsExams, n)
	for i := range items {
		items[i].ExamName = strconv.Itoa(i)

		keeper.SetProfessorsExams(ctx, items[i])
	}
	return items
}

func TestProfessorsExamsGet(t *testing.T) {
	keeper, ctx := keepertest.UniversitychainitKeeper(t)
	items := createNProfessorsExams(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetProfessorsExams(ctx,
			item.ExamName,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestProfessorsExamsRemove(t *testing.T) {
	keeper, ctx := keepertest.UniversitychainitKeeper(t)
	items := createNProfessorsExams(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveProfessorsExams(ctx,
			item.ExamName,
		)
		_, found := keeper.GetProfessorsExams(ctx,
			item.ExamName,
		)
		require.False(t, found)
	}
}

func TestProfessorsExamsGetAll(t *testing.T) {
	keeper, ctx := keepertest.UniversitychainitKeeper(t)
	items := createNProfessorsExams(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllProfessorsExams(ctx)),
	)
}
