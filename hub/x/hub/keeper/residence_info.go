package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"hub/x/hub/types"
)

// SetResidenceInfo set residenceInfo in the store
func (k Keeper) SetResidenceInfo(ctx sdk.Context, residenceInfo types.ResidenceInfo) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ResidenceInfoKey))
	b := k.cdc.MustMarshal(&residenceInfo)
	store.Set([]byte{0}, b)
}

// GetResidenceInfo returns residenceInfo
func (k Keeper) GetResidenceInfo(ctx sdk.Context) (val types.ResidenceInfo, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ResidenceInfoKey))

	b := store.Get([]byte{0})
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveResidenceInfo removes residenceInfo from the store
func (k Keeper) RemoveResidenceInfo(ctx sdk.Context) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ResidenceInfoKey))
	store.Delete([]byte{0})
}
