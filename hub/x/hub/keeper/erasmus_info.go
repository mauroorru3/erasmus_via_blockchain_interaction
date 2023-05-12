package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"hub/x/hub/types"
)

// SetErasmusInfo set erasmusInfo in the store
func (k Keeper) SetErasmusInfo(ctx sdk.Context, erasmusInfo types.ErasmusInfo) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ErasmusInfoKey))
	b := k.cdc.MustMarshal(&erasmusInfo)
	store.Set([]byte{0}, b)
}

// GetErasmusInfo returns erasmusInfo
func (k Keeper) GetErasmusInfo(ctx sdk.Context) (val types.ErasmusInfo, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ErasmusInfoKey))

	b := store.Get([]byte{0})
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveErasmusInfo removes erasmusInfo from the store
func (k Keeper) RemoveErasmusInfo(ctx sdk.Context) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ErasmusInfoKey))
	store.Delete([]byte{0})
}
