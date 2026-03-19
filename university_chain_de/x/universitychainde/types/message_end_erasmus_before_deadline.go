package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgEndErasmusBeforeDeadline = "end_erasmus_before_deadline"

var _ sdk.Msg = &MsgEndErasmusBeforeDeadline{}

func NewMsgEndErasmusBeforeDeadline(creator string, university string, studentIndex string) *MsgEndErasmusBeforeDeadline {
	return &MsgEndErasmusBeforeDeadline{
		Creator:      creator,
		University:   university,
		StudentIndex: studentIndex,
	}
}

func (msg *MsgEndErasmusBeforeDeadline) Route() string {
	return RouterKey
}

func (msg *MsgEndErasmusBeforeDeadline) Type() string {
	return TypeMsgEndErasmusBeforeDeadline
}

func (msg *MsgEndErasmusBeforeDeadline) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgEndErasmusBeforeDeadline) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgEndErasmusBeforeDeadline) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
