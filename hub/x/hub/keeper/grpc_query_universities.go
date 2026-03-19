package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"hub/x/hub/types"
)

func (k Keeper) UniversitiesAll(c context.Context, req *types.QueryAllUniversitiesRequest) (*types.QueryAllUniversitiesResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var universitiess []types.Universities
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	universitiesStore := prefix.NewStore(store, types.KeyPrefix(types.UniversitiesKeyPrefix))

	pageRes, err := query.Paginate(universitiesStore, req.Pagination, func(key []byte, value []byte) error {
		var universities types.Universities
		if err := k.cdc.Unmarshal(value, &universities); err != nil {
			return err
		}

		universitiess = append(universitiess, universities)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllUniversitiesResponse{Universities: universitiess, Pagination: pageRes}, nil
}

func (k Keeper) Universities(c context.Context, req *types.QueryGetUniversitiesRequest) (*types.QueryGetUniversitiesResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetUniversities(
		ctx,
		req.UniversityName,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetUniversitiesResponse{Universities: val}, nil
}
