package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"hub/x/hub/types"
)

// SetStudentInfo set studentInfo in the store
func (k Keeper) SetStudentInfo(ctx sdk.Context, studentInfo types.StudentInfo) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.StudentInfoKey))
	b := k.cdc.MustMarshal(&studentInfo)
	store.Set([]byte{0}, b)
}

// GetStudentInfo returns studentInfo
func (k Keeper) GetStudentInfo(ctx sdk.Context) (val types.StudentInfo, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.StudentInfoKey))

	b := store.Get([]byte{0})
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveStudentInfo removes studentInfo from the store
func (k Keeper) RemoveStudentInfo(ctx sdk.Context) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.StudentInfoKey))
	store.Delete([]byte{0})
}
