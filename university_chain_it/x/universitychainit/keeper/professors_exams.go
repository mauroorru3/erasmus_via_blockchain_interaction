package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"university_chain_it/x/universitychainit/types"
)

// SetProfessorsExams set a specific professorsExams in the store from its index
func (k Keeper) SetProfessorsExams(ctx sdk.Context, professorsExams types.ProfessorsExams) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ProfessorsExamsKeyPrefix))
	b := k.cdc.MustMarshal(&professorsExams)
	store.Set(types.ProfessorsExamsKey(
		professorsExams.ExamName,
	), b)
}

// GetProfessorsExams returns a professorsExams from its index
func (k Keeper) GetProfessorsExams(
	ctx sdk.Context,
	examName string,

) (val types.ProfessorsExams, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ProfessorsExamsKeyPrefix))

	b := store.Get(types.ProfessorsExamsKey(
		examName,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveProfessorsExams removes a professorsExams from the store
func (k Keeper) RemoveProfessorsExams(
	ctx sdk.Context,
	examName string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ProfessorsExamsKeyPrefix))
	store.Delete(types.ProfessorsExamsKey(
		examName,
	))
}

// GetAllProfessorsExams returns all professorsExams
func (k Keeper) GetAllProfessorsExams(ctx sdk.Context) (list []types.ProfessorsExams) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ProfessorsExamsKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.ProfessorsExams
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
