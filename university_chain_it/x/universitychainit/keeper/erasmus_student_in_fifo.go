package keeper

import (
	"university_chain_it/x/universitychainit/types"
	"university_chain_it/x/universitychainit/utilfunc"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) RemoveFromFifo(ctx sdk.Context, student *types.StoredStudent, uniInfo *types.UniversityInfo) {
	// Does it have a predecessor?

	if student.ErasmusData.PreviousStudentFifo != "" {
		beforeElement, found := k.GetStoredStudent(ctx, student.ErasmusData.PreviousStudentFifo)
		if !found {
			panic("Element before in Fifo was not found")
		}
		beforeElement.ErasmusData.NextStudentFifo = student.ErasmusData.NextStudentFifo
		k.SetStoredStudent(ctx, beforeElement)
		if student.ErasmusData.NextStudentFifo == "" {
			uniInfo.FifoTailErasmus = beforeElement.Index
		}
		// Is it at the FIFO head?
	} else if uniInfo.FifoHeadErasmus == student.Index {
		uniInfo.FifoHeadErasmus = student.ErasmusData.NextStudentFifo
	}
	// Does it have a successor?
	if student.ErasmusData.NextStudentFifo != "" {
		afterElement, found := k.GetStoredStudent(ctx, student.ErasmusData.NextStudentFifo)
		if !found {
			panic("Element after in Fifo was not found")
		}
		afterElement.ErasmusData.PreviousStudentFifo = student.ErasmusData.PreviousStudentFifo
		k.SetStoredStudent(ctx, afterElement)
		if student.ErasmusData.PreviousStudentFifo == "" {
			uniInfo.FifoHeadErasmus = afterElement.Index
		}
		// Is it at the FIFO tail?
	} else if uniInfo.FifoTailErasmus == student.Index {
		uniInfo.FifoTailErasmus = student.ErasmusData.PreviousStudentFifo
	}
	student.ErasmusData.PreviousStudentFifo = ""
	student.ErasmusData.NextStudentFifo = ""

	err := utilfunc.CloseErasmusPeriod(student)
	if err != nil {
		panic("Error in CloseErasmusPeriod function")
	}
}

func (k Keeper) SendToFifoTail(ctx sdk.Context, student *types.StoredStudent, uniInfo *types.UniversityInfo) {

	if uniInfo.FifoHeadErasmus == "" && uniInfo.FifoTailErasmus == "" {
		uniInfo.FifoHeadErasmus = student.Index
		uniInfo.FifoTailErasmus = student.Index
	} else {
		currentTail, found := k.GetStoredStudent(ctx, uniInfo.FifoTailErasmus)
		if !found {
			panic("Current Fifo tail was not found")
		}
		currentTail.ErasmusData.NextStudentFifo = student.Index
		k.SetStoredStudent(ctx, currentTail)

		student.ErasmusData.PreviousStudentFifo = currentTail.Index
		uniInfo.FifoTailErasmus = student.Index
	}
}
