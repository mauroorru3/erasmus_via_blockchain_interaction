package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"hub/x/hub/types"
)

// SetPersonalInfo set personalInfo in the store
func (k Keeper) SetPersonalInfo(ctx sdk.Context, personalInfo types.PersonalInfo) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PersonalInfoKey))
	b := k.cdc.MustMarshal(&personalInfo)
	store.Set([]byte{0}, b)
}

// GetPersonalInfo returns personalInfo
func (k Keeper) GetPersonalInfo(ctx sdk.Context) (val types.PersonalInfo, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PersonalInfoKey))

	b := store.Get([]byte{0})
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemovePersonalInfo removes personalInfo from the store
func (k Keeper) RemovePersonalInfo(ctx sdk.Context) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PersonalInfoKey))
	store.Delete([]byte{0})
}
