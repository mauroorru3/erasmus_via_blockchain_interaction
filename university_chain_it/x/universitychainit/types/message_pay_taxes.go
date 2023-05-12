package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgPayTaxes = "pay_taxes"

var _ sdk.Msg = &MsgPayTaxes{}

func NewMsgPayTaxes(creator string, university string, studentIndex string) *MsgPayTaxes {
	return &MsgPayTaxes{
		Creator:      creator,
		University:   university,
		StudentIndex: studentIndex,
	}
}

func (msg *MsgPayTaxes) Route() string {
	return RouterKey
}

func (msg *MsgPayTaxes) Type() string {
	return TypeMsgPayTaxes
}

func (msg *MsgPayTaxes) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgPayTaxes) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgPayTaxes) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
