package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgExtendErasmus = "extend_erasmus"

var _ sdk.Msg = &MsgExtendErasmus{}

func NewMsgExtendErasmus(creator string, university string, studentIndex string, durationInMonths string) *MsgExtendErasmus {
	return &MsgExtendErasmus{
		Creator:          creator,
		University:       university,
		StudentIndex:     studentIndex,
		DurationInMonths: durationInMonths,
	}
}

func (msg *MsgExtendErasmus) Route() string {
	return RouterKey
}

func (msg *MsgExtendErasmus) Type() string {
	return TypeMsgExtendErasmus
}

func (msg *MsgExtendErasmus) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgExtendErasmus) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgExtendErasmus) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
