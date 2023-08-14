package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"hub/x/hub/types"
)

// SetTaxesInfo set taxesInfo in the store
func (k Keeper) SetTaxesInfo(ctx sdk.Context, taxesInfo types.TaxesInfo) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TaxesInfoKey))
	b := k.cdc.MustMarshal(&taxesInfo)
	store.Set([]byte{0}, b)
}

// GetTaxesInfo returns taxesInfo
func (k Keeper) GetTaxesInfo(ctx sdk.Context) (val types.TaxesInfo, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TaxesInfoKey))

	b := store.Get([]byte{0})
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveTaxesInfo removes taxesInfo from the store
func (k Keeper) RemoveTaxesInfo(ctx sdk.Context) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TaxesInfoKey))
	store.Delete([]byte{0})
}
