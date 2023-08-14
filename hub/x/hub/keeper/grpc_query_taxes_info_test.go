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

func TestTaxesInfoQuery(t *testing.T) {
	keeper, ctx := keepertest.HubKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	item := createTestTaxesInfo(keeper, ctx)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetTaxesInfoRequest
		response *types.QueryGetTaxesInfoResponse
		err      error
	}{
		{
			desc:     "First",
			request:  &types.QueryGetTaxesInfoRequest{},
			response: &types.QueryGetTaxesInfoResponse{TaxesInfo: item},
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.TaxesInfo(wctx, tc.request)
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
