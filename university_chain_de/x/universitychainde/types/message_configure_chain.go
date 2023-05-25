package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgConfigureChain = "configure_chain"

var _ sdk.Msg = &MsgConfigureChain{}

func NewMsgConfigureChain(creator string) *MsgConfigureChain {
	return &MsgConfigureChain{
		Creator: creator,
	}
}

func (msg *MsgConfigureChain) Route() string {
	return RouterKey
}

func (msg *MsgConfigureChain) Type() string {
	return TypeMsgConfigureChain
}

func (msg *MsgConfigureChain) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgConfigureChain) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgConfigureChain) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
