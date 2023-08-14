package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"hub/x/hub/types"
)

func (k Keeper) TaxesInfo(c context.Context, req *types.QueryGetTaxesInfoRequest) (*types.QueryGetTaxesInfoResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetTaxesInfo(ctx)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetTaxesInfoResponse{TaxesInfo: val}, nil
}
