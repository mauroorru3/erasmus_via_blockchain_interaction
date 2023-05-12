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

func (k Keeper) ForeignUniversitiesAll(c context.Context, req *types.QueryAllForeignUniversitiesRequest) (*types.QueryAllForeignUniversitiesResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var foreignUniversitiess []types.ForeignUniversities
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	foreignUniversitiesStore := prefix.NewStore(store, types.KeyPrefix(types.ForeignUniversitiesKeyPrefix))

	pageRes, err := query.Paginate(foreignUniversitiesStore, req.Pagination, func(key []byte, value []byte) error {
		var foreignUniversities types.ForeignUniversities
		if err := k.cdc.Unmarshal(value, &foreignUniversities); err != nil {
			return err
		}

		foreignUniversitiess = append(foreignUniversitiess, foreignUniversities)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllForeignUniversitiesResponse{ForeignUniversities: foreignUniversitiess, Pagination: pageRes}, nil
}

func (k Keeper) ForeignUniversities(c context.Context, req *types.QueryGetForeignUniversitiesRequest) (*types.QueryGetForeignUniversitiesResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetForeignUniversities(
		ctx,
		req.UniversityName,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetForeignUniversitiesResponse{ForeignUniversities: val}, nil
}
