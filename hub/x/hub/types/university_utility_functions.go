package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var universities map[string]int = map[string]int{
	"tum":                 1,
	"humboldt_university": 1,
	"unipi":               1,
	"uniroma1":            1,
}

func (foreignUniversity Universities) GetForeignUniversityAddress() (address sdk.AccAddress, err error) {
	address, errAddress := sdk.AccAddressFromBech32(foreignUniversity.UniversitiesKey)
	return address, sdkerrors.Wrapf(errAddress, ErrInvalidStudentAddress.Error(), foreignUniversity.UniversitiesKey)
}

func (foreignUniversity Universities) Validate() (err error) {
	_, err = foreignUniversity.GetForeignUniversityAddress()
	if err != nil {
		return err
	}
	_, found := universities[foreignUniversity.UniversityName]
	if !found {
		return ErrWrongNameUniversity
	}
	return err
}
