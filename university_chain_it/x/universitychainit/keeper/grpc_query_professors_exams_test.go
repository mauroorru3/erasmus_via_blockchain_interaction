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

func TestProfessorsExamsQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.UniversitychainitKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNProfessorsExams(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetProfessorsExamsRequest
		response *types.QueryGetProfessorsExamsResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetProfessorsExamsRequest{
				ExamName: msgs[0].ExamName,
			},
			response: &types.QueryGetProfessorsExamsResponse{ProfessorsExams: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetProfessorsExamsRequest{
				ExamName: msgs[1].ExamName,
			},
			response: &types.QueryGetProfessorsExamsResponse{ProfessorsExams: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetProfessorsExamsRequest{
				ExamName: strconv.Itoa(100000),
			},
			err: status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.ProfessorsExams(wctx, tc.request)
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

func TestProfessorsExamsQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.UniversitychainitKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNProfessorsExams(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllProfessorsExamsRequest {
		return &types.QueryAllProfessorsExamsRequest{
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
			resp, err := keeper.ProfessorsExamsAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.ProfessorsExams), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.ProfessorsExams),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.ProfessorsExamsAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.ProfessorsExams), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.ProfessorsExams),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.ProfessorsExamsAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.ProfessorsExams),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.ProfessorsExamsAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
