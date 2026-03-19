package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"hub/x/hub/types"
)

// SetContactInfo set contactInfo in the store
func (k Keeper) SetContactInfo(ctx sdk.Context, contactInfo types.ContactInfo) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ContactInfoKey))
	b := k.cdc.MustMarshal(&contactInfo)
	store.Set([]byte{0}, b)
}

// GetContactInfo returns contactInfo
func (k Keeper) GetContactInfo(ctx sdk.Context) (val types.ContactInfo, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ContactInfoKey))

	b := store.Get([]byte{0})
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveContactInfo removes contactInfo from the store
func (k Keeper) RemoveContactInfo(ctx sdk.Context) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ContactInfoKey))
	store.Delete([]byte{0})
}
