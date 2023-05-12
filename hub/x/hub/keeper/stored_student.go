package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"hub/x/hub/types"
)

// SetStoredStudent set a specific storedStudent in the store from its index
func (k Keeper) SetStoredStudent(ctx sdk.Context, storedStudent types.StoredStudent) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.StoredStudentKeyPrefix))
	b := k.cdc.MustMarshal(&storedStudent)
	store.Set(types.StoredStudentKey(
		storedStudent.Index,
	), b)
}

// GetStoredStudent returns a storedStudent from its index
func (k Keeper) GetStoredStudent(
	ctx sdk.Context,
	index string,

) (val types.StoredStudent, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.StoredStudentKeyPrefix))

	b := store.Get(types.StoredStudentKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveStoredStudent removes a storedStudent from the store
func (k Keeper) RemoveStoredStudent(
	ctx sdk.Context,
	index string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.StoredStudentKeyPrefix))
	store.Delete(types.StoredStudentKey(
		index,
	))
}

// GetAllStoredStudent returns all storedStudent
func (k Keeper) GetAllStoredStudent(ctx sdk.Context) (list []types.StoredStudent) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.StoredStudentKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.StoredStudent
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
