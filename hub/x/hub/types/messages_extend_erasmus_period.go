package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSendExtendErasmusPeriod = "send_extend_erasmus_period"

var _ sdk.Msg = &MsgSendExtendErasmusPeriod{}

func NewMsgSendExtendErasmusPeriod(
	creator string,
	port string,
	channelID string,
	timeoutTimestamp uint64,
	durationInMonths uint64,
	destinationUniversityName string,
	foreignIndex string,
	finalDate string,
) *MsgSendExtendErasmusPeriod {
	return &MsgSendExtendErasmusPeriod{
		Creator:                   creator,
		Port:                      port,
		ChannelID:                 channelID,
		TimeoutTimestamp:          timeoutTimestamp,
		DurationInMonths:          durationInMonths,
		DestinationUniversityName: destinationUniversityName,
		ForeignIndex:              foreignIndex,
		FinalDate:                 finalDate,
	}
}

func (msg *MsgSendExtendErasmusPeriod) Route() string {
	return RouterKey
}

func (msg *MsgSendExtendErasmusPeriod) Type() string {
	return TypeMsgSendExtendErasmusPeriod
}

func (msg *MsgSendExtendErasmusPeriod) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSendExtendErasmusPeriod) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSendExtendErasmusPeriod) ValidateBasic() error {
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
