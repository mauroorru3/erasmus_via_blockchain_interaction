package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"hub/x/hub/types"
)

func (k Keeper) PersonalInfo(c context.Context, req *types.QueryGetPersonalInfoRequest) (*types.QueryGetPersonalInfoResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetPersonalInfo(ctx)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetPersonalInfoResponse{PersonalInfo: val}, nil
}
