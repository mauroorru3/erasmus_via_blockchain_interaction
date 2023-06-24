package keeper

import (
	"encoding/json"
	"time"
	"university_chain_de/x/universitychainde/types"
	"university_chain_de/x/universitychainde/utilfunc"

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

func (k Keeper) InsertInTheErasmusFIFOQueue(ctx sdk.Context, student *types.StoredStudent, uniInfo *types.UniversityInfo) {

	utilfunc.PrintLogs("InsertInTheErasmusFIFOQueue")
	finish := false
	if uniInfo.FifoHeadErasmus == "" && uniInfo.FifoTailErasmus == "" {
		uniInfo.FifoHeadErasmus = student.Index
		uniInfo.FifoTailErasmus = student.Index
		student.ErasmusData.PreviousStudentFifo = ""
		student.ErasmusData.NextStudentFifo = ""
		utilfunc.PrintLogs("InsertInTheErasmusFIFOQueue qui 1")
	} else {
		currentTail, found := k.GetStoredStudent(ctx, uniInfo.FifoTailErasmus)
		if !found {
			panic("Current Fifo tail was not found")
		}
		goOn, err := CheckRemainingTime(currentTail, *student)
		if err != nil {
			panic(err)
		}
		if goOn {
			if currentTail.ErasmusData.PreviousStudentFifo != "" {
				utilfunc.PrintLogs("InsertInTheErasmusFIFOQueue qui 2")
				currentTail, found = k.GetStoredStudent(ctx, currentTail.ErasmusData.PreviousStudentFifo)
				if !found {
					panic("Previous student was not found")
				}
			} else {
				utilfunc.PrintLogs("InsertInTheErasmusFIFOQueue qui 3")
				uniInfo.FifoHeadErasmus = student.Index
				student.ErasmusData.PreviousStudentFifo = ""
				student.ErasmusData.NextStudentFifo = currentTail.Index
				currentTail.ErasmusData.PreviousStudentFifo = student.Index
				k.SetStoredStudent(ctx, currentTail)
				finish = true

			}
		} else {
			utilfunc.PrintLogs("InsertInTheErasmusFIFOQueue qui 4")
			currentTail.ErasmusData.NextStudentFifo = student.Index
			k.SetStoredStudent(ctx, currentTail)
			student.ErasmusData.PreviousStudentFifo = currentTail.Index
			uniInfo.FifoTailErasmus = student.Index
			finish = true
		}
		for !finish {
			goOn, err := CheckRemainingTime(currentTail, *student)
			if err != nil {
				panic(err)
			}
			if goOn {
				if currentTail.ErasmusData.PreviousStudentFifo != "" {
					currentTail, found = k.GetStoredStudent(ctx, currentTail.ErasmusData.PreviousStudentFifo)
					if !found {
						panic("Previous student was not found")
					}
				} else {
					uniInfo.FifoHeadErasmus = student.Index
					student.ErasmusData.PreviousStudentFifo = ""
					student.ErasmusData.NextStudentFifo = currentTail.Index
					currentTail.ErasmusData.PreviousStudentFifo = student.Index
					k.SetStoredStudent(ctx, currentTail)
					finish = true

				}
			} else {
				student.ErasmusData.NextStudentFifo = currentTail.ErasmusData.NextStudentFifo
				student.ErasmusData.PreviousStudentFifo = currentTail.Index
				currentTail.ErasmusData.NextStudentFifo = student.Index
				k.SetStoredStudent(ctx, currentTail)
				finish = true
			}
		}
	}

}

func CheckRemainingTime(tail types.StoredStudent, student types.StoredStudent) (ok bool, err error) {

	var erasmusCareerTail []utilfunc.ErasmusCareerStruct

	err = json.Unmarshal([]byte(tail.ErasmusData.Career), &erasmusCareerTail)
	if err != nil {
		return ok, err
	}

	lenCareer := len(erasmusCareerTail)

	finishDate, err := time.Parse(utilfunc.DeadlineLayout, erasmusCareerTail[lenCareer-1].End_date)
	if err != nil {
		return ok, err
	}

	startDate, err := time.Parse(utilfunc.DeadlineLayout, erasmusCareerTail[lenCareer-1].Start_date)
	if err != nil {
		return ok, err
	}

	differenceTail := finishDate.Sub(startDate)

	//--------------------

	var erasmusCareerStudent []utilfunc.ErasmusCareerStruct

	err = json.Unmarshal([]byte(student.ErasmusData.Career), &erasmusCareerStudent)
	if err != nil {
		return ok, err
	}

	lenCareerStudent := len(erasmusCareerStudent)

	finishDate, err = time.Parse(utilfunc.DeadlineLayout, erasmusCareerStudent[lenCareerStudent-1].End_date)
	if err != nil {
		return ok, err
	}

	startDate, err = time.Parse(utilfunc.DeadlineLayout, erasmusCareerStudent[lenCareerStudent-1].Start_date)
	if err != nil {
		return ok, err
	}

	differenceStudent := finishDate.Sub(startDate)

	if differenceTail.Seconds() > differenceStudent.Seconds() {
		return true, err
	} else {
		return false, err
	}

}
