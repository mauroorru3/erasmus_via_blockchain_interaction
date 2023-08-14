package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	keepertest "university_chain_it/testutil/keeper"
	"university_chain_it/testutil/nullify"
	"university_chain_it/x/universitychainit/keeper"
	"university_chain_it/x/universitychainit/types"
)

func createTestContactInfo(keeper *keeper.Keeper, ctx sdk.Context) types.ContactInfo {
	item := types.ContactInfo{}
	keeper.SetContactInfo(ctx, item)
	return item
}

func TestContactInfoGet(t *testing.T) {
	keeper, ctx := keepertest.UniversitychainitKeeper(t)
	item := createTestContactInfo(keeper, ctx)
	rst, found := keeper.GetContactInfo(ctx)
	require.True(t, found)
	require.Equal(t,
		nullify.Fill(&item),
		nullify.Fill(&rst),
	)
}

func TestContactInfoRemove(t *testing.T) {
	keeper, ctx := keepertest.UniversitychainitKeeper(t)
	createTestContactInfo(keeper, ctx)
	keeper.RemoveContactInfo(ctx)
	_, found := keeper.GetContactInfo(ctx)
	require.False(t, found)
}
