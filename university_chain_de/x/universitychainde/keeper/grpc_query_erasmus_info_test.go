package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "university_chain_de/testutil/keeper"
	"university_chain_de/testutil/nullify"
	"university_chain_de/x/universitychainde/types"
)

func TestErasmusInfoQuery(t *testing.T) {
	keeper, ctx := keepertest.UniversitychaindeKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	item := createTestErasmusInfo(keeper, ctx)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetErasmusInfoRequest
		response *types.QueryGetErasmusInfoResponse
		err      error
	}{
		{
			desc:     "First",
			request:  &types.QueryGetErasmusInfoRequest{},
			response: &types.QueryGetErasmusInfoResponse{ErasmusInfo: item},
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.ErasmusInfo(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}
}
