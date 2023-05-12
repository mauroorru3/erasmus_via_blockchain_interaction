package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "hub/testutil/keeper"
	"hub/testutil/nullify"
	"hub/x/hub/types"
)

func TestPersonalInfoQuery(t *testing.T) {
	keeper, ctx := keepertest.HubKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	item := createTestPersonalInfo(keeper, ctx)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetPersonalInfoRequest
		response *types.QueryGetPersonalInfoResponse
		err      error
	}{
		{
			desc:     "First",
			request:  &types.QueryGetPersonalInfoRequest{},
			response: &types.QueryGetPersonalInfoResponse{PersonalInfo: item},
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.PersonalInfo(wctx, tc.request)
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
