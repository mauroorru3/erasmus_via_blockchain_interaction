package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgInsertStudentPersonalInfo = "insert_student_personal_info"

var _ sdk.Msg = &MsgInsertStudentPersonalInfo{}

func NewMsgInsertStudentPersonalInfo(creator string, university string, studentIndex string, gender string, dateOfBirth string, primaryNationality string, countryOfBirth string, provinceOfBirth string, townOfBirth string, taxCode string, incomeBracket uint32) *MsgInsertStudentPersonalInfo {
	return &MsgInsertStudentPersonalInfo{
		Creator:            creator,
		University:         university,
		StudentIndex:       studentIndex,
		Gender:             gender,
		DateOfBirth:        dateOfBirth,
		PrimaryNationality: primaryNationality,
		CountryOfBirth:     countryOfBirth,
		ProvinceOfBirth:    provinceOfBirth,
		TownOfBirth:        townOfBirth,
		TaxCode:            taxCode,
		IncomeBracket:      incomeBracket,
	}
}

func (msg *MsgInsertStudentPersonalInfo) Route() string {
	return RouterKey
}

func (msg *MsgInsertStudentPersonalInfo) Type() string {
	return TypeMsgInsertStudentPersonalInfo
}

func (msg *MsgInsertStudentPersonalInfo) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgInsertStudentPersonalInfo) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgInsertStudentPersonalInfo) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
