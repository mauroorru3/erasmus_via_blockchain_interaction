package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgInsertExamGrade = "insert_exam_grade"

var _ sdk.Msg = &MsgInsertExamGrade{}

func NewMsgInsertExamGrade(creator string, university string, studentIndex string, examName string, grade string) *MsgInsertExamGrade {
	return &MsgInsertExamGrade{
		Creator:      creator,
		University:   university,
		StudentIndex: studentIndex,
		ExamName:     examName,
		Grade:        grade,
	}
}

func (msg *MsgInsertExamGrade) Route() string {
	return RouterKey
}

func (msg *MsgInsertExamGrade) Type() string {
	return TypeMsgInsertExamGrade
}

func (msg *MsgInsertExamGrade) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgInsertExamGrade) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgInsertExamGrade) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
