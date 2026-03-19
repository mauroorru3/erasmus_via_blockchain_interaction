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

func (k Keeper) ProfessorsExamsAll(c context.Context, req *types.QueryAllProfessorsExamsRequest) (*types.QueryAllProfessorsExamsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var professorsExamss []types.ProfessorsExams
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	professorsExamsStore := prefix.NewStore(store, types.KeyPrefix(types.ProfessorsExamsKeyPrefix))

	pageRes, err := query.Paginate(professorsExamsStore, req.Pagination, func(key []byte, value []byte) error {
		var professorsExams types.ProfessorsExams
		if err := k.cdc.Unmarshal(value, &professorsExams); err != nil {
			return err
		}

		professorsExamss = append(professorsExamss, professorsExams)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllProfessorsExamsResponse{ProfessorsExams: professorsExamss, Pagination: pageRes}, nil
}

func (k Keeper) ProfessorsExams(c context.Context, req *types.QueryGetProfessorsExamsRequest) (*types.QueryGetProfessorsExamsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetProfessorsExams(
		ctx,
		req.ExamName,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetProfessorsExamsResponse{ProfessorsExams: val}, nil
}
