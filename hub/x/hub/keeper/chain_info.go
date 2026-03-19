package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"hub/x/hub/types"
)

// SetChainInfo set chainInfo in the store
func (k Keeper) SetChainInfo(ctx sdk.Context, chainInfo types.ChainInfo) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ChainInfoKey))
	b := k.cdc.MustMarshal(&chainInfo)
	store.Set([]byte{0}, b)
}

// GetChainInfo returns chainInfo
func (k Keeper) GetChainInfo(ctx sdk.Context) (val types.ChainInfo, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ChainInfoKey))

	b := store.Get([]byte{0})
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveChainInfo removes chainInfo from the store
func (k Keeper) RemoveChainInfo(ctx sdk.Context) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ChainInfoKey))
	store.Delete([]byte{0})
}
