package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSendFinalErasmusData = "send_final_erasmus_data"

var _ sdk.Msg = &MsgSendFinalErasmusData{}

func NewMsgSendFinalErasmusData(
	creator string,
	port string,
	channelID string,
	timeoutTimestamp uint64,
	erasmusData *ErasmusInfo,
	homeIndex string,
) *MsgSendFinalErasmusData {
	return &MsgSendFinalErasmusData{
		Creator:          creator,
		Port:             port,
		ChannelID:        channelID,
		TimeoutTimestamp: timeoutTimestamp,
		ErasmusData:      erasmusData,
		HomeIndex:        homeIndex,
	}
}

func (msg *MsgSendFinalErasmusData) Route() string {
	return RouterKey
}

func (msg *MsgSendFinalErasmusData) Type() string {
	return TypeMsgSendFinalErasmusData
}

func (msg *MsgSendFinalErasmusData) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSendFinalErasmusData) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSendFinalErasmusData) ValidateBasic() error {
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
