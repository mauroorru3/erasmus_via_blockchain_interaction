package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"university_chain_it/x/universitychainit/types"
)

// SetForeignUniversities set a specific foreignUniversities in the store from its index
func (k Keeper) SetForeignUniversities(ctx sdk.Context, foreignUniversities types.ForeignUniversities) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ForeignUniversitiesKeyPrefix))
	b := k.cdc.MustMarshal(&foreignUniversities)
	store.Set(types.ForeignUniversitiesKey(
		foreignUniversities.UniversityName,
	), b)
}

// GetForeignUniversities returns a foreignUniversities from its index
func (k Keeper) GetForeignUniversities(
	ctx sdk.Context,
	universityName string,

) (val types.ForeignUniversities, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ForeignUniversitiesKeyPrefix))

	b := store.Get(types.ForeignUniversitiesKey(
		universityName,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveForeignUniversities removes a foreignUniversities from the store
func (k Keeper) RemoveForeignUniversities(
	ctx sdk.Context,
	universityName string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ForeignUniversitiesKeyPrefix))
	store.Delete(types.ForeignUniversitiesKey(
		universityName,
	))
}

// GetAllForeignUniversities returns all foreignUniversities
func (k Keeper) GetAllForeignUniversities(ctx sdk.Context) (list []types.ForeignUniversities) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ForeignUniversitiesKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.ForeignUniversities
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
