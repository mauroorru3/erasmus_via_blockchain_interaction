package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "university_chain_it/testutil/keeper"
	"university_chain_it/x/universitychainit/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.UniversitychainitKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
