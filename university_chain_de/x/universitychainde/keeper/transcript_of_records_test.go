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

func createTestTranscriptOfRecords(keeper *keeper.Keeper, ctx sdk.Context) types.TranscriptOfRecords {
	item := types.TranscriptOfRecords{}
	keeper.SetTranscriptOfRecords(ctx, item)
	return item
}

func TestTranscriptOfRecordsGet(t *testing.T) {
	keeper, ctx := keepertest.UniversitychaindeKeeper(t)
	item := createTestTranscriptOfRecords(keeper, ctx)
	rst, found := keeper.GetTranscriptOfRecords(ctx)
	require.True(t, found)
	require.Equal(t,
		nullify.Fill(&item),
		nullify.Fill(&rst),
	)
}

func TestTranscriptOfRecordsRemove(t *testing.T) {
	keeper, ctx := keepertest.UniversitychaindeKeeper(t)
	createTestTranscriptOfRecords(keeper, ctx)
	keeper.RemoveTranscriptOfRecords(ctx)
	_, found := keeper.GetTranscriptOfRecords(ctx)
	require.False(t, found)
}
