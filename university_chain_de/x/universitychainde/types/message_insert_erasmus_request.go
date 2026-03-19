package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgInsertErasmusRequest = "insert_erasmus_request"

var _ sdk.Msg = &MsgInsertErasmusRequest{}

func NewMsgInsertErasmusRequest(creator string, university string, studentIndex string, durationInMonths string, foreignUniversityName string, erasmusType string) *MsgInsertErasmusRequest {
	return &MsgInsertErasmusRequest{
		Creator:               creator,
		University:            university,
		StudentIndex:          studentIndex,
		DurationInMonths:      durationInMonths,
		ForeignUniversityName: foreignUniversityName,
		ErasmusType:           erasmusType,
	}
}

func (msg *MsgInsertErasmusRequest) Route() string {
	return RouterKey
}

func (msg *MsgInsertErasmusRequest) Type() string {
	return TypeMsgInsertErasmusRequest
}

func (msg *MsgInsertErasmusRequest) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgInsertErasmusRequest) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgInsertErasmusRequest) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
