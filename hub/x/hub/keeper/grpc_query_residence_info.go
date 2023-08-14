package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"hub/x/hub/types"
)

func (k Keeper) ResidenceInfo(c context.Context, req *types.QueryGetResidenceInfoRequest) (*types.QueryGetResidenceInfoResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetResidenceInfo(ctx)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetResidenceInfoResponse{ResidenceInfo: val}, nil
}
