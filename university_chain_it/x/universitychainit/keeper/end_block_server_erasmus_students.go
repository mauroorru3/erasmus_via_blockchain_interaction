package keeper

import (
	"context"
	"time"
	"university_chain_it/x/universitychainit/utilfunc"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) TerminateExpiredErasmusPeriods(goCtx context.Context) {

	ctx := sdk.UnwrapSDKContext(goCtx)

	uniList := k.GetAllUniversityInfo(ctx)
	lenUniList := len(uniList)

	for i := 0; i < lenUniList; i++ {

		studentIndex := uniList[i].FifoHeadErasmus
		finish := false

		for !finish {
			// Finished moving along
			if studentIndex == "" {
				finish = true
			} else {
				storedStudent, found := k.GetStoredStudent(ctx, studentIndex)

				if !found {
					panic("Fifo head game not found " + uniList[i].FifoHeadErasmus)
				}
				deadline, err := utilfunc.GetErasmusDeadline(storedStudent)
				if err != nil {
					panic(err)
				}

				s := utilfunc.FormatDeadline(ctx.BlockTime())
				formattedStartDate, _ := time.Parse(utilfunc.DeadlineLayout, s)
				if deadline.Before(formattedStartDate) {
					//if deadline.Before(ctx.BlockTime()) {

					// Erasmus period is past deadline
					k.RemoveFromFifo(ctx, &storedStudent, &uniList[i])
					k.SetStoredStudent(ctx, storedStudent)
					// Move along FIFO
					studentIndex = uniList[i].FifoHeadErasmus
				} else {
					finish = true
				}

			}
		}
		k.SetUniversityInfo(ctx, uniList[i])

	}

}
