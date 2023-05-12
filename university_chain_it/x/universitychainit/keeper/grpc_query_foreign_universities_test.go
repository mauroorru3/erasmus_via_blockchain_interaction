package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "university_chain_it/testutil/keeper"
	"university_chain_it/testutil/nullify"
	"university_chain_it/x/universitychainit/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestForeignUniversitiesQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.UniversitychainitKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNForeignUniversities(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetForeignUniversitiesRequest
		response *types.QueryGetForeignUniversitiesResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetForeignUniversitiesRequest{
				UniversityName: msgs[0].UniversityName,
			},
			response: &types.QueryGetForeignUniversitiesResponse{ForeignUniversities: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetForeignUniversitiesRequest{
				UniversityName: msgs[1].UniversityName,
			},
			response: &types.QueryGetForeignUniversitiesResponse{ForeignUniversities: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetForeignUniversitiesRequest{
				UniversityName: strconv.Itoa(100000),
			},
			err: status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.ForeignUniversities(wctx, tc.request)
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

func TestForeignUniversitiesQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.UniversitychainitKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNForeignUniversities(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllForeignUniversitiesRequest {
		return &types.QueryAllForeignUniversitiesRequest{
			Pagination: &query.PageRequest{
				Key:        next,
				Offset:     offset,
				Limit:      limit,
				CountTotal: total,
			},
		}
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.ForeignUniversitiesAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.ForeignUniversities), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.ForeignUniversities),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.ForeignUniversitiesAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.ForeignUniversities), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.ForeignUniversities),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.ForeignUniversitiesAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.ForeignUniversities),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.ForeignUniversitiesAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
