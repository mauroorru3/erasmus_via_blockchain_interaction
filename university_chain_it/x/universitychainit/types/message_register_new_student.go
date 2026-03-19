package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgRegisterNewStudent = "register_new_student"

var _ sdk.Msg = &MsgRegisterNewStudent{}

func NewMsgRegisterNewStudent(creator string, university string, name string, surname string, courseType string, courseOfStudy string, departmentName string) *MsgRegisterNewStudent {
	return &MsgRegisterNewStudent{
		Creator:        creator,
		University:     university,
		Name:           name,
		Surname:        surname,
		CourseType:     courseType,
		CourseOfStudy:  courseOfStudy,
		DepartmentName: departmentName,
	}
}

func (msg *MsgRegisterNewStudent) Route() string {
	return RouterKey
}

func (msg *MsgRegisterNewStudent) Type() string {
	return TypeMsgRegisterNewStudent
}

func (msg *MsgRegisterNewStudent) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgRegisterNewStudent) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRegisterNewStudent) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
