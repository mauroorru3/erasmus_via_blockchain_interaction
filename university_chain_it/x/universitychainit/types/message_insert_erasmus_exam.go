package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgInsertErasmusExam = "insert_erasmus_exam"

var _ sdk.Msg = &MsgInsertErasmusExam{}

func NewMsgInsertErasmusExam(creator string, university string, studentIndex string, examName string) *MsgInsertErasmusExam {
	return &MsgInsertErasmusExam{
		Creator:      creator,
		University:   university,
		StudentIndex: studentIndex,
		ExamName:     examName,
	}
}

func (msg *MsgInsertErasmusExam) Route() string {
	return RouterKey
}

func (msg *MsgInsertErasmusExam) Type() string {
	return TypeMsgInsertErasmusExam
}

func (msg *MsgInsertErasmusExam) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgInsertErasmusExam) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgInsertErasmusExam) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
