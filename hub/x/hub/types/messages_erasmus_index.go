package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSendErasmusIndex = "send_erasmus_index"

var _ sdk.Msg = &MsgSendErasmusIndex{}

func NewMsgSendErasmusIndex(
	creator string,
	port string,
	channelID string,
	timeoutTimestamp uint64,
	chain string,
	index string,
	foreignIndex string,
) *MsgSendErasmusIndex {
	return &MsgSendErasmusIndex{
		Creator:          creator,
		Port:             port,
		ChannelID:        channelID,
		TimeoutTimestamp: timeoutTimestamp,
		Chain:            chain,
		Index:            index,
		ForeignIndex:     foreignIndex,
	}
}

func (msg *MsgSendErasmusIndex) Route() string {
	return RouterKey
}

func (msg *MsgSendErasmusIndex) Type() string {
	return TypeMsgSendErasmusIndex
}

func (msg *MsgSendErasmusIndex) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSendErasmusIndex) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSendErasmusIndex) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	if msg.Port == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid packet port")
	}
	if msg.ChannelID == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid packet channel")
	}
	if msg.TimeoutTimestamp == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid packet timeout")
	}
	return nil
}
