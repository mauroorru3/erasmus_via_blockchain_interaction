package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"university_chain_it/x/universitychainit/types"
)

// SetUniversityInfo set a specific universityInfo in the store from its index
func (k Keeper) SetUniversityInfo(ctx sdk.Context, universityInfo types.UniversityInfo) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.UniversityInfoKeyPrefix))
	b := k.cdc.MustMarshal(&universityInfo)
	store.Set(types.UniversityInfoKey(
		universityInfo.UniversityName,
	), b)
}

// GetUniversityInfo returns a universityInfo from its index
func (k Keeper) GetUniversityInfo(
	ctx sdk.Context,
	universityName string,

) (val types.UniversityInfo, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.UniversityInfoKeyPrefix))

	b := store.Get(types.UniversityInfoKey(
		universityName,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveUniversityInfo removes a universityInfo from the store
func (k Keeper) RemoveUniversityInfo(
	ctx sdk.Context,
	universityName string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.UniversityInfoKeyPrefix))
	store.Delete(types.UniversityInfoKey(
		universityName,
	))
}

// GetAllUniversityInfo returns all universityInfo
func (k Keeper) GetAllUniversityInfo(ctx sdk.Context) (list []types.UniversityInfo) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.UniversityInfoKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.UniversityInfo
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
