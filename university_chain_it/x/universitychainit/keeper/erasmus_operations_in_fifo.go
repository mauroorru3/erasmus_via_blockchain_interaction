package keeper

import (
	"encoding/json"
	"time"
	"university_chain_it/x/universitychainit/types"
	"university_chain_it/x/universitychainit/utilfunc"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// function to remove the operation from the queue of currently active operations

func (k Keeper) RemoveOperationFromFifo(ctx sdk.Context, student *types.StoredStudent, uniInfo *types.UniversityInfo) {

	utilfunc.PrintLogs("RemoveOperationFromFifo", ctx)

	if student.OperationInfo.PreviousStudentOperationFifo != "" {
		beforeElement, found := k.GetStoredStudent(ctx, student.OperationInfo.PreviousStudentOperationFifo)
		if !found {
			panic("Element before in Fifo was not found")
		}
		beforeElement.OperationInfo.NextStudentOperationFifo = student.OperationInfo.NextStudentOperationFifo
		k.SetStoredStudent(ctx, beforeElement)
		if student.OperationInfo.NextStudentOperationFifo == "" {
			uniInfo.FifoTailOperation = beforeElement.Index
		}

	} else if uniInfo.FifoHeadOperation == student.Index {
		uniInfo.FifoHeadOperation = student.OperationInfo.NextStudentOperationFifo
	}
	// Does it have a successor?
	if student.OperationInfo.NextStudentOperationFifo != "" {
		afterElement, found := k.GetStoredStudent(ctx, student.OperationInfo.NextStudentOperationFifo)
		if !found {
			panic("Element after in Fifo was not found")
		}
		afterElement.OperationInfo.PreviousStudentOperationFifo = student.OperationInfo.PreviousStudentOperationFifo
		k.SetStoredStudent(ctx, afterElement)
		if student.OperationInfo.PreviousStudentOperationFifo == "" {
			uniInfo.FifoHeadOperation = afterElement.Index
		}
		// Is it at the FIFO tail?
	} else if uniInfo.FifoTailOperation == student.Index {
		uniInfo.FifoTailOperation = student.OperationInfo.PreviousStudentOperationFifo
	}

	student.OperationInfo.PreviousStudentOperationFifo = ""
	student.OperationInfo.NextStudentOperationFifo = ""

}

// function to insert the operation from the queue of currently active operations

func (k Keeper) InsertOperationInFIFOQueue(ctx sdk.Context, student *types.StoredStudent, uniInfo *types.UniversityInfo) {

	utilfunc.PrintLogs("InsertOperationInFIFOQueue", ctx)

	finish := false
	if uniInfo.FifoHeadOperation == "" && uniInfo.FifoTailOperation == "" {
		uniInfo.FifoHeadOperation = student.Index
		uniInfo.FifoTailOperation = student.Index
		student.OperationInfo.PreviousStudentOperationFifo = ""
		student.OperationInfo.NextStudentOperationFifo = ""
	} else {
		currentTail, found := k.GetStoredStudent(ctx, uniInfo.FifoTailOperation)
		if !found {
			panic("Current Fifo tail was not found")
		}
		goOn, err := CheckOperationRemainingTime(ctx, currentTail, *student)
		if err != nil {
			panic(err)
		}
		if goOn {
			if currentTail.OperationInfo.PreviousStudentOperationFifo != "" {
				currentTail, found = k.GetStoredStudent(ctx, currentTail.OperationInfo.PreviousStudentOperationFifo)
				if !found {
					panic("Previous student was not found")
				}
			} else {
				uniInfo.FifoHeadOperation = student.Index
				student.OperationInfo.PreviousStudentOperationFifo = ""
				student.OperationInfo.NextStudentOperationFifo = currentTail.Index
				currentTail.OperationInfo.PreviousStudentOperationFifo = student.Index
				k.SetStoredStudent(ctx, currentTail)
				finish = true

			}
		} else {
			currentTail.OperationInfo.NextStudentOperationFifo = student.Index
			k.SetStoredStudent(ctx, currentTail)
			student.OperationInfo.PreviousStudentOperationFifo = currentTail.Index
			uniInfo.FifoTailOperation = student.Index
			finish = true
		}
		for !finish {
			goOn, err := CheckOperationRemainingTime(ctx, currentTail, *student)
			if err != nil {
				panic(err)
			}
			if goOn {
				if currentTail.OperationInfo.PreviousStudentOperationFifo != "" {
					currentTail, found = k.GetStoredStudent(ctx, currentTail.OperationInfo.PreviousStudentOperationFifo)
					if !found {
						panic("Previous student was not found")
					}
				} else {
					uniInfo.FifoHeadOperation = student.Index
					student.OperationInfo.PreviousStudentOperationFifo = ""
					student.OperationInfo.NextStudentOperationFifo = currentTail.Index
					currentTail.OperationInfo.PreviousStudentOperationFifo = student.Index
					k.SetStoredStudent(ctx, currentTail)
					finish = true

				}
			} else {
				student.OperationInfo.NextStudentOperationFifo = currentTail.OperationInfo.NextStudentOperationFifo
				student.OperationInfo.PreviousStudentOperationFifo = currentTail.Index
				currentTail.OperationInfo.NextStudentOperationFifo = student.Index
				k.SetStoredStudent(ctx, currentTail)
				finish = true
			}
		}
	}

}

// function that checks the timers of two successive operations for expiration

func CheckOperationRemainingTime(ctx sdk.Context, tail types.StoredStudent, student types.StoredStudent) (ok bool, err error) {

	utilfunc.PrintLogs("CheckOperationRemainingTime", ctx)

	var operationDataTail utilfunc.OperationDataStruct

	err = json.Unmarshal([]byte(tail.OperationInfo.StudentOperationDetails), &operationDataTail)
	if err != nil {
		return ok, err
	}

	finishDate, err := time.Parse(utilfunc.DeadlineLayout, operationDataTail.OperationDeadline)
	if err != nil {
		return ok, err
	}

	startDate := ctx.BlockTime()
	differenceTail := finishDate.Sub(startDate)

	//--------------------

	var operationDataStudent utilfunc.OperationDataStruct

	err = json.Unmarshal([]byte(tail.OperationInfo.StudentOperationDetails), &operationDataStudent)
	if err != nil {
		return ok, err
	}

	finishDate, err = time.Parse(utilfunc.DeadlineLayout, operationDataStudent.OperationDeadline)
	if err != nil {
		return ok, err
	}

	startDate = ctx.BlockTime()

	differenceStudent := finishDate.Sub(startDate)

	if differenceTail.Seconds() > differenceStudent.Seconds() {
		return true, err
	} else {
		return false, err
	}

}

// function that inserts the timer and the id of the ongoing inter-chain operation into the student structure

func (k Keeper) InsertOperationDeadline(ctx sdk.Context, student *types.StoredStudent, operationID string, nRetry int) (err error) {

	utilfunc.PrintLogs("InsertOperationDeadline", ctx)

	var operationData utilfunc.OperationDataStruct

	startDate := ctx.BlockTime()

	operationData.OperationID = operationID

	switch operationID {

	case "1":

		endDate := startDate.Add(time.Duration(600 * time.Second * time.Duration(nRetry)))
		operationData.OperationDeadline = utilfunc.FormatDeadline(endDate)

	case "2":

		endDate := startDate.Add(time.Duration(600 * time.Second * time.Duration(nRetry)))
		operationData.OperationDeadline = utilfunc.FormatDeadline(endDate)

	case "3":

		endDate := startDate.Add(time.Duration(600 * time.Second * time.Duration(nRetry)))
		operationData.OperationDeadline = utilfunc.FormatDeadline(endDate)

	default:
		endDate := startDate.Add(time.Duration(600 * time.Second * time.Duration(nRetry)))
		operationData.OperationDeadline = utilfunc.FormatDeadline(endDate)
	}

	resultByteJSON, err := json.Marshal(operationData)
	if err != nil {
		return err
	}

	student.OperationInfo.StudentOperationDetails = string(resultByteJSON)
	k.SetStoredStudent(ctx, *student)
	return err

}

// function that deletes the timer and the id of the ongoing inter-chain operation into the student structure

func (k Keeper) RemoveOperationDeadline(ctx sdk.Context, student *types.StoredStudent) (err error) {

	utilfunc.PrintLogs("RemoveOperationDeadline", ctx)

	var operationData utilfunc.OperationDataStruct

	operationData.OperationID = ""
	operationData.OperationDeadline = ""

	resultByteJSON, err := json.Marshal(operationData)
	if err != nil {
		return err
	}

	student.OperationInfo.StudentOperationDetails = string(resultByteJSON)
	k.SetStoredStudent(ctx, *student)
	return err

}

// function that returns the value of the operation timer

func GetOperationDeadline(student types.StoredStudent) (date time.Time, err error) {

	var operationData utilfunc.OperationDataStruct

	err = json.Unmarshal([]byte(student.OperationInfo.StudentOperationDetails), &operationData)
	if err != nil {
		return date, err
	}

	finishDate, err := time.Parse(utilfunc.DeadlineLayout, operationData.OperationDeadline)
	if err != nil {
		return date, err
	}

	return finishDate, err
}

func GetOperationID(student types.StoredStudent) (s string, err error) {

	var operationData utilfunc.OperationDataStruct

	err = json.Unmarshal([]byte(student.OperationInfo.StudentOperationDetails), &operationData)
	if err != nil {
		return s, err
	}

	return operationData.OperationID, err
}

// function that removes the timer and operation from the inter-chain operation queue

func (k Keeper) ClearOperationQueue(ctx sdk.Context, student *types.StoredStudent) (err error) {

	utilfunc.PrintLogs("ClearOperationQueue", ctx)

	homeUni, found := k.GetUniversityInfo(ctx, student.StudentData.UniversityName)
	if !found {
		return err
	}

	err = k.RemoveOperationDeadline(ctx, student)
	if err != nil {
		return err
	}
	k.RemoveOperationFromFifo(ctx, student, &homeUni)

	k.SetUniversityInfo(ctx, homeUni)
	k.SetStoredStudent(ctx, *student)
	return err
}

// function that adds the timer and operation from the inter-chain operation queue

func (k Keeper) AddOperationQueue(ctx sdk.Context, student *types.StoredStudent, homeUni *types.UniversityInfo, operationType string, nRetry int) (err error) {

	utilfunc.PrintLogs("AddOperationQueue", ctx)

	err = k.InsertOperationDeadline(ctx, student, operationType, nRetry)
	if err != nil {
		return err

	}
	k.InsertOperationInFIFOQueue(ctx, student, homeUni)

	k.SetUniversityInfo(ctx, *homeUni)
	k.SetStoredStudent(ctx, *student)
	return err
}
