package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"hub/x/hub/types"
)

// SetTranscriptOfRecords set transcriptOfRecords in the store
func (k Keeper) SetTranscriptOfRecords(ctx sdk.Context, transcriptOfRecords types.TranscriptOfRecords) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TranscriptOfRecordsKey))
	b := k.cdc.MustMarshal(&transcriptOfRecords)
	store.Set([]byte{0}, b)
}

// GetTranscriptOfRecords returns transcriptOfRecords
func (k Keeper) GetTranscriptOfRecords(ctx sdk.Context) (val types.TranscriptOfRecords, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TranscriptOfRecordsKey))

	b := store.Get([]byte{0})
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveTranscriptOfRecords removes transcriptOfRecords from the store
func (k Keeper) RemoveTranscriptOfRecords(ctx sdk.Context) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TranscriptOfRecordsKey))
	store.Delete([]byte{0})
}
