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

func TestStudentInfoQuery(t *testing.T) {
	keeper, ctx := keepertest.UniversitychaindeKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	item := createTestStudentInfo(keeper, ctx)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetStudentInfoRequest
		response *types.QueryGetStudentInfoResponse
		err      error
	}{
		{
			desc:     "First",
			request:  &types.QueryGetStudentInfoRequest{},
			response: &types.QueryGetStudentInfoResponse{StudentInfo: item},
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.StudentInfo(wctx, tc.request)
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
