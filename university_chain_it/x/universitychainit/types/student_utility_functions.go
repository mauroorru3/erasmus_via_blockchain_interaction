package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (storedStudent StoredStudent) GetStudentAddress() (address sdk.AccAddress, err error) {
	address, errAddress := sdk.AccAddressFromBech32(storedStudent.StudentData.StudentKey)
	return address, sdkerrors.Wrapf(errAddress, ErrInvalidStudentAddress.Error(), storedStudent.StudentData.StudentKey)
}

func (storedStudent StoredStudent) Validate() (err error) {
	_, err = storedStudent.GetStudentAddress()
	if err != nil {
		return err
	}
	return err
}
