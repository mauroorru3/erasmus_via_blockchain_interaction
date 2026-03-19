package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"hub/x/hub/types"
)

// SetUniversities set a specific universities in the store from its index
func (k Keeper) SetUniversities(ctx sdk.Context, universities types.Universities) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.UniversitiesKeyPrefix))
	b := k.cdc.MustMarshal(&universities)
	store.Set(types.UniversitiesKey(
		universities.UniversityName,
	), b)
}

// GetUniversities returns a universities from its index
func (k Keeper) GetUniversities(
	ctx sdk.Context,
	universityName string,

) (val types.Universities, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.UniversitiesKeyPrefix))

	b := store.Get(types.UniversitiesKey(
		universityName,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveUniversities removes a universities from the store
func (k Keeper) RemoveUniversities(
	ctx sdk.Context,
	universityName string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.UniversitiesKeyPrefix))
	store.Delete(types.UniversitiesKey(
		universityName,
	))
}

// GetAllUniversities returns all universities
func (k Keeper) GetAllUniversities(ctx sdk.Context) (list []types.Universities) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.UniversitiesKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Universities
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
