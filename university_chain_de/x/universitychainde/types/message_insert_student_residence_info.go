package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgInsertStudentResidenceInfo = "insert_student_residence_info"

var _ sdk.Msg = &MsgInsertStudentResidenceInfo{}

func NewMsgInsertStudentResidenceInfo(creator string, university string, studentIndex string, country string, province string, town string, postCode string, address string, houseNumber string, homePhone string) *MsgInsertStudentResidenceInfo {
	return &MsgInsertStudentResidenceInfo{
		Creator:      creator,
		University:   university,
		StudentIndex: studentIndex,
		Country:      country,
		Province:     province,
		Town:         town,
		PostCode:     postCode,
		Address:      address,
		HouseNumber:  houseNumber,
		HomePhone:    homePhone,
	}
}

func (msg *MsgInsertStudentResidenceInfo) Route() string {
	return RouterKey
}

func (msg *MsgInsertStudentResidenceInfo) Type() string {
	return TypeMsgInsertStudentResidenceInfo
}

func (msg *MsgInsertStudentResidenceInfo) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgInsertStudentResidenceInfo) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgInsertStudentResidenceInfo) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
