package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "university_chain_de/testutil/keeper"
	"university_chain_de/x/universitychainde/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.UniversitychaindeKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
