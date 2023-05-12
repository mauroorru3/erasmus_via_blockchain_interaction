package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgInsertStudentContactInfo = "insert_student_contact_info"

var _ sdk.Msg = &MsgInsertStudentContactInfo{}

func NewMsgInsertStudentContactInfo(creator string, university string, studentIndex string, contactAddress string, email string, mobilePhone string) *MsgInsertStudentContactInfo {
	return &MsgInsertStudentContactInfo{
		Creator:        creator,
		University:     university,
		StudentIndex:   studentIndex,
		ContactAddress: contactAddress,
		Email:          email,
		MobilePhone:    mobilePhone,
	}
}

func (msg *MsgInsertStudentContactInfo) Route() string {
	return RouterKey
}

func (msg *MsgInsertStudentContactInfo) Type() string {
	return TypeMsgInsertStudentContactInfo
}

func (msg *MsgInsertStudentContactInfo) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgInsertStudentContactInfo) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgInsertStudentContactInfo) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
