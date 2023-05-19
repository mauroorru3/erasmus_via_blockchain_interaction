package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSendEndErasmusPeriodRequest = "send_end_erasmus_period_request"

var _ sdk.Msg = &MsgSendEndErasmusPeriodRequest{}

func NewMsgSendEndErasmusPeriodRequest(
	creator string,
	port string,
	channelID string,
	timeoutTimestamp uint64,
	startingUniversityName string,
	destinationUniversityName string,
	index string,
	foreignIndex string,
	homeIndex string,
) *MsgSendEndErasmusPeriodRequest {
	return &MsgSendEndErasmusPeriodRequest{
		Creator:                   creator,
		Port:                      port,
		ChannelID:                 channelID,
		TimeoutTimestamp:          timeoutTimestamp,
		StartingUniversityName:    startingUniversityName,
		DestinationUniversityName: destinationUniversityName,
		Index:                     index,
		ForeignIndex:              foreignIndex,
		HomeIndex:                 homeIndex,
	}
}

func (msg *MsgSendEndErasmusPeriodRequest) Route() string {
	return RouterKey
}

func (msg *MsgSendEndErasmusPeriodRequest) Type() string {
	return TypeMsgSendEndErasmusPeriodRequest
}

func (msg *MsgSendEndErasmusPeriodRequest) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSendEndErasmusPeriodRequest) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSendEndErasmusPeriodRequest) ValidateBasic() error {
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
