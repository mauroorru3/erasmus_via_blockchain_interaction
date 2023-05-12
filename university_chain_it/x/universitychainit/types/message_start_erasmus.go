package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgStartErasmus = "start_erasmus"

var _ sdk.Msg = &MsgStartErasmus{}

func NewMsgStartErasmus(creator string, university string, studentIndex string) *MsgStartErasmus {
	return &MsgStartErasmus{
		Creator:      creator,
		University:   university,
		StudentIndex: studentIndex,
	}
}

func (msg *MsgStartErasmus) Route() string {
	return RouterKey
}

func (msg *MsgStartErasmus) Type() string {
	return TypeMsgStartErasmus
}

func (msg *MsgStartErasmus) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgStartErasmus) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgStartErasmus) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
