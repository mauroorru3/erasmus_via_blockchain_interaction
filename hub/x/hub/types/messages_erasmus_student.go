package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSendErasmusStudent = "send_erasmus_student"

var _ sdk.Msg = &MsgSendErasmusStudent{}

func NewMsgSendErasmusStudent(
	creator string,
	port string,
	channelID string,
	timeoutTimestamp uint64,
	student *StoredStudent,
) *MsgSendErasmusStudent {
	return &MsgSendErasmusStudent{
		Creator:          creator,
		Port:             port,
		ChannelID:        channelID,
		TimeoutTimestamp: timeoutTimestamp,
		Student:          student,
	}
}

func (msg *MsgSendErasmusStudent) Route() string {
	return RouterKey
}

func (msg *MsgSendErasmusStudent) Type() string {
	return TypeMsgSendErasmusStudent
}

func (msg *MsgSendErasmusStudent) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSendErasmusStudent) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSendErasmusStudent) ValidateBasic() error {
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
