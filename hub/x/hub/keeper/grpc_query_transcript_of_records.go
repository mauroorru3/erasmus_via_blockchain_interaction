package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"hub/x/hub/types"
)

func (k Keeper) TranscriptOfRecords(c context.Context, req *types.QueryGetTranscriptOfRecordsRequest) (*types.QueryGetTranscriptOfRecordsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetTranscriptOfRecords(ctx)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetTranscriptOfRecordsResponse{TranscriptOfRecords: val}, nil
}
