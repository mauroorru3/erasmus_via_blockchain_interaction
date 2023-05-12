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

func TestContactInfoQuery(t *testing.T) {
	keeper, ctx := keepertest.HubKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	item := createTestContactInfo(keeper, ctx)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetContactInfoRequest
		response *types.QueryGetContactInfoResponse
		err      error
	}{
		{
			desc:     "First",
			request:  &types.QueryGetContactInfoRequest{},
			response: &types.QueryGetContactInfoResponse{ContactInfo: item},
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.ContactInfo(wctx, tc.request)
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
