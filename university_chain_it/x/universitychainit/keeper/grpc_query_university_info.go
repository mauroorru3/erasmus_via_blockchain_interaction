package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"university_chain_it/x/universitychainit/types"
)

func (k Keeper) UniversityInfoAll(c context.Context, req *types.QueryAllUniversityInfoRequest) (*types.QueryAllUniversityInfoResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var universityInfos []types.UniversityInfo
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	universityInfoStore := prefix.NewStore(store, types.KeyPrefix(types.UniversityInfoKeyPrefix))

	pageRes, err := query.Paginate(universityInfoStore, req.Pagination, func(key []byte, value []byte) error {
		var universityInfo types.UniversityInfo
		if err := k.cdc.Unmarshal(value, &universityInfo); err != nil {
			return err
		}

		universityInfos = append(universityInfos, universityInfo)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllUniversityInfoResponse{UniversityInfo: universityInfos, Pagination: pageRes}, nil
}

func (k Keeper) UniversityInfo(c context.Context, req *types.QueryGetUniversityInfoRequest) (*types.QueryGetUniversityInfoResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetUniversityInfo(
		ctx,
		req.UniversityName,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetUniversityInfoResponse{UniversityInfo: val}, nil
}
