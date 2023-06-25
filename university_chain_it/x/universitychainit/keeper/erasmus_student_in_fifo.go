package keeper

import (
	"encoding/json"
	"time"
	"university_chain_it/x/universitychainit/types"
	"university_chain_it/x/universitychainit/utilfunc"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) RemoveFromFifo(ctx sdk.Context, student *types.StoredStudent, uniInfo *types.UniversityInfo) {

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

	err := utilfunc.CloseErasmusPeriod(ctx, student)
	if err != nil {
		panic("Error in CloseErasmusPeriod function")
	}
}

func (k Keeper) InsertInTheErasmusFIFOQueue(ctx sdk.Context, student *types.StoredStudent, uniInfo *types.UniversityInfo) {

	finish := false
	if uniInfo.FifoHeadErasmus == "" && uniInfo.FifoTailErasmus == "" {
		uniInfo.FifoHeadErasmus = student.Index
		uniInfo.FifoTailErasmus = student.Index
		student.ErasmusData.PreviousStudentFifo = ""
		student.ErasmusData.NextStudentFifo = ""
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

func (k Keeper) CheckAndInCaseMoveStutent(ctx sdk.Context, student *types.StoredStudent, uniInfo *types.UniversityInfo) {

	if student.ErasmusData.NextStudentFifo != "" {
		nextElem, found := k.GetStoredStudent(ctx, student.ErasmusData.NextStudentFifo)
		if !found {
			panic("Element after was not found")
		}
		count := 0
		finish := false
		for !finish {
			goOn, err := CheckRemainingTime(nextElem, *student)
			if err != nil {
				panic(err)
			}
			if !goOn { //-> next element has less remaining time
				if nextElem.ErasmusData.NextStudentFifo != "" {
					nextElem, found = k.GetStoredStudent(ctx, nextElem.ErasmusData.NextStudentFifo)
					if !found {
						panic("Previous student was not found")
					}
					count++
				} else { // NextStudentFifo="" so last element
					if student.ErasmusData.PreviousStudentFifo == "" { // PreviousStudentFifo="" originar position was head
						uniInfo.FifoHeadErasmus = student.ErasmusData.NextStudentFifo
						rightElem, found := k.GetStoredStudent(ctx, student.ErasmusData.NextStudentFifo)
						if !found {
							panic("Element after was not found")
						}
						rightElem.ErasmusData.PreviousStudentFifo = ""
						k.SetStoredStudent(ctx, rightElem)
						uniInfo.FifoTailErasmus = student.Index
						student.ErasmusData.NextStudentFifo = ""
						student.ErasmusData.PreviousStudentFifo = nextElem.Index
						nextElem.ErasmusData.NextStudentFifo = student.Index
						k.SetStoredStudent(ctx, nextElem)
					} else {
						// update previous element pointer
						previousElem, found := k.GetStoredStudent(ctx, student.ErasmusData.PreviousStudentFifo)
						if !found {
							panic("Element after was not found")
						}
						previousElem.ErasmusData.NextStudentFifo = student.ErasmusData.NextStudentFifo
						k.SetStoredStudent(ctx, previousElem)
						// update next element pointer
						rightElem, found := k.GetStoredStudent(ctx, student.ErasmusData.NextStudentFifo)
						if !found {
							panic("Element after was not found")
						}
						rightElem.ErasmusData.PreviousStudentFifo = student.ErasmusData.PreviousStudentFifo
						k.SetStoredStudent(ctx, rightElem)
						// update new positions
						uniInfo.FifoTailErasmus = student.Index
						student.ErasmusData.NextStudentFifo = ""
						student.ErasmusData.PreviousStudentFifo = nextElem.Index
						nextElem.ErasmusData.NextStudentFifo = student.Index
						k.SetStoredStudent(ctx, nextElem)
					}

					finish = true

				}
			} else { //-> next element has longer remaining time
				if count > 0 {
					if student.ErasmusData.PreviousStudentFifo == "" { // PreviousStudentFifo="" originar position was head
						uniInfo.FifoHeadErasmus = student.ErasmusData.NextStudentFifo
						rightElem, found := k.GetStoredStudent(ctx, student.ErasmusData.NextStudentFifo)
						if !found {
							panic("Element after was not found")
						}
						rightElem.ErasmusData.PreviousStudentFifo = ""
						k.SetStoredStudent(ctx, rightElem)
						student.ErasmusData.NextStudentFifo = nextElem.Index
						student.ErasmusData.PreviousStudentFifo = nextElem.ErasmusData.PreviousStudentFifo
						nextElem.ErasmusData.PreviousStudentFifo = student.Index
						k.SetStoredStudent(ctx, nextElem)
						leftElem, found := k.GetStoredStudent(ctx, nextElem.ErasmusData.PreviousStudentFifo)
						if !found {
							panic("Element after was not found")
						}
						leftElem.ErasmusData.NextStudentFifo = student.Index
						k.SetStoredStudent(ctx, leftElem)
					} else {
						// update previous element pointer
						previousElem, found := k.GetStoredStudent(ctx, student.ErasmusData.PreviousStudentFifo)
						if !found {
							panic("Element after was not found")
						}
						previousElem.ErasmusData.NextStudentFifo = student.ErasmusData.NextStudentFifo
						k.SetStoredStudent(ctx, previousElem)
						// update next element pointer
						rightElem, found := k.GetStoredStudent(ctx, student.ErasmusData.NextStudentFifo)
						if !found {
							panic("Element after was not found")
						}
						rightElem.ErasmusData.PreviousStudentFifo = student.ErasmusData.PreviousStudentFifo
						k.SetStoredStudent(ctx, rightElem)
						// update new positions
						student.ErasmusData.NextStudentFifo = nextElem.Index
						student.ErasmusData.PreviousStudentFifo = nextElem.ErasmusData.PreviousStudentFifo
						nextElem.ErasmusData.PreviousStudentFifo = student.Index
						k.SetStoredStudent(ctx, nextElem)
						leftElem, found := k.GetStoredStudent(ctx, nextElem.ErasmusData.PreviousStudentFifo)
						if !found {
							panic("Element after was not found")
						}
						leftElem.ErasmusData.NextStudentFifo = student.Index
						k.SetStoredStudent(ctx, leftElem)
					}

				}
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
